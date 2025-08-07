---
title: React
slug: /sdk/client-side/react
toc_max_heading_level: 4
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

The Bucketeer React SDK provides React hooks and context for seamless feature flag integration in React applications.

:::warning React SDK Version (Beta)

Bucketeer React SDK is a beta version. Breaking changes may be introduced before general release.

:::

## Features

- 🚀 React Context and Hooks for easy integration
- 🔧 TypeScript support with full type safety
- ⚡ Real-time feature flag updates
- 🎯 Multiple variation types (boolean, string, number, object)
- 🧪 User attribute management

## Requirements

Before starting, ensure that you follow the [Getting Started](/getting-started) guide.

**React Version Support:**
- ✅ Supported: React 18.2.0 - 18.3.x
- ⚠️ May work: React 18.0.0 - 18.1.x (not officially supported)
- ❌ Not supported: React 19.0.0 and above

## Installation

<Tabs>
<TabItem value="npm" label="npm">

```sh showLineNumbers
npm install @bucketeer/react-client-sdk
```

</TabItem>
<TabItem value="yarn" label="Yarn">

```sh showLineNumbers
yarn add @bucketeer/react-client-sdk
```

</TabItem>
</Tabs>

## Setup

### 1. Initialize the Client

Configure and initialize the Bucketeer client in your app's root component:

<Tabs>
<TabItem value="jsx" label="App.jsx">

```jsx showLineNumbers
import React, { useEffect, useState } from 'react';
import {
  BucketeerProvider,
  defineBKTConfigForReact,
  defineBKTUser,
  initializeBKTClient,
  getBKTClient,
  destroyBKTClient,
} from '@bucketeer/react-client-sdk';

const config = defineBKTConfigForReact({
  apiKey: 'YOUR_API_KEY',
  apiEndpoint: 'YOUR_API_ENDPOINT',
  appVersion: '1.0.0',
  featureTag: 'web', // Optional but recommended
});

const user = defineBKTUser({
  id: 'user-123',
  customAttributes: {
    platform: 'web',
    version: '1.0.0',
  },
});

export default function App() {
  const [client, setClient] = useState(null);

  useEffect(() => {
    const init = async () => {
      try {
        await initializeBKTClient(config, user);
        const bktClient = getBKTClient();
        setClient(bktClient);
      } catch (error) {
        if (error.name === 'TimeoutException') {
          // Client is still initialized despite timeout
          console.warn('Initialization timed out, but client is ready');
          const bktClient = getBKTClient();
          setClient(bktClient);
        } else {
          console.error('Failed to initialize:', error);
        }
      }
    };

    init();
    return () => destroyBKTClient();
  }, []);

  if (!client) {
    return <div>Loading Bucketeer...</div>;
  }

  return (
    <BucketeerProvider client={client}>
      <YourAppContent />
    </BucketeerProvider>
  );
}
```

</TabItem>
</Tabs>

:::info Configuration Options

The **featureTag** setting filters which feature flags to evaluate. We strongly recommend using tags to improve performance and reduce cache size.

Optional configurations:
- **pollingInterval** - Default is 10 minutes (minimum 60 seconds)
- **eventsFlushInterval** - Default is 30 seconds  
- **eventsMaxQueueSize** - Default is 50 events
- **storageKeyPrefix** - Default is empty

:::

### 2. Use Feature Flag Hooks

Use the provided hooks to access feature flags in your components:

<Tabs>
<TabItem value="jsx" label="MyComponent.jsx">

```jsx showLineNumbers
import React from 'react';
import {
  useBooleanVariation,
  useStringVariation,
  useNumberVariation,
  useObjectVariation,
} from '@bucketeer/react-client-sdk';

function MyComponent() {
  // Boolean feature flag
  const showNewFeature = useBooleanVariation('show-new-feature', false);
  
  // String feature flag
  const theme = useStringVariation('app-theme', 'light');
  
  // Number feature flag
  const maxItems = useNumberVariation('max-items', 10);
  
  // Object feature flag
  const config = useObjectVariation('app-config', { timeout: 5000 });

  return (
    <div>
      {showNewFeature && <NewFeature />}
      <div>Theme: {theme}</div>
      <div>Max items: {maxItems}</div>
      <div>Timeout: {config.timeout}ms</div>
    </div>
  );
}
```

