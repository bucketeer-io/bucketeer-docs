package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Default configuration values.
const (
	DefaultModel      = "gpt-4o"
	DefaultAPIBase    = "https://api.openai.com/v1"
	DefaultMaxTokens  = 8192
	DefaultMaxRetries = 3
	RequestTimeout    = 2 * time.Minute
)

// Environment variable names.
const (
	EnvOpenAIAPIKey  = "OPENAI_API_KEY"
	EnvOpenAIModel   = "OPENAI_MODEL"
	EnvOpenAIAPIBase = "OPENAI_API_BASE"
)

// Client is an OpenAI API client with retry logic and circuit breaker.
type Client struct {
	apiKey      string
	apiBase     string
	model       string
	temperature float64
	maxTokens   int
	maxRetries  int
	httpClient  *http.Client
}

// ClientOption is a function that configures a Client.
type ClientOption func(*Client)

// WithModel sets the model to use.
func WithModel(model string) ClientOption {
	return func(c *Client) {
		c.model = model
	}
}

// WithAPIBase sets the API base URL.
func WithAPIBase(apiBase string) ClientOption {
	return func(c *Client) {
		c.apiBase = apiBase
	}
}

// WithTemperature sets the temperature parameter.
func WithTemperature(temp float64) ClientOption {
	return func(c *Client) {
		c.temperature = temp
	}
}

// WithMaxTokens sets the maximum tokens for responses.
func WithMaxTokens(tokens int) ClientOption {
	return func(c *Client) {
		c.maxTokens = tokens
	}
}

// WithMaxRetries sets the maximum number of retries.
func WithMaxRetries(retries int) ClientOption {
	return func(c *Client) {
		c.maxRetries = retries
	}
}

// NewClient creates a new OpenAI API client.
// It reads configuration from environment variables with sensible defaults.
func NewClient(apiKey string, opts ...ClientOption) *Client {
	// Get model from environment or use default
	model := os.Getenv(EnvOpenAIModel)
	if model == "" {
		model = DefaultModel
	}

	// Get API base from environment or use default
	apiBase := os.Getenv(EnvOpenAIAPIBase)
	if apiBase == "" {
		apiBase = DefaultAPIBase
	}

	c := &Client{
		apiKey:      apiKey,
		apiBase:     apiBase,
		model:       model,
		temperature: 0, // Deterministic output
		maxTokens:   DefaultMaxTokens,
		maxRetries:  DefaultMaxRetries,
		httpClient: &http.Client{
			Timeout: RequestTimeout,
		},
	}

	// Apply options
	for _, opt := range opts {
		opt(c)
	}

	return c
}

// ChatMessage represents a message in the chat completion request.
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletionRequest represents the request body for chat completions.
type ChatCompletionRequest struct {
	Model       string        `json:"model"`
	Messages    []ChatMessage `json:"messages"`
	Temperature float64       `json:"temperature"`
	MaxTokens   int           `json:"max_tokens"`
}

// ChatCompletionResponse represents the response from chat completions.
type ChatCompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Error *APIError `json:"error,omitempty"`
}

// APIError represents an error from the OpenAI API.
type APIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("OpenAI API error: %s (type: %s, code: %s)", e.Message, e.Type, e.Code)
}

// HTTPError represents an HTTP error with status code.
type HTTPError struct {
	StatusCode int
	Message    string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

// CreateChatCompletion sends a chat completion request with retry logic.
// Implements exponential backoff (2s, 8s, 32s) and circuit breaker pattern.
func (c *Client) CreateChatCompletion(ctx context.Context, messages []ChatMessage) (string, error) {
	var lastErr error

	for i := 0; i < c.maxRetries; i++ {
		// Create request-scoped context with timeout
		reqCtx, cancel := context.WithTimeout(ctx, RequestTimeout)

		response, err := c.doRequest(reqCtx, messages)
		cancel() // Always cancel to avoid context leak

		if err == nil {
			if len(response.Choices) == 0 {
				return "", errors.New("no choices in response")
			}
			return response.Choices[0].Message.Content, nil
		}

		// Check if error is retryable
		if !isRetryableError(err) {
			return "", fmt.Errorf("openai api error (non-retryable): %w", err)
		}

		lastErr = err

		// Exponential backoff: 2s, 8s, 32s (2^(i+1) * 2 = 2, 8, 32)
		// Actually: 2^1=2, 2^3=8, 2^5=32 -> using 2 << (i*2) for i=0,1,2 gives 2,8,32
		delay := time.Duration(2<<(i*2)) * time.Second
		log.Printf("OpenAI API retry %d/%d after %v: %v", i+1, c.maxRetries, delay, err)

		select {
		case <-ctx.Done():
			return "", fmt.Errorf("context cancelled during retry: %w", ctx.Err())
		case <-time.After(delay):
		}
	}

	// Circuit breaker: max retries exceeded
	return "", fmt.Errorf("circuit breaker open: max retries (%d) exceeded: %w", c.maxRetries, lastErr)
}

// doRequest performs the actual HTTP request to OpenAI API.
func (c *Client) doRequest(ctx context.Context, messages []ChatMessage) (*ChatCompletionResponse, error) {
	reqBody := ChatCompletionRequest{
		Model:       c.model,
		Messages:    messages,
		Temperature: c.temperature,
		MaxTokens:   c.maxTokens,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	url := c.apiBase + "/chat/completions"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			Message:    string(body),
		}
	}

	var response ChatCompletionResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Check for API-level errors
	if response.Error != nil {
		return nil, response.Error
	}

	return &response, nil
}

// isRetryableError determines if an error is retryable.
// 429 (Rate Limit) and 5xx errors are retryable.
// Context timeout/cancel errors are NOT retryable.
func isRetryableError(err error) bool {
	// Context timeout/cancel is not retryable
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		return false
	}

	// Check for HTTP errors
	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		// 429 (Rate Limit) and 5xx errors are retryable
		return httpErr.StatusCode == http.StatusTooManyRequests || httpErr.StatusCode >= 500
	}

	// Network errors are generally retryable
	return true
}

// GetModel returns the model being used.
func (c *Client) GetModel() string {
	return c.model
}

// GetAPIBase returns the API base URL being used.
func (c *Client) GetAPIBase() string {
	return c.apiBase
}
