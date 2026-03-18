package openai

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	sdkopenai "github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/shared"
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

// Client is an OpenAI API client using the official SDK.
type Client struct {
	sdkClient   sdkopenai.Client
	model       string
	apiBase     string
	temperature float64
	maxTokens   int
	maxRetries  int
}

// ChatMessage represents a message in the chat completion request.
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
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

// NewClient creates a new OpenAI API client using the official SDK.
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
		model:       model,
		apiBase:     apiBase,
		temperature: 0, // Deterministic output
		maxTokens:   DefaultMaxTokens,
		maxRetries:  DefaultMaxRetries,
	}

	// Apply options
	for _, opt := range opts {
		opt(c)
	}

	// Initialize SDK client with configured values
	c.sdkClient = sdkopenai.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL(c.apiBase),
		option.WithMaxRetries(c.maxRetries),
		option.WithHTTPClient(&http.Client{Timeout: RequestTimeout}),
	)

	return c
}

// RequestOption configures a per-request option for chat completions.
// This is distinct from ClientOption which configures the client itself.
type RequestOption func(*sdkopenai.ChatCompletionNewParams)

// WithJSONResponse sets response_format to json_object, guaranteeing valid JSON output.
// IMPORTANT: The system or user message MUST contain the word "JSON"
// when using this option, otherwise the API returns a 400 error.
func WithJSONResponse() RequestOption {
	return func(p *sdkopenai.ChatCompletionNewParams) {
		jsonObj := shared.NewResponseFormatJSONObjectParam()
		p.ResponseFormat = sdkopenai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONObject: &jsonObj,
		}
	}
}

// CreateChatCompletion sends a chat completion request via the SDK.
// The SDK handles retry logic with exponential backoff automatically.
func (c *Client) CreateChatCompletion(ctx context.Context, messages []ChatMessage, opts ...RequestOption) (string, error) {
	params := sdkopenai.ChatCompletionNewParams{
		Model:               c.model,
		Messages:            toSDKMessages(messages),
		Temperature:         sdkopenai.Float(c.temperature),
		MaxCompletionTokens: sdkopenai.Int(int64(c.maxTokens)),
	}

	// Apply per-request options
	for _, opt := range opts {
		opt(&params)
	}

	log.Printf("Calling OpenAI API (model: %s)", c.model)
	completion, err := c.sdkClient.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", fmt.Errorf("openai api error: %w", err)
	}

	if len(completion.Choices) == 0 {
		return "", errors.New("no choices in response")
	}

	return completion.Choices[0].Message.Content, nil
}

// toSDKMessages converts ChatMessage slice to SDK message union types.
func toSDKMessages(messages []ChatMessage) []sdkopenai.ChatCompletionMessageParamUnion {
	result := make([]sdkopenai.ChatCompletionMessageParamUnion, len(messages))
	for i, m := range messages {
		switch m.Role {
		case "system":
			result[i] = sdkopenai.SystemMessage(m.Content)
		case "user":
			result[i] = sdkopenai.UserMessage(m.Content)
		case "assistant":
			result[i] = sdkopenai.AssistantMessage(m.Content)
		default:
			log.Printf("WARNING: unknown message role %q, treating as user message", m.Role)
			result[i] = sdkopenai.UserMessage(m.Content)
		}
	}
	return result
}

// GetModel returns the model being used.
func (c *Client) GetModel() string {
	return c.model
}

// GetAPIBase returns the API base URL being used.
func (c *Client) GetAPIBase() string {
	return c.apiBase
}
