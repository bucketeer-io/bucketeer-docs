---
title: React
slug: /sdk/client-side/react
toc_max_heading_level: 4
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

This category contains topics explaining how to configure Bucketeer's React SDK.

:::warning React SDK Version (Beta)

Bucketeer React SDK is a beta version. Breaking changes may be introduced before general release.

:::

## Features

- 🚀 React Context and Hooks for easy integration
- 🔧 TypeScript support with full type safety
- 🔄 Automatic re-rendering on flag changes

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

### 1. Configure and Initialize Client

First, import the necessary components and configure the SDK:

<Tabs>
<TabItem value="jsx" label="Importing">

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
```

</TabItem>
</Tabs>

Configure the SDK and user settings:

<Tabs>
<TabItem value="jsx" label="Configuration">

```jsx showLineNumbers
const config = defineBKTConfigForReact({
  apiKey: 'YOUR_API_KEY',
  apiEndpoint: 'YOUR_API_ENDPOINT',
  appVersion: '1.0.0',
  featureTag: 'web', // Optional but recommended
});

const user = defineBKTUser({
  id: 'USER_ID',
  customAttributes: {},
});
```

</TabItem>
</Tabs>

:::warning Important

Use `defineBKTConfigForReact` instead of `defineBKTConfig` when working with React SDK. This ensures proper React-specific configuration.

:::

Initialize the client in your root component:

<Tabs>
<TabItem value="jsx" label="Initialization">

```jsx showLineNumbers
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
          console.error('Failed to initialize Bucketeer:', error);
          // Keep client as null - app will use default values
          setClient(null);
        }
      }
    };

    init();
    return () => destroyBKTClient();
  }, []);

  return (
    <BucketeerProvider client={client}>
      {!client ? (
        <div>Loading Bucketeer...</div>
      ) : (
        <YourAppContent />
      )}
    </BucketeerProvider>
  );
}
```

</TabItem>
</Tabs>

:::info Default timeout

The initialize process default timeout is **5 seconds**.<br />
Once initialization is finished, all the requests in the SDK use a timeout of **30 seconds**.

:::

#### Custom timeout

You can customize the initialization timeout if needed:

<Tabs>
<TabItem value="jsx" label="Custom Timeout">

```jsx showLineNumbers
const timeout = 2000; // Custom timeout: 2 seconds (in milliseconds)
const initialFetchPromise = initializeBKTClient(config, user, timeout);

initialFetchPromise
  .then(() => {
    const client = getBKTClient();
    setClient(client);
    console.log('Bucketeer client initialized successfully');
  })
  .catch((error) => {
    if (error.name === 'TimeoutException') {
      // Client is still usable despite timeout
      const client = getBKTClient();
      setClient(client);
      console.warn('Initialization timed out, but client is ready');
    } else {
      console.error('Failed to initialize:', error);
    }
  });
```

</TabItem>
</Tabs>

:::info Initialization Timeout Error

During the initialization process, timeout errors are **not** related to the initialization itself. They arise from a timeout request, indicating the variations data from the server weren't received. The SDK will work as usual and update the variations in the next polling request.

:::

### 2. Setup BucketeerProvider

The `BucketeerProvider` is a React Context Provider that makes the Bucketeer client available to all child components through React hooks.

<Tabs>
<TabItem value="jsx" label="Provider Setup">

```jsx showLineNumbers
import { BucketeerProvider } from '@bucketeer/react-client-sdk';

