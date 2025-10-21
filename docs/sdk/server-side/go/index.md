---
title: Go
sidebar_position: 1
slug: /sdk/server-side/go
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

This category contains topics explaining how to configure Bucketeer's Go SDK.

:::tip Go SDK Version (Stable)

Bucketeer Go SDK has reached the production stage, offering you a stable and reliable experience.

:::

## Getting started

Before starting, ensure that you follow the [Getting Started](/getting-started) guide.

### Installing dependency

Install the latest version of the client using the following command.

<Tabs>
<TabItem value="go" label="Go">

```go showLineNumbers
go get github.com/bucketeer-io/go-server-sdk
```

</TabItem>
</Tabs>

### Importing SDK

Import the Bucketeer SDK into your application code.

<Tabs>
<TabItem value="go" label="Go">

```go showLineNumbers
import (
	"github.com/bucketeer-io/go-server-sdk/pkg/bucketeer"
	"github.com/bucketeer-io/go-server-sdk/pkg/bucketeer/user"
)
```

</TabItem>
</Tabs>

### Initializing client

The SDK supports local and remote evaluations.

- [Local evaluation](#evaluating-users-within-sdk-locally): Evaluates end users within SDK locally
- [Remote evaluation](#remote-evaluation): Evaluate end users on the server

:::info

The **tag** setting is the tag you configure when creating a Feature Flag. It will evaluate all the Feature Flags in the environment when it is not configured.<br />
**We strongly recommend** using tags to speed up the evaluation process and reduce the response latency.

:::

#### Custom configuration

:::info Custom configuration

Depending on your use, you may want to change the optional configurations available.

- **eventQueueCapacity** (Default is 100)
- **numEventFlushWorkers** (Default is 50)
- **eventFlushInterval** (Default is 1 minute)
- **eventFlushSize** (Default is 100)
- **enableDebugLog** (Default is false)
- **errorLogger** (Default is `log.DefaultErrorLogger`)
- **enableLocalEvaluation** (Default is false)
- **cachePollingInterval** (Default is 1 minute)

For more information, please check the Option implementation [here](https://github.com/bucketeer-io/go-server-sdk/blob/master/pkg/bucketeer/option.go).

:::

#### Remote evaluation

To evaluate users on the server side you must create an API Key using the `Client SDK` role.

<Tabs>
<TabItem value="go" label="Go">

```go showLineNumbers
ctx, cancel := context.WithTimeout(context.Background(), timeout)
defer cancel()
client, err := bucketeer.NewSDK(
  ctx,
  bucketeer.WithAPIKey("YOUR_API_KEY"),
  bucketeer.WithAPIEndpoint("YOUR_API_ENDPOINT"),
  bucketeer.WithTag("YOUR_FEATURE_TAG"),
)
if err != nil {
  log.Fatalf("Failed initialize the new client: %v", err)
}
```

</TabItem>
</Tabs>

Once the SDK is configured, please check this [section](#evaluating-user) to learn how to get the variation for a user.

#### Evaluating users within SDK locally

By evaluating users locally you can improve response time significantly.<br />
To evaluate them you must create an API Key using the `Server SDK` role.

The SDK will poll the Feature Flags and Segment Users from the server, and cache them in memory.

:::caution

The Server SDK API Key has access to all Feature Flags and Segment Users in the environment.<br />
Keep in mind that it might contain sensitive information, so be careful when sharing the key with others.

:::

When initializing the SDK you must enable the local evaluation setting.

<Tabs>
<TabItem value="go" label="Go">

```go showLineNumbers
ctx, cancel := context.WithTimeout(context.Background(), timeout)
defer cancel()
client, err := bucketeer.NewSDK(
  ctx,
  bucketeer.WithAPIKey("YOUR_API_KEY"),
  bucketeer.WithAPIEndpoint("YOUR_API_ENDPOINT"),
  bucketeer.WithTag("YOUR_FEATURE_TAG"),
  bucketeer.WithEnableLocalEvaluation(true), // <--- Enable the local evaluation
  bucketeer.WithCachePollingInterval(10*time.Minute), // <--- Change the default interval if needed
)
if err != nil {
  log.Fatalf("Failed initialize the new client: %v", err)
}
```

Once the SDK is configured, please check this [section](#evaluating-user) to learn how to get the variation for a user.

</TabItem>
</Tabs>

## Supported features

### Evaluating user

The variation method determines whether or not a feature flag is enabled for a specific user.<br />
To check which variation a specific user will receive, you can use the client like below.

<Tabs>
<TabItem value="go" label="Go">

```go showLineNumbers
user := user.NewUser(
  "END_USER_ID",
  nil, // The user attributes are optional
)
showNewFeature := client.BoolVariation(ctx, user, "YOUR_FEATURE_FLAG_ID", false)
if (showNewFeature) {
  // The Application code to show the new feature
} else {
  // The code to run when the feature is off
}
```

</TabItem>
</Tabs>

:::note

The variation method will return the default value if the feature flag doesn't exist, or is archived, or if the request fails.

:::

### Variation types

The Bucketeer SDK supports the following variation types.

<Tabs>
<TabItem value="go" label="Go">

```go showLineNumbers
BoolVariation(ctx context.Context, user *user.User, featureID string, defaultValue bool) bool

IntVariation(ctx context.Context, user *user.User, featureID string, defaultValue int) int

Int64Variation(ctx context.Context, user *user.User, featureID string, defaultValue int64) int64

Float64Variation(ctx context.Context, user *user.User, featureID string, defaultValue float64) float64

StringVariation(ctx context.Context, user *user.User, featureID, defaultValue string) string

JSONVariation(ctx context.Context, user *user.User, featureID string, dst interface{})
```

</TabItem>
</Tabs>

### Getting evaluation details

The following methods will return the **evaluation details** for a specific feature flag. If the feature flag is missing in the SDK's cache, the variable `reason` value will be `CLIENT`, which means the default value was returned.

This is useful if you use another A/B Test solution with Bucketeer and need to know the variation name, reason, and other information.

#### Interfaces

<Tabs>
<TabItem value="go" label="Go">

```go showLineNumbers
type SDK interface {
    BoolVariationDetails(
      ctx context.Context,
      user *user.User,
      featureID string,
      defaultValue bool) model.BKTEvaluationDetails[bool]

    IntVariationDetails(
      ctx context.Context,
      user *user.User,
      featureID string,
      defaultValue int) model.BKTEvaluationDetails[int]

    Int64VariationDetails(
      ctx context.Context,
      user *user.User,
      featureID string,
      defaultValue int64) model.BKTEvaluationDetails[int64]

    Float64VariationDetails(
      ctx context.Context,
      user *user.User,
      featureID string,
      defaultValue float64) model.BKTEvaluationDetails[float64]

    StringVariationDetails(
      ctx context.Context,
      user *user.User,
      featureID,
      defaultValue string) model.BKTEvaluationDetails[string]

    ObjectVariationDetails(
      ctx context.Context,
      user *user.User,
      featureID string,
      defaultValue interface{}) model.BKTEvaluationDetails[interface{}]
}
```

</TabItem>
</Tabs>

#### Object

<Tabs>
<TabItem value="go" label="Go">

```go showLineNumbers
type BKTEvaluationDetails[T EvaluationValue] struct {
    FeatureID      string           // The ID of the feature flag.
    FeatureVersion int32            // The version of the feature flag.
    UserID         string           // The ID of the user being evaluated.
    VariationID    string           // The ID of the assigned variation.
    VariationName  string           // The name of the assigned variation.
    VariationValue T                // The value of the assigned variation.
    Reason         EvaluationReason // The reason for the evaluation result.
}

// EvaluationValue defines the allowed types for variation values
type EvaluationValue interface {
    int | int64 | float64 | string | bool | interface{}
}

type EvaluationReason string

const (
    EvaluationReasonTarget       EvaluationReason = "TARGET"        // Evaluated using individual targeting.
    EvaluationReasonRule         EvaluationReason = "RULE"          // Evaluated using a custom rule.
    EvaluationReasonDefault      EvaluationReason = "DEFAULT"       // Evaluated using the default strategy.
    EvaluationReasonClient       EvaluationReason = "CLIENT"        // The flag is missing in the cache; the default value was returned.
    EvaluationReasonOffVariation EvaluationReason = "OFF_VARIATION" // Evaluated using the off variation.
    EvaluationReasonPrerequisite EvaluationReason = "PREREQUISITE"  // Evaluated using a prerequisite.
)
```

</TabItem>
</Tabs>

#### Usage

<Tabs>
<TabItem value="go" label="Go">

```go showLineNumbers
user := user.NewUser(
  "END_USER_ID",
  nil, // The user attributes are optional
)
details := client.BoolVariationDetails(ctx, user, "YOUR_FEATURE_FLAG_ID", false)
if details.VariationValue {
  // The Application code to show the new feature
} else {
  // The code to run when the feature is off
}
```

</TabItem>
</Tabs>

### Reporting custom events

This method lets you save user actions in your application as events. You can connect these events to metrics in the experiments or in the kill switch (auto operation) on the console UI.

In addition, you can pass a number value to the goal event. These values will sum and show as <br />
`Value total` on the experiments console UI. This is useful if you have a goal event for tracking how much a user spent on your application buying items, etc.

<Tabs>
<TabItem value="go" label="Go">

```go showLineNumbers
user := user.NewUser(
  "END_USER_ID",
  nil, // The user attributes are optional
)
client.Track(ctx, user, "YOUR_GOAL_ID")

// Reporting with value
client.TrackValue(ctx, user, "YOUR_GOAL_ID", 10.50)
```

</TabItem>
</Tabs>

### User attributes configuration

This feature will give you robust and granular control over what users can see on your application.<br />
You can add rules using these attributes on the console UI's feature flag's targeting tab. [See more](/feature-flags/creating-feature-flags/targeting#user-attributes).

<Tabs>
<TabItem value="go" label="Go">

```go showLineNumbers
attributes := make(map[string]string)
attributes["app_version"] = "1.0.0"
attributes["os_version"] = "11.0.0"
attributes["device_model"] = "pixel-5"
attributes["language"] = "english"
attributes["genre"] = "female"

user := user.NewUser(
  "END_USER_ID",
  attributes,
)
showNewFeature := client.BoolVariation(ctx, user, "YOUR_FEATURE_FLAG_ID", false)
```

</TabItem>
</Tabs>

### Closing client

This method will send all pending analytics events to the Bucketeer server as soon as possible and tear down all the SDK activities and resources.<br />
The application should call this before the application stops. Otherwise, the remaining events can be lost.

<Tabs>
<TabItem value="go" label="Go">

```go showLineNumbers
if err := client.Close(ctx); err != nil {
	return fmt.Errorf("Failed to close the client: %v", err)
}
```

</TabItem>
</Tabs>
