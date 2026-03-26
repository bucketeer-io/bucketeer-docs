---
title: iOS
slug: /open-feature/ios
---

# iOS Provider

:::warning Swift Provider (Beta)

This is a beta version. Breaking changes may be introduced before general release.

:::

This is the official Swift OpenFeature provider for [Bucketeer](https://bucketeer.io/), designed for iOS and tvOS applications.

### Installation

#### Swift Package Manager

1. In Xcode, select **File > Swift Packages > Add Package Dependency**.
2. Enter the repository URL: `https://github.com/bucketeer-io/openfeature-swift-client-sdk.git`.
3. Select the product **BucketeerOpenFeature** and add it to your app target.

**Requirements:** Xcode 16.0+, Swift 5.0+, iOS 14.0+, tvOS 14.0+

### Usage

#### Initialize the provider

The Bucketeer provider needs to be created and then set in the global `OpenFeatureAPI`.

```swift showLineNumbers
import BucketeerOpenFeature
import OpenFeature
import Bucketeer

do {
    let config = try BKTConfig.Builder()
        .with(apiKey: "YOUR_API_KEY")
        .with(apiEndpoint: "YOUR_API_ENDPOINT")
        .with(featureTag: "YOUR_FEATURE_TAG")
        .with(appVersion: "1.2.3")
        .build()
    
    let provider = BucketeerProvider(config: config)
    
    let context = MutableContext(
        targetingKey: "user-123",
        structure: MutableStructure(attributes: [:])
    )

    await OpenFeatureAPI.shared.setProviderAndWait(provider: provider, initialContext: context)
} catch {
    // Error handling
}
```

See our [documentation](https://docs.bucketeer.io/sdk/client-side/ios) for more SDK configuration.

#### Update the Evaluation Context

You can update the evaluation context with the new attributes if the user attributes change.

```swift showLineNumbers
let ctx = MutableContext(targetingKey: "user-123", structure: MutableStructure(
    attributes: ["buyer": "true"]
))
OpenFeatureAPI.shared.setEvaluationContext(evaluationContext: ctx)
```

:::warning
Changing the `targetingKey` is not supported in the current implementation of the BucketeerProvider. To change the user ID, the BucketeerProvider must be removed and reinitialized.
:::

#### Evaluate a feature flag

```swift showLineNumbers
let client = OpenFeatureAPI.shared.getClient()

// Bool
client.getBooleanValue(key: "my-flag", defaultValue: false)

// String
client.getStringValue(key: "my-flag", defaultValue: "default")

// Integer
client.getIntegerValue(key: "my-flag", defaultValue: 1)

// Double
client.getDoubleValue(key: "my-flag", defaultValue: 1.1)
```
