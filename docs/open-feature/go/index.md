---
title: Go
slug: /open-feature/go
---

# Go Server Provider

:::warning Go Provider (Beta)

This is a beta version. Breaking changes may be introduced before general release.

:::

This is the official Go OpenFeature server-side provider for [Bucketeer](https://bucketeer.io/).

### Installation

```bash showLineNumbers
go get github.com/bucketeer-io/openfeature-go-server-sdk
```

**Requirements:** Go 1.21+

### Usage

#### Initialize the provider

This example follows the standard OpenFeature pattern: create a provider, set it globally, and then create clients to evaluate feature flags.

```go showLineNumbers
import (
	"context"
	"github.com/bucketeer-io/go-server-sdk/pkg/bucketeer"
	provider "github.com/bucketeer-io/openfeature-go-server-sdk/pkg"
	"github.com/open-feature/go-sdk/openfeature"
)

func main() {
	// SDK configuration
	options := provider.ProviderOptions{
		bucketeer.WithAPIKey("YOUR_API_KEY"),
		bucketeer.WithAPIEndpoint("YOUR_API_ENDPOINT"),
		bucketeer.WithTag("YOUR_FEATURE_TAG"),
	}

	// Create provider
	p, err := provider.NewProviderWithContext(context.Background(), options)
	if err != nil {
		// Error handling
	}

	// Set the provider and create client
	openfeature.SetProvider(p)
	client := openfeature.NewClient("my-app")

	// User configuration
	userID := "user-123"
	evalCtx := openfeature.NewEvaluationContext(userID, map[string]interface{}{
		openfeature.TargetingKey: userID,
	})

	// Evaluate feature flag
	result, err := client.BooleanValueDetails(context.Background(), "feature-flag-id", false, evalCtx)
	if err != nil {
		// Handle error
	}
}
```

See our [documentation](https://docs.bucketeer.io/sdk/server-side/go) for more SDK configuration.

#### Evaluation Context

The evaluation context allows the client to specify contextual data that Bucketeer uses to evaluate the feature flags. The `targetingKey` is the user ID (Unique ID) and cannot be empty.

```go showLineNumbers
userID := "user-123"
evalCtx := openfeature.NewEvaluationContext(userID, map[string]interface{}{
    openfeature.TargetingKey: userID,
    "department": "engineering",
    "role": "developer",
})
```

#### Evaluate a feature flag

The OpenFeature client supports evaluating different types of feature flags.

```go showLineNumbers
// Boolean Evaluation
result, err := client.BooleanValueDetails(context.Background(), "bool-flag", false, evalCtx)

// String Evaluation
result, err := client.StringValueDetails(context.Background(), "string-flag", "default", evalCtx)

// Integer Evaluation
result, err := client.IntValueDetails(context.Background(), "int-flag", int64(100), evalCtx)

// Float Evaluation
result, err := client.FloatValueDetails(context.Background(), "float-flag", 3.14, evalCtx)

// Object Evaluation
defaultObject := map[string]interface{}{"key": "value"}
result, err := client.ObjectValueDetails(context.Background(), "obj-flag", defaultObject, evalCtx)
```
