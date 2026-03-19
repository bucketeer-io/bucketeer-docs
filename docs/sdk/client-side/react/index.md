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

## Getting started

### Installing dependency

Install the dependency in your application.

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

### Importing client

Import the necessary components and configure the SDK:

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

### Configuring client

Configure the SDK config and user configuration.

:::info

The **featureTag** setting is the tag you configure when creating a Feature Flag. It will evaluate all the Feature Flags in the environment when it is not configured.

**We strongly recommend** using tags to speed up the evaluation process and reduce the cache size in the client.

:::

All the settings in the example below are required.

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

Use `defineBKTConfigForReact` instead of `defineBKTConfig` when working with React SDK. This function includes React-specific optimizations and ensures proper integration with React's rendering cycle.

:::

### Initializing client

Initialize the client in your root component:

<Tabs>
<TabItem value="jsx" label="Initialization">

```jsx showLineNumbers
export default function App() {
  const [client, setClient] = useState(null);
  const [isInitialized, setIsInitialized] = useState(false);

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
      } finally {
        setIsInitialized(true);
      }
    };

    init();
    return () => destroyBKTClient();
  }, []);

  if (!isInitialized) {
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

:::info Default timeout

The initialization process has a default timeout of **5 seconds**.<br />
Once initialization is finished, all the requests in the SDK use a timeout of **30 seconds**.

:::

#### Custom timeout

You can customize the initialization timeout if needed:

<Tabs>
<TabItem value="jsx" label="Custom Timeout">

```jsx showLineNumbers
const timeout = 2000; // Default is 5 seconds (In milliseconds)
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
      console.error('Failed to initialize with BKTException:', error);
    }
  });
```

</TabItem>
</Tabs>

:::info Initialization Timeout Error

During the initialization process, timeout errors are **not** related to the initialization itself. They arise from a timeout request, indicating the variations data from the server weren't received. The SDK will work as usual and update the variations in the next polling request.

:::

### Setup BucketeerProvider

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

:::info Custom configuration

Depending on your use, you may want to change the optional configurations available.

- **pollingInterval** - Minimum 60 seconds. Default is 10 minutes (In Milliseconds)
- **eventsFlushInterval** - Default is 30 seconds (In Milliseconds)
- **eventsMaxQueueSize** - Default is 50 events
- **storageKeyPrefix** - Default is empty
- **userAgent** - Default is `window.navigator.userAgent`
- **fetch** - Default is `globalThis.fetch`

:::

:::note

The Bucketeer SDK doesn't save the user data. The Application must save and set it when initializing the client SDK.

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

### Evaluating user

The variation hooks determine whether or not a feature flag is enabled for a specific user.<br />
To check which variation a specific user will receive, you can use the hooks like below.

```jsx showLineNumbers
const showNewFeature = useBooleanVariation('YOUR_FEATURE_FLAG_ID', false);
if (showNewFeature) {
    // The Application code to show the new feature
} else {
    // The code to run when the feature is off
}
```

:::note

The variation hooks will return the default value if the feature flag is missing in the SDK.

:::

### Variation types

The Bucketeer SDK supports the following variation types.

:::caution Deprecated

The `jsonVariation` interface is deprecated. Please use the `objectVariation` instead.

:::

```typescript showLineNumbers
useBooleanVariation(featureId: string, defaultValue: boolean): boolean;

useStringVariation(featureId: string, defaultValue: string): string;

useNumberVariation(featureId: string, defaultValue: number): number;

