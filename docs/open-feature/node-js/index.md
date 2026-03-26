---
title: Node.js
slug: /open-feature/node-js
---

# Node.js Server Provider

:::warning Node.js Provider (Beta)

This is a beta version. Breaking changes may be introduced before general release.

:::

This is the official Node.js OpenFeature server-side provider for [Bucketeer](https://bucketeer.io/).

### Installation

```bash showLineNumbers
npm install @bucketeer/openfeature-node-server-sdk @openfeature/server-sdk @bucketeer/node-server-sdk @openfeature/core
```

**Note:** This package requires `@openfeature/server-sdk` and `@bucketeer/node-server-sdk` as peer dependencies. `@openfeature/server-sdk` requires `@openfeature/core` as a peer dependency.

### Usage

#### Initialize the provider

The Bucketeer provider needs to be created and then set in the global OpenFeature instance.

```typescript showLineNumbers
import { OpenFeature } from '@openfeature/server-sdk';
import { BucketeerProvider, defineBKTConfig } from '@bucketeer/openfeature-node-server-sdk';

const config = defineBKTConfig({
  apiKey: 'BUCKETEER_API_KEY',
  apiEndpoint: 'BUCKETEER_API_ENDPOINT',
  featureTag: 'FEATURE_TAG',
  appVersion: '1.2.3',
});

// Initialize the provider
const provider = new BucketeerProvider(config);

// Set the provider and wait for initialization
await OpenFeature.setProviderAndWait(provider);
```

See our [documentation](https://docs.bucketeer.io/sdk/server-side/node-js) for more SDK configuration options.

#### Evaluation Context

In server-side applications, evaluation context is typically provided per request rather than set globally. This allows for user-specific flag evaluations based on request data.

The `targetingKey` is the user ID (Unique ID) and cannot be empty.

```typescript showLineNumbers
const client = OpenFeature.getClient();

// Define evaluation context per request/user
const evaluationContext = {
  targetingKey: 'user-123', // Required: unique user identifier
  email: 'user@example.com', // User attributes for targeting
  plan: 'premium',
  region: 'us-east-1',
};

// Evaluate flags with context
const featureEnabled = await client.getBooleanValue('new-feature', false, evaluationContext);
```

#### Evaluate a feature flag

After the provider is set and ready, you can evaluate feature flags using the OpenFeature client. Always provide evaluation context for each evaluation.

```typescript showLineNumbers
const client = OpenFeature.getClient();

const context = {
  targetingKey: 'user-123',
  email: 'user@example.com',
};

// boolean flag
const booleanValue = await client.getBooleanValue('my-feature-flag', false, context);

// string flag
const stringValue = await client.getStringValue('my-feature-flag', 'default-value', context);

// number flag
const numberValue = await client.getNumberValue('my-feature-flag', 0, context);

// object flag
const objectValue = await client.getObjectValue('my-feature-flag', {}, context);
```
