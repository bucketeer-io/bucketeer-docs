---
title: Android
slug: /open-feature/android
---

# Android Provider

:::warning Kotlin Provider (Beta)

This is a beta version. Breaking changes may be introduced before general release.

:::

This is the official Kotlin OpenFeature provider for [Bucketeer](https://bucketeer.io/), designed for Android clients.

### Installation

Add the dependency to your `build.gradle` file:

```groovy showLineNumbers
dependencies {
    implementation 'io.bucketeer:openfeature-kotlin-client-sdk:LATEST_VERSION'
}
```

### Usage

#### Initialize the provider

The Bucketeer provider needs to be created and then set in the global `OpenFeatureAPI`.

```kotlin showLineNumbers
import dev.openfeature.sdk.*
import io.bucketeer.openfeatureprovider.BucketeerProvider
import io.bucketeer.sdk.android.*
import kotlinx.coroutines.*

val config = BKTConfig.builder()
    .apiKey("API_KEY")
    .apiEndpoint("API_ENDPOINT")
    .featureTag("android")
    .appVersion(BuildConfig.VERSION_NAME)
    .build()

// Evaluation context
val initContext = ImmutableContext(
    targetingKey = "USER_ID",
    attributes = mapOf("attr1" to Value.String("value1"))
)

val provider = BucketeerProvider(context, config, lifecycleScope)

lifecycleScope.launch {
    OpenFeatureAPI.setProviderAndWait(provider, Dispatchers.IO, initContext)
    
    val client = OpenFeatureAPI.getClient()
    val flag = client.getBooleanValue("feature-id", defaultValue = false)
}
```

Note: `lifecycleScope` should be the activity or fragment lifecycle scope.

See our [documentation](https://docs.bucketeer.io/sdk/client-side/android) for more SDK configuration.

#### Update the Evaluation Context

You can update the evaluation context with the new attributes if the user attributes change.

```kotlin showLineNumbers
val newContext = ImmutableContext(
    targetingKey = "USER_ID",
    attributes = mapOf("attr2" to Value.String("value2"))
)
OpenFeatureAPI.setEvaluationContext(newContext)
```

:::warning
Changing the `targetingKey` is not supported in the current implementation of the BucketeerProvider. To change the user ID, the BucketeerProvider must be removed and reinitialized.
:::

#### Evaluate a feature flag

```kotlin showLineNumbers
val client = OpenFeatureAPI.getClient()

// Bool
client.getBooleanValue("my-flag", defaultValue = false)

// String
client.getStringValue("my-flag", defaultValue = "default")

// Integer
client.getIntegerValue("my-flag", defaultValue = 1)

// Double
client.getDoubleValue("my-flag", defaultValue = 1.1)
```