// The returned value will be either a BKTJsonObject or a BKTJsonArray. If no result is found, it will return the provided `defaultValue`, which can be of any type within `BKTValue`.
useObjectVariation(featureId: string, defaultValue: BKTValue): BKTValue;
```

#### Polling

The initialization process starts polling the latest evaluations from the Bucketeer server in the background using the interval `pollingInterval` configuration. React SDK **does not support** Background fetch.

#### Polling retry behavior

The Bucketeer SDK regularly polls the latest evaluations from the server based on the pollingInterval parameter. By default, the `pollingInterval` is set to 10 minutes, but you can adjust it to suit your needs.

If a polling request fails, the SDK initiates a retry procedure. The SDK attempts a new polling request every minute up to 5 times. If all five retry attempts fail, the SDK sends a new polling request once the `pollingInterval` time elapses. The table below shows this scenario:

<div className="center-table">
<table>
<thead>
  <tr>
    <th>Polling Time</th>
    <th>Retry Time</th>
    <th>Request Status</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>10:00</td>
    <td>-</td>
    <td>Fail</td>
  </tr>
  <tr>
    <td>- </td>
    <td>10:01</td>
    <td>Fail</td>
  </tr>
  <tr>
    <td>- </td>
    <td>10:02</td>
    <td>Fail</td>
  </tr>
  <tr>
    <td>- </td>
    <td>10:03</td>
    <td>Fail</td>
  </tr>
  <tr>
    <td>-</td>
    <td>10:04</td>
    <td>Fail</td>
  </tr>
  <tr>
    <td>-</td>
    <td>10:05</td>
    <td>Fail</td>
  </tr>
  <tr>
    <td>10:10</td>
    <td>-</td>
    <td>Successful</td>
  </tr>
</tbody>
</table>
</div>

The polling counter, which uses the `pollingInterval` information, resets in case of a successful retry. The table below shows the described scenario.

<div className="center-table">
<table>
<thead>
  <tr>
    <th>Polling Time</th>
    <th>Retry Time</th>
    <th>Request status</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>10:00</td>
    <td>-</td>
    <td>Fail</td>
  </tr>
  <tr>
    <td>- </td>
    <td>10:01</td>
    <td>Successful</td>
  </tr>
  <tr>
    <td>10:11</td>
    <td>-</td>
    <td>Successful</td>
  </tr>
</tbody>
</table>
</div>

### Handling exceptions

While most of the time error is handled internally, some methods throw `BKTException` when something goes wrong.  
Those methods are:

- **`initializeBKTClient()`**
- **`BKTClient#fetchEvaluations()`**
- **`BKTClient#flush()`**

These methods return `Promise` and might reject with `BKTException`, so you should make sure to catch the error.

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

The `BKTEvaluationDetails<T>` object contains detailed information about feature flag evaluation:

<Tabs>
<TabItem value="jsx" label="Interface">

```typescript showLineNumbers
export interface BKTEvaluationDetails<T extends BKTValue> {
  readonly featureId: string;       // The ID of the feature flag.
  readonly featureVersion: number;  // The version of the feature flag.
  readonly userId: string;          // The ID of the user being evaluated.
  readonly variationId: string;     // The ID of the assigned variation.
  readonly variationName: string;   // The name of the assigned variation.
  readonly variationValue: T;       // The value of the assigned variation.
  readonly reason:
    | 'TARGET'        // Evaluated using individual targeting.
    | 'RULE'          // Evaluated using a custom rule.
    | 'DEFAULT'       // Evaluated using the default strategy.
    | 'CLIENT'        // The flag is missing in the cache; the default value was returned.
    | 'OFF_VARIATION' // Evaluated using the off variation.
    | 'PREREQUISITE'; // Evaluated using a prerequisite.
}
```

</TabItem>
</Tabs>

:::note Hook Behavior

- All hooks automatically **re-render** your component when flag values change
- Hooks **must** be used within components wrapped by `BucketeerProvider`
- If a flag is not found, hooks return the provided `defaultValue`

:::

## Supported features

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

### Reporting custom events

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

### Updating user evaluations

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
      console.error('Failed to update with BKTException:', error);
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

### Flushing events

This method will send all pending analytics events to the Bucketeer server as soon as possible. This process is asynchronous, so it returns before it is complete.

<Tabs>
<TabItem value="jsx" label="FlushEvents.jsx">

```jsx showLineNumbers
import { useContext } from 'react';
import { BucketeerContext } from '@bucketeer/react-client-sdk';

function FlushButton() {
  const { client } = useContext(BucketeerContext);

  const handleFlush = async () => {
    try {
      await client?.flush();
      console.log('Events flushed successfully');
    } catch (error) {
      console.error('Failed to flush events with BKTException:', error);
    }
  };

  return (
    <button onClick={handleFlush}>
      Flush Events
    </button>
  );
}
```

</TabItem>
</Tabs>

:::note

In regular use, you don't need to call the flush method because the events are sent every **30 seconds** in the background.

:::

### User attributes configuration

This feature will give you robust and granular control over what users can see on your application. You can add rules using these attributes on the console UI's feature flag's targeting tab.

<Tabs>
<TabItem value="jsx" label="UserConfiguration.jsx">

