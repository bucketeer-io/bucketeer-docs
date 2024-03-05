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

Initializing the client.

:::info

The **tag** setting is the tag you configure when creating a Feature Flag. It will evaluate all the Feature Flags in the environment when it is not configured.<br />
**We strongly recommend** using tags to speed up the evaluation process and reduce the response latency.

:::
<Tabs>
<TabItem value="go" label="Go">

```go showLineNumbers
ctx, cancel := context.WithTimeout(context.Background(), timeout)
defer cancel()
client, err := bucketeer.NewSDK(
  ctx,
  bucketeer.WithAPIKey("YOUR_API_KEY"),
  bucketeer.WithHost("YOUR_API_ENDPOINT"),
  bucketeer.WithTag("YOUR_FEATURE_TAG"),
)
if err != nil {
  log.Fatalf("Failed initialize the new client: %v", err)
}
```

</TabItem>
</Tabs>

:::info Custom configuration

Depending on your use, you may want to change the optional configurations available.

- **eventQueueCapacity** (Default is 100)
- **numEventFlushWorkers** (Default is 50)
- **eventFlushInterval** (Default is 1 minute)
- **eventFlushSize** (Default is 100)
- **enableDebugLog** (Default is false)
- **errorLogger** (Default is `log.DefaultErrorLogger`)

For more information, please check the Option implementation [here](https://github.com/bucketeer-io/go-server-sdk/blob/master/pkg/bucketeer/option.go).

:::

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