// Wrap your app with BucketeerProvider
function App() {
  const [client, setClient] = useState(null);
  
  // ... initialization code ...

  return (
    <BucketeerProvider client={client}>
      <Header />
      <MainContent />
      <Footer />
    </BucketeerProvider>
  );
}
```

</TabItem>
</Tabs>

The `BucketeerProvider` provides:
- **client**: The initialized Bucketeer client instance
- **lastUpdated**: Timestamp of the last evaluation update (for advanced use cases)

:::info Graceful Degradation

If Bucketeer initialization fails, you can pass `client={null}` to `BucketeerProvider`. All hooks will automatically return their default values, allowing your app to continue working normally without feature flags.

:::

:::note

All Bucketeer hooks (`useBooleanVariation`, `useStringVariation`, etc.) must be used within components that are wrapped by `BucketeerProvider`.

:::

:::info Configuration Options

The **featureTag** setting filters which feature flags to evaluate. We strongly recommend using tags to improve performance and reduce cache size.

Optional configurations:
- **pollingInterval** - Default is 10 minutes (minimum 60 seconds)
- **eventsFlushInterval** - Default is 30 seconds  
- **eventsMaxQueueSize** - Default is 50 events
- **storageKeyPrefix** - Default is empty

:::

### 3. Use Feature Flag Hooks

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

These hooks return the current value of a feature flag and automatically re-render your component when the flag value changes.

---

#### useBooleanVariation

```typescript
useBooleanVariation(flagId: string, defaultValue: boolean): boolean
```

Returns a boolean feature flag value.

**Parameters:**
- `flagId` (string) - The feature flag identifier
- `defaultValue` (boolean) - Default value if flag is not available

**Returns:** `boolean`

<Tabs>
<TabItem value="jsx" label="Example">

```jsx showLineNumbers
import { useBooleanVariation } from '@bucketeer/react-client-sdk';

function FeatureComponent() {
  const showNewUI = useBooleanVariation('new-ui-enabled', false);
  
  return (
    <div>
      {showNewUI ? (
        <NewUserInterface />
      ) : (
        <LegacyUserInterface />
      )}
    </div>
  );
}
```

</TabItem>
</Tabs>

---

#### useStringVariation

```typescript
useStringVariation(flagId: string, defaultValue: string): string
```

Returns a string feature flag value.

**Parameters:**
- `flagId` (string) - The feature flag identifier
- `defaultValue` (string) - Default value if flag is not available

**Returns:** `string`

<Tabs>
<TabItem value="jsx" label="Example">

```jsx showLineNumbers
import { useStringVariation } from '@bucketeer/react-client-sdk';

function ThemeComponent() {
  const theme = useStringVariation('app-theme', 'light');
  
  return (
    <div className={`theme-${theme}`}>
      <h1>Current theme: {theme}</h1>
      {/* Your themed content */}
    </div>
  );
}
```

</TabItem>
</Tabs>

---

#### useNumberVariation

```typescript
useNumberVariation(flagId: string, defaultValue: number): number
```

Returns a number feature flag value.

**Parameters:**
- `flagId` (string) - The feature flag identifier
- `defaultValue` (number) - Default value if flag is not available

**Returns:** `number`

<Tabs>
<TabItem value="jsx" label="Example">

```jsx showLineNumbers
import { useNumberVariation } from '@bucketeer/react-client-sdk';

function ProductList() {
  const itemsPerPage = useNumberVariation('items-per-page', 20);
  const maxPrice = useNumberVariation('max-price-filter', 1000);
  
  return (
    <div>
      <p>Showing {itemsPerPage} items per page</p>
      <p>Max price filter: ${maxPrice}</p>
      {/* Product listing logic */}
    </div>
  );
}
```

</TabItem>
</Tabs>

---

#### useObjectVariation

```typescript
useObjectVariation<T>(flagId: string, defaultValue: T): T
```

Returns a JSON/object feature flag value with type safety.

**Parameters:**
- `flagId` (string) - The feature flag identifier
- `defaultValue` (T) - Default value if flag is not available

**Returns:** `T` (the type must extend `BKTValue`)

<Tabs>
<TabItem value="jsx" label="Example">

```jsx showLineNumbers
import { useObjectVariation } from '@bucketeer/react-client-sdk';

