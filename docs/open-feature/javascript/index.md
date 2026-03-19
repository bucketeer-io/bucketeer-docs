---
title: JavaScript
slug: /open-feature/javascript
---

# Web (JavaScript / TypeScript)

:::warning Web Provider (Beta)

This is a beta version. Breaking changes may be introduced before general release.

:::

This is the official JavaScript OpenFeature provider for [Bucketeer](https://bucketeer.io/).

### Installation

```bash showLineNumbers
npm install @bucketeer/openfeature-js-client-sdk
```

:::note
npm versions 7 and above will automatically install the required peer dependencies: 
`@openfeature/web-sdk` and `@bucketeer/js-client-sdk`.
If you got an error about missing peer dependencies, please install them manually:

```bash showLineNumbers
npm install @openfeature/core @openfeature/web-sdk @bucketeer/js-client-sdk
```
:::

### Usage

#### Configuration

Use `defineBKTConfig` to create your configuration.

```js showLineNumbers
import { defineBKTConfig } from '@bucketeer/openfeature-js-client-sdk';

const config = defineBKTConfig({
  apiEndpoint: 'BUCKETEER_API_ENDPOINT',
  apiKey: 'BUCKETEER_API_KEY',
  featureTag: 'FEATURE_TAG',
  appVersion: '1.2.3',
  fetch: window.fetch,
})
```

See our [documentation](https://docs.bucketeer.io/sdk/client-side/javascript#configuring-client) for more SDK configuration.

#### Initialization

Initialize and set the Bucketeer provider to OpenFeature.

```js showLineNumbers
import { OpenFeature } from '@openfeature/web-sdk';
import { defineBKTConfig, BucketeerProvider } from '@bucketeer/openfeature-js-client-sdk';

const config = defineBKTConfig({
  apiEndpoint: 'BUCKETEER_API_ENDPOINT',
  apiKey: 'BUCKETEER_API_KEY',
  featureTag: 'FEATURE_TAG',
  appVersion: '1.2.3',
  fetch: window.fetch,
})

const initEvaluationContext = {
  targetingKey: 'USER_ID',
  app_version: '1.2.3',
}
await OpenFeature.setContext(initEvaluationContext)
const provider = new BucketeerProvider(config)
await OpenFeature.setProviderAndWait(provider)
```

#### Evaluate a feature flag

After the provider is set and the provider's status is `ClientProviderEvents.Ready`, you can evaluate a feature flag using the OpenFeature client.

```js showLineNumbers
const client = OpenFeature.getClient();

// boolean flag
const flagValueBool = client.getBooleanValue('my-feature-flag', false);

// string flag
const flagValueStr = client.getStringValue('my-feature-flag', 'default-value');

// number flag
const flagValueNum = client.getNumberValue('my-number-flag', 0);

// object flag
const flagValueObj = client.getObjectValue('my-object-flag', {});
```

More details can be found in the [OpenFeature Web SDK documentation](https://openfeature.dev/docs/reference/sdks/client/web/#usage).

#### Update the Evaluation Context

The evaluation context allows the client to specify contextual data that Bucketeer uses to evaluate feature flags.
The `targetingKey` is the user ID (Unique ID) and cannot be empty.

You can update the evaluation context with the new attributes if the user attributes change.

```js showLineNumbers
const newEvaluationContext = {
  targetingKey: 'USER_ID',
  app_version: '2.0.0',
  age: 25,
  country: 'US',
}
await OpenFeature.setContext(newEvaluationContext)
```

:::warning
Changing the `targetingKey` is not supported in the current implementation of the BucketeerProvider.
:::

To change the user ID, the BucketeerProvider must be removed and reinitialized.

```js showLineNumbers
await OpenFeature.clearProviders()
await OpenFeature.clearContext()

// Reinitialize the provider with new targetingKey
const newEvaluationContext = {
  targetingKey: 'USER_ID_NEW',
  app_version: '2.0.0',
  age: 25,
  country: 'US',
}

const config = defineBKTConfig({
  apiEndpoint: 'BUCKETEER_API_ENDPOINT',
  apiKey: 'BUCKETEER_API_KEY',
  featureTag: 'FEATURE_TAG',
  appVersion: '1.2.3',
  fetch: window.fetch,
})

await OpenFeature.setContext(newEvaluationContext)
const provider = new BucketeerProvider(config)
await OpenFeature.setProviderAndWait(provider)
```