```jsx showLineNumbers
import { defineBKTUser, defineBKTConfigForReact, initializeBKTClient } from '@bucketeer/react-client-sdk';

const attributes = {
  app_version: '1.0.0',
  os_version: '11.0.0',
  device_model: 'pixel-5',
  language: 'english',
  genre: 'female'
};

const user = defineBKTUser({
  id: 'USER_ID',
  customAttributes: attributes
});

const config = defineBKTConfigForReact({
  apiKey: 'YOUR_API_KEY',
  apiEndpoint: 'YOUR_API_ENDPOINT',
  appVersion: '1.0.0',
  featureTag: 'web',
});

await initializeBKTClient(config, user);
```

</TabItem>
</Tabs>

### Getting user information

This method will return the current user configured in the SDK. This is useful when you want to check the current user id and attributes before updating them through [updateUserAttributes](#updating-user-attributes).

<Tabs>
<TabItem value="jsx" label="GetUserInfo.jsx">

```jsx showLineNumbers
import { useContext } from 'react';
import { BucketeerContext } from '@bucketeer/react-client-sdk';

function UserInfo() {
  const { client } = useContext(BucketeerContext);

  const handleGetUser = () => {
    const user = client?.currentUser();
    console.log('Current user:', user);
  };

  return (
    <button onClick={handleGetUser}>
      Get User Info
    </button>
  );
}
```

</TabItem>
</Tabs>

### Listening to evaluation updates

The SDK can notify when the evaluation is updated.
The listener can detect both automatic polling and manual fetching.

<Tabs>
<TabItem value="jsx" label="EvaluationListener.jsx">

```jsx showLineNumbers
import { useContext, useEffect } from 'react';
import { BucketeerContext } from '@bucketeer/react-client-sdk';

function EvaluationListener() {
  const { client } = useContext(BucketeerContext);

  useEffect(() => {
    if (!client) return;

    // Add listener - returned value is used when you want to remove listener
    const key = client.addEvaluationUpdateListener(() => {
      const showNewFeature = client.booleanVariation("YOUR_FEATURE_FLAG_ID", false);
      if (showNewFeature) {
        // The Application code to show the new feature
      } else {
        // The code to run when the feature is off
      }
    });

    // Cleanup: Remove listener on component unmount
    return () => {
      client.removeEvaluationUpdateListener(key);
      // Or remove all listeners: client.clearEvaluationUpdateListeners();
    };
  }, [client]);

  return <div>Listening for evaluation updates...</div>;
}
```

</TabItem>
</Tabs>

### Destroying client

There are cases you might want to switch the user ID or reduce resources when the application is in the background.<br />
For those cases, you can call the destroy function, which will clear the client instance.

<Tabs>
<TabItem value="jsx" label="DestroyClient.jsx">

```jsx showLineNumbers
import { useEffect } from 'react';
import { destroyBKTClient } from '@bucketeer/react-client-sdk';

function App() {
  useEffect(() => {
    // Cleanup on component unmount
    return () => destroyBKTClient();
  }, []);

  // Your app content...
}
```

</TabItem>
</Tabs>

:::tip

If you want to switch the user ID, please call the [flush](#flushing-events) interface before calling the destroy, so that all the pending events can be sent before clearing the client instance, then call the [initialize](#initializing-client) interface with the new user information.

:::

## Best Practices

1. **Use feature tags** to improve performance by filtering relevant flags
2. **Handle initialization timeouts** gracefully - the client may still work
3. **Update user attributes** when user context changes (login, plan changes, etc.)
4. **Use TypeScript** for better type safety with the provided types (see [Re-exported from JavaScript SDK](#re-exported-from-javascript-sdk) section)
5. **Avoid frequent manual fetches** - let automatic polling handle updates

## Re-exported from JavaScript SDK

The React SDK re-exports all functionality from the JavaScript SDK, allowing you to use any JavaScript SDK features directly:

This means you can access all JavaScript SDK functionality without needing to install the JavaScript SDK separately. You can use features like:
- Direct access to the Bucketeer client instance (by `getBKTClient()`) without `useContext`
- Direct client methods (`client.booleanVariation()`, `client.track()`, etc.)
- Advanced configuration options
- Error handling utilities
- All TypeScript types and interfaces

:::tip

When using React components, prefer the React hooks (ie: `useBooleanVariation`) over direct client methods for better integration with React's rendering cycle. Use direct client methods only when necessary (e.g., in event handlers or outside React components).

:::

For complete API reference and advanced features, see the [JavaScript SDK documentation](/sdk/client-side/javascript).
