---
title: React Native
slug: /open-feature/react-native
---

# React Native

:::warning React Native Provider (Beta)

This is a beta version. Breaking changes may be introduced before general release.

:::

This is the official React Native OpenFeature provider for [Bucketeer](https://bucketeer.io/).

### Installation

```bash showLineNumbers
npm install @bucketeer/openfeature-js-client-sdk @openfeature/react-sdk
```

:::note
npm versions 7 and above will automatically install the required peer dependencies.
If you got an error about missing peer dependencies, please install them manually:

```bash showLineNumbers
npm install @openfeature/core @openfeature/web-sdk @bucketeer/js-client-sdk
```
:::

:::important
The Bucketeer React Native provider relies on `@react-native-async-storage/async-storage` for local caching and `react-native-uuid` for generating IDs for Bucketeer SDK events.

**Expo Users:**
Simply installing the Bucketeer OpenFeature JS provider is sufficient.
No additional steps are required for Android.
For iOS, navigate to the iOS folder and run `pod install`.
```bash showLineNumbers
cd ios && pod install  # For iOS
```

**Non-Expo Users:**
You must explicitly install `@react-native-async-storage/async-storage` and `react-native-uuid` as direct dependencies in your project. 
No additional steps are required for Android.
For iOS, navigate to the iOS folder and run `pod install`.
This is necessary because React Native's auto-linking feature does not support transitive dependencies (see [CLI Issue #1347](https://github.com/react-native-community/cli/issues/1347)).

```bash showLineNumbers
npm install @react-native-async-storage/async-storage react-native-uuid
cd ios && pod install  # For iOS
```

**Note on Optional Peer Dependencies:**
These dependencies are marked as **optional** peer dependencies in `package.json` to avoid forcing Web and Node.js consumers to install packages they don't need. However, `BucketeerReactNativeProvider` will throw a `ProviderFatalError` if `react-native-uuid` is missing at runtime, as it is strictly required for the React Native environment.

If `@react-native-async-storage/async-storage` is not installed, the SDK will gracefully fall back to in-memory storage.

For more details, see: https://react-native-async-storage.github.io/async-storage/docs/install/
:::

### Usage

Please use the [OpenFeature React SDK](https://openfeature.dev/docs/reference/sdks/client/web/react/) to use feature flags in your React Native application.

#### Configuration & Initialization

Use `defineBKTConfigForReactNative` to create your configuration and set up the `OpenFeatureProvider`. Make sure to use the global `fetch` API.

```js showLineNumbers
import { OpenFeatureProvider, OpenFeature } from '@openfeature/react-sdk';
import { defineBKTConfigForReactNative, BucketeerReactNativeProvider } from '@bucketeer/openfeature-js-client-sdk';

const config = defineBKTConfigForReactNative({
  apiEndpoint: 'BUCKETEER_API_ENDPOINT',
  apiKey: 'BUCKETEER_API_KEY',
  featureTag: 'FEATURE_TAG',
  appVersion: '1.2.3',
  fetch: fetch, // Use global fetch in React Native
})

const initEvaluationContext = {
  targetingKey: 'USER_ID',
  app_version: '1.2.3',
}
await OpenFeature.setContext(initEvaluationContext)
const provider = new BucketeerReactNativeProvider(config)
OpenFeature.setProvider(provider)

// Note: There is no need to await setProvider in React Native,
// because provider initialization is handled internally by the OpenFeature React SDK.

function App() {
  return (
    <OpenFeatureProvider>
      <YourApp />
    </OpenFeatureProvider>
  )
}
```

See our [documentation](https://docs.bucketeer.io/sdk/client-side/javascript#configuring-client) for more SDK configuration.

:::important
Use `defineBKTConfigForReactNative` for the standard setup — you do not need to provide an `idGenerator`. The `BucketeerReactNativeProvider` automatically loads and injects the correct React Native implementation (`react-native-uuid`) during initialization.
:::

#### Evaluate a feature flag

The OpenFeature React SDK provides hooks for evaluating feature flags.

```js showLineNumbers
import { useBooleanFlagValue, useStringFlagValue, useNumberFlagValue, useObjectFlagValue } from '@openfeature/react-sdk';

// boolean flag
const flagValueBool = useBooleanFlagValue('my-feature-flag', false);

// string flag
const flagValueStr = useStringFlagValue('my-feature-flag', 'default-value');

// number flag
const flagValueNum = useNumberFlagValue('my-number-flag', 0);

// object flag
const flagValueObj = useObjectFlagValue('my-object-flag', {});
```

More details can be found in the [OpenFeature React SDK documentation](https://openfeature.dev/docs/reference/sdks/client/web/react#usage).

#### Update the Evaluation Context

The evaluation context allows the client to specify contextual data that Bucketeer uses to evaluate feature flags.
The `targetingKey` is the user ID (Unique ID) and cannot be empty.

You can update the evaluation context with the new attributes if the user attributes change.

```js showLineNumbers
import { OpenFeature } from '@openfeature/react-sdk';

const newEvaluationContext = {
  targetingKey: 'USER_ID',
  app_version: '2.0.0',
  age: 25,
  country: 'US',
}
await OpenFeature.setContext(newEvaluationContext)
```

:::warning
Changing the `targetingKey` is not supported in the current implementation of the BucketeerReactNativeProvider. To change the user ID, the Provider must be removed and reinitialized exactly as demonstrated in the [JavaScript / Web section](/open-feature/javascript#update-the-evaluation-context).
:::