</TabItem>
</Tabs>

## Hook Reference

### Variation Hooks

These hooks return the current value of a feature flag:

- `useBooleanVariation(flagId, defaultValue)` - Returns boolean
- `useStringVariation(flagId, defaultValue)` - Returns string  
- `useNumberVariation(flagId, defaultValue)` - Returns number
- `useObjectVariation(flagId, defaultValue)` - Returns object/array

### Evaluation Details Hooks

These hooks return both the value and detailed evaluation information:

- `useBooleanVariationDetails(flagId, defaultValue)`
- `useStringVariationDetails(flagId, defaultValue)`  
- `useNumberVariationDetails(flagId, defaultValue)`
- `useObjectVariationDetails(flagId, defaultValue)`

<Tabs>
<TabItem value="jsx" label="Example">

```jsx showLineNumbers
import { useBooleanVariationDetails } from '@bucketeer/react-client-sdk';

function AdvancedComponent() {
  const featureDetails = useBooleanVariationDetails('advanced-feature', false);
  
  return (
    <div>
      <p>Feature enabled: {featureDetails.variationValue}</p>
      <p>Variation: {featureDetails.variationName}</p>
      <p>Reason: {featureDetails.reason}</p>
    </div>
  );
}
```

</TabItem>
</Tabs>

## Advanced Usage

### Updating User Attributes

<Tabs>
<TabItem value="jsx" label="UserProfile.jsx">

```jsx showLineNumbers
import { useContext } from 'react';
import { BucketeerContext } from '@bucketeer/react-client-sdk';

function UserProfile() {
  const { client } = useContext(BucketeerContext);

  const handleUpgrade = () => {
    client.updateUserAttributes({
      plan: 'premium',
      tier: 'gold',
    });
  };

  return (
    <button onClick={handleUpgrade}>
      Upgrade to Premium
    </button>
  );
}
```

</TabItem>
</Tabs>

### Tracking Events

<Tabs>
<TabItem value="jsx" label="EventTracking.jsx">

```jsx showLineNumbers
import { useContext } from 'react';
import { BucketeerContext } from '@bucketeer/react-client-sdk';

function PurchaseButton() {
  const { client } = useContext(BucketeerContext);

  const handlePurchase = (amount) => {
    // Track purchase event with value
    client.track('purchase_completed', amount);
  };

  return (
    <button onClick={() => handlePurchase(99.99)}>
      Buy Now
    </button>
  );
}
```

</TabItem>
</Tabs>

### Manual Evaluation Updates

<Tabs>
<TabItem value="jsx" label="RefreshButton.jsx">

```jsx showLineNumbers
import { useContext } from 'react';
import { BucketeerContext } from '@bucketeer/react-client-sdk';

function RefreshButton() {
  const { client } = useContext(BucketeerContext);

  const handleRefresh = async () => {
    try {
      await client.fetchEvaluations(5000); // 5 second timeout
      console.log('Evaluations updated');
    } catch (error) {
      console.error('Failed to update:', error);
    }
  };

  return (
    <button onClick={handleRefresh}>
      Refresh Feature Flags
    </button>
  );
}
```

</TabItem>
</Tabs>

## Best Practices

1. **Use feature tags** to improve performance by filtering relevant flags
2. **Handle initialization timeouts** gracefully - the client may still work
3. **Update user attributes** when user context changes (login, plan changes, etc.)
4. **Use TypeScript** for better type safety with the provided types
5. **Avoid frequent manual fetches** - let automatic polling handle updates

## TypeScript Support

The React SDK provides full TypeScript support with exported types:

```typescript
import type {
  BKTClient,
  BKTUser,
  BKTConfig,
  BKTEvaluationDetails,
  BKTValue,
} from '@bucketeer/react-client-sdk';
```

For complete API reference and advanced features, see the [JavaScript SDK documentation](/sdk/client-side/javascript).