function ConfigComponent() {
  const apiConfig = useObjectVariation('api-config', {
    timeout: 5000,
    retries: 3,
    baseUrl: 'https://api.example.com'
  });
  
  const features = useObjectVariation('enabled-features', [
    'chat', 'notifications', 'analytics'
  ]);
  
  return (
    <div>
      <p>API Timeout: {apiConfig.timeout}ms</p>
      <p>Retries: {apiConfig.retries}</p>
      <p>Enabled features: {features.join(', ')}</p>
    </div>
  );
}
```

</TabItem>
</Tabs>

---

### Evaluation Details Hooks

These hooks return both the feature flag value and detailed evaluation information, useful for debugging or analytics.

---

#### useBooleanVariationDetails

```typescript
useBooleanVariationDetails(flagId: string, defaultValue: boolean): BKTEvaluationDetails<boolean>
```

Returns a boolean feature flag value with detailed evaluation information.

**Parameters:**
- `flagId` (string) - The feature flag identifier
- `defaultValue` (boolean) - Default value if flag is not available

**Returns:** `BKTEvaluationDetails<boolean>`

---

#### useStringVariationDetails

```typescript
useStringVariationDetails(flagId: string, defaultValue: string): BKTEvaluationDetails<string>
```

Returns a string feature flag value with detailed evaluation information.

**Parameters:**
- `flagId` (string) - The feature flag identifier
- `defaultValue` (string) - Default value if flag is not available

**Returns:** `BKTEvaluationDetails<string>`

---

#### useNumberVariationDetails

```typescript
useNumberVariationDetails(flagId: string, defaultValue: number): BKTEvaluationDetails<number>
```

Returns a number feature flag value with detailed evaluation information.

**Parameters:**
- `flagId` (string) - The feature flag identifier
- `defaultValue` (number) - Default value if flag is not available

**Returns:** `BKTEvaluationDetails<number>`

---

#### useObjectVariationDetails

```typescript
useObjectVariationDetails<T>(flagId: string, defaultValue: T): BKTEvaluationDetails<T>
```

Returns a JSON/object feature flag value with detailed evaluation information.

**Parameters:**
- `flagId` (string) - The feature flag identifier
- `defaultValue` (T) - Default value if flag is not available

**Returns:** `BKTEvaluationDetails<T>`

<Tabs>
<TabItem value="jsx" label="Detailed Evaluation Example">

```jsx showLineNumbers
import { useBooleanVariationDetails } from '@bucketeer/react-client-sdk';

function AdvancedComponent() {
  const featureDetails = useBooleanVariationDetails('advanced-feature', false);
  
  // Log evaluation details for debugging
  console.log('Feature evaluation:', {
    featureId: featureDetails.featureId,
    variationId: featureDetails.variationId,
    variationName: featureDetails.variationName,
    reason: featureDetails.reason,
    userId: featureDetails.userId
  });
  
  return (
    <div>
      <p>Feature enabled: {featureDetails.variationValue}</p>
      <p>Variation: {featureDetails.variationName}</p>
      <p>Evaluation reason: {featureDetails.reason}</p>
      
      {featureDetails.variationValue && <AdvancedFeature />}
      
      {/* Show debug info in development */}
      {process.env.NODE_ENV === 'development' && (
        <details>
          <summary>Debug Info</summary>
          <pre>{JSON.stringify(featureDetails, null, 2)}</pre>
        </details>
      )}
    </div>
  );
}
```

</TabItem>
</Tabs>

---

#### Evaluation Details Object

The `BKTEvaluationDetails<T>` object contains detailed information about feature flag evaluation. For complete details about this interface, see the [JavaScript SDK Evaluation Details documentation](/sdk/client-side/javascript#object).

:::note Hook Behavior

- All hooks automatically **re-render** your component when flag values change
- Hooks **must** be used within components wrapped by `BucketeerProvider`
- If a flag is not found, hooks return the provided `defaultValue`
- Hooks are **reactive** - they respond to polling updates and manual fetches

:::

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

## Re-exported from JavaScript SDK

The React SDK re-exports all functionality from the JavaScript SDK, allowing you to use any JavaScript SDK features directly:

This means you can access all JavaScript SDK functionality without needing to install the JavaScript SDK separately. You can use features like:

- Direct client methods (`client.booleanVariation()`, `client.track()`, etc.)
- Advanced configuration options
- Error handling utilities
- All TypeScript types and interfaces

:::tip

When using React components, prefer the React hooks (ie: `useBooleanVariation`) over direct client methods for better integration with React's rendering cycle. Use direct client methods only when necessary (e.g., in event handlers or outside React components).

:::

For complete API reference and advanced features, see the [JavaScript SDK documentation](/sdk/client-side/javascript).
