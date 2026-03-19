---
title: React Native
slug: /sdk/client-side/react-native
toc_max_heading_level: 4
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

This category contains topics explaining how to configure Bucketeer's React Native SDK.

:::warning React Native SDK Version (Beta)

Bucketeer React Native SDK is a beta version. Breaking changes may be introduced before general release.

:::

## Key Points

- Most APIs and usage are **identical** to the [React SDK](/sdk/client-side/react)
- The main difference: use `defineBKTConfigForReactNative` instead of `defineBKTConfigForReact`
- Requires AsyncStorage as a native dependency

## Requirements

Before starting, ensure that you follow the [Getting Started](/getting-started) guide.

**React Native Version Support:**
- ✅ Supported: React 18.2.0 - 18.3.x and React Native 0.76.0 - 0.78.x
- ⚠️ May work: React 18.0.0 - 18.1.x (not officially supported)
- ❌ Not supported: React 19.0.0 and above, React Native 0.79.0 and above

## Getting started

### Installing dependency

Install the dependency in your application.

<Tabs>
<TabItem value="npm" label="npm">

```sh showLineNumbers
npm install @bucketeer/react-native-client-sdk
```

</TabItem>
<TabItem value="yarn" label="Yarn">

```sh showLineNumbers
yarn add @bucketeer/react-native-client-sdk
```

</TabItem>
</Tabs>

#### AsyncStorage Dependency

This SDK uses `@react-native-async-storage/async-storage` for bootstrapping, which is a **native dependency**.

<Tabs>
<TabItem value="expo" label="Expo Projects">

Adding the Bucketeer React Native SDK from npm and re-running should suffice. If it doesn't work, install AsyncStorage explicitly:

```sh showLineNumbers
npx expo install @react-native-async-storage/async-storage
```

</TabItem>
<TabItem value="bare" label="Bare React Native">

You need to explicitly add AsyncStorage as a dependency and re-run pod install:

```sh showLineNumbers
npm install @react-native-async-storage/async-storage
cd ios && pod install  # For iOS
```

</TabItem>
</Tabs>

:::info

Auto-linking does not work with transitive dependencies, so you must install AsyncStorage explicitly in bare React Native projects. For more details, see the [AsyncStorage documentation](https://react-native-async-storage.github.io/async-storage/docs/install/).

:::

### Importing client

Import the necessary components and configure the SDK:

<Tabs>
<TabItem value="tsx" label="Importing">

```tsx showLineNumbers
import React, { useEffect, useState } from 'react';
import {
  BucketeerProvider,
  defineBKTConfigForReactNative, // Note: React Native specific
  defineBKTUser,
  initializeBKTClient,
  getBKTClient,
  destroyBKTClient,
  type BKTClient,
} from '@bucketeer/react-native-client-sdk';
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
<TabItem value="tsx" label="Configuration">

```tsx showLineNumbers
const config = defineBKTConfigForReactNative({
  apiKey: 'YOUR_API_KEY',
  apiEndpoint: 'YOUR_API_ENDPOINT',
  appVersion: '1.0.0',
  featureTag: 'mobile', // Optional but recommended
});

const user = defineBKTUser({
  id: 'USER_ID',
  customAttributes: {},
});
```

</TabItem>
</Tabs>

:::warning Important

Use `defineBKTConfigForReactNative` instead of `defineBKTConfigForReact` when working with React Native. This ensures proper React Native-specific configuration including AsyncStorage integration.

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

### Initializing client

Initialize the client in your root component:

<Tabs>
<TabItem value="tsx" label="Initialization">

```tsx showLineNumbers
export default function App() {
  const [client, setClient] = useState<BKTClient | null>(null);
  const [isInitialized, setIsInitialized] = useState(false);

  useEffect(() => {
    const init = async () => {
      try {
        await initializeBKTClient(config, user);
        const bktClient = getBKTClient();
        setClient(bktClient);
      } catch (error) {
        if (error instanceof Error && error.name === 'TimeoutException') {
          // Client is still initialized despite timeout
          console.warn('Initialization timed out, but client is ready');
          const bktClient = getBKTClient();
          setClient(bktClient);
        } else {
          console.error('Failed to initialize:', error);
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
    return <LoadingComponent />; // Your loading component
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
<TabItem value="tsx" label="Custom Timeout">

```tsx showLineNumbers
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

### Use Feature Flag Hooks

All [React SDK hooks](/sdk/client-side/react#hook-reference) work identically in React Native:

<Tabs>
<TabItem value="tsx" label="MyScreen.tsx">

```tsx showLineNumbers
import React from 'react';
import { View, Text } from 'react-native';
import {
  useBooleanVariation,
  useStringVariation,
  useNumberVariation,
  useObjectVariation,
} from '@bucketeer/react-native-client-sdk';

function MyScreen() {
  // Boolean feature flag
  const showNewFeature = useBooleanVariation('new-feature-enabled', false);
  
  // String feature flag
  const theme = useStringVariation('app-theme', 'light');
  
  // Number feature flag
  const maxItems = useNumberVariation('max-items', 10);
  
  // Object feature flag
  const config = useObjectVariation('app-config', { timeout: 5000 });

  return (
    <View>
      {showNewFeature && <NewFeatureComponent />}
      <Text>Theme: {theme}</Text>
      <Text>Max items: {maxItems}</Text>
      <Text>Timeout: {config.timeout}ms</Text>
    </View>
  );
}
```

</TabItem>
</Tabs>

### Evaluating user

The variation hooks determine whether or not a feature flag is enabled for a specific user.<br />
To check which variation a specific user will receive, you can use the hooks like below.

```tsx showLineNumbers
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

The initialization process starts polling the latest evaluations from the Bucketeer server in the background using the interval `pollingInterval` configuration. React Native SDK **does not support** Background fetch.

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

## Supported features

### Getting evaluation details

All React Native hooks also have details versions that return evaluation information:

```typescript showLineNumbers
useBooleanVariationDetails(flagId: string, defaultValue: boolean): BKTEvaluationDetails<boolean>;
useStringVariationDetails(flagId: string, defaultValue: string): BKTEvaluationDetails<string>;
useNumberVariationDetails(flagId: string, defaultValue: number): BKTEvaluationDetails<number>;
useObjectVariationDetails<T>(flagId: string, defaultValue: T): BKTEvaluationDetails<T>;
```

#### Evaluation Details Object

The `BKTEvaluationDetails<T>` object contains detailed information about feature flag evaluation:

<Tabs>
<TabItem value="tsx" label="Interface">

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

### Updating user evaluations

Depending on the use case, you may need to ensure the evaluations in the SDK are up to date before requesting the variation.

<Tabs>
<TabItem value="tsx" label="RefreshButton.tsx">

```tsx showLineNumbers
import { useContext } from 'react';
import { BucketeerContext } from '@bucketeer/react-native-client-sdk';
import { Button } from 'react-native';

function RefreshButton() {
  const { client } = useContext(BucketeerContext);

  const handleRefresh = async () => {
    try {
      await client?.fetchEvaluations(5000); // 5 second timeout
      console.log('Evaluations updated');
    } catch (error) {
      console.error('Failed to update with BKTException:', error);
    }
  };

  return (
    <Button title="Refresh Feature Flags" onPress={handleRefresh} />
  );
}
```

</TabItem>
</Tabs>

### Reporting custom events

This method lets you save user actions in your application as events. You can connect these events to metrics in the experiments console UI.

<Tabs>
<TabItem value="tsx" label="EventTracking.tsx">

```tsx showLineNumbers
import { useContext } from 'react';
import { BucketeerContext } from '@bucketeer/react-native-client-sdk';
import { Button } from 'react-native';

function PurchaseButton() {
  const { client } = useContext(BucketeerContext);

  const handlePurchase = (amount: number) => {
    // Track purchase event with value
    client?.track('purchase_completed', amount);
  };

  return (
    <Button title="Buy Now" onPress={() => handlePurchase(99.99)} />
  );
}
```

</TabItem>
</Tabs>

### Flushing events

This method will send all pending analytics events to the Bucketeer server as soon as possible. This process is asynchronous, so it returns before it is complete.

<Tabs>
<TabItem value="tsx" label="FlushEvents.tsx">

```tsx showLineNumbers
import { useContext } from 'react';
import { BucketeerContext } from '@bucketeer/react-native-client-sdk';
import { Button } from 'react-native';

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
    <Button title="Flush Events" onPress={handleFlush} />
  );
}
```

</TabItem>
</Tabs>

:::note

In regular use, you don't need to call the flush method because the events are sent every **30 seconds** in the background.

:::

### Updating user attributes

This method will update all the current user attributes. This is useful in case the user attributes update dynamically on the application after initializing the SDK.

<Tabs>
<TabItem value="tsx" label="UserProfile.tsx">

```tsx showLineNumbers
import { useContext } from 'react';
import { BucketeerContext } from '@bucketeer/react-native-client-sdk';
import { Platform } from 'react-native';

function UserProfile() {
  const { client } = useContext(BucketeerContext);

  const handleUpgrade = () => {
    client?.updateUserAttributes({
      plan: 'premium',
      tier: 'gold',
      platform: Platform.OS, // 'ios' or 'android'
      version: Platform.Version,
    });
  };

  return (
    <Button title="Upgrade to Premium" onPress={handleUpgrade} />
  );
}
```

</TabItem>
</Tabs>

:::caution

This updating method will override the current data.

:::

### User attributes configuration

This feature will give you robust and granular control over what users can see on your application. You can add rules using these attributes on the console UI's feature flag's targeting tab.

<Tabs>
<TabItem value="tsx" label="UserConfiguration.tsx">

```tsx showLineNumbers
import { defineBKTUser, defineBKTConfigForReactNative, initializeBKTClient } from '@bucketeer/react-native-client-sdk';
import { Platform } from 'react-native';

const attributes = {
  app_version: '1.0.0',
  os_version: Platform.Version.toString(),
  device_model: Platform.OS,
  language: 'english',
  platform: Platform.OS, // 'ios' or 'android'
};

const user = defineBKTUser({
  id: 'USER_ID',
  customAttributes: attributes
});

const config = defineBKTConfigForReactNative({
  apiKey: 'YOUR_API_KEY',
  apiEndpoint: 'YOUR_API_ENDPOINT',
  appVersion: '1.0.0',
  featureTag: 'mobile',
});

await initializeBKTClient(config, user);
```

</TabItem>
</Tabs>

### Getting user information

This method will return the current user configured in the SDK. This is useful when you want to check the current user id and attributes before updating them through [updateUserAttributes](#updating-user-attributes).

<Tabs>
<TabItem value="tsx" label="GetUserInfo.tsx">

```tsx showLineNumbers
import { useContext } from 'react';
import { BucketeerContext } from '@bucketeer/react-native-client-sdk';
import { Button } from 'react-native';

function UserInfo() {
  const { client } = useContext(BucketeerContext);

  const handleGetUser = () => {
    const user = client?.currentUser();
    console.log('Current user:', user);
  };

  return (
    <Button title="Get User Info" onPress={handleGetUser} />
  );
}
```

</TabItem>
</Tabs>

### Listening to evaluation updates

The SDK can notify when the evaluation is updated.
The listener can detect both automatic polling and manual fetching.

<Tabs>
<TabItem value="tsx" label="EvaluationListener.tsx">

```tsx showLineNumbers
import { useContext, useEffect } from 'react';
import { BucketeerContext } from '@bucketeer/react-native-client-sdk';

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

  return null; // This is a listener component
}
```

</TabItem>
</Tabs>

### Destroying client

There are cases you might want to switch the user ID or reduce resources when the application is in the background.<br />
For those cases, you can call the destroy function, which will clear the client instance.

<Tabs>
<TabItem value="tsx" label="DestroyClient.tsx">

```tsx showLineNumbers
import { useEffect } from 'react';
import { destroyBKTClient } from '@bucketeer/react-native-client-sdk';

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

## Re-exported from React SDK

The React Native SDK re-exports all functionality from the [React SDK](/sdk/client-side/react) and [JavaScript SDK](/sdk/client-side/javascript), allowing you to use any React SDK features directly:

This means you can access all React SDK and JavaScript SDK functionality without needing to install the React SDK and JavaScript SDK separately. You can use features like:

- All React hooks (`useBooleanVariation`, `useStringVariation`, etc.)
- Direct access to the Bucketeer client instance (by `getBKTClient()`) without `useContext`
- Direct client methods (`client.booleanVariation()`, `client.track()`, etc.)
- Advanced configuration options
- Error handling utilities
- All TypeScript types and interfaces

:::tip

When using React Native components, prefer the React hooks (ie: `useBooleanVariation`) over direct client methods for better integration with React's rendering cycle. Use direct client methods only when necessary (e.g., in event handlers or outside React components).

:::

## Best Practices

1. **Always use `defineBKTConfigForReactNative`** - Essential for proper AsyncStorage integration
2. **Use feature tags** to improve performance by filtering relevant flags  
3. **Handle initialization timeouts** gracefully - the client may still work
4. **Handle app state changes** - Refresh evaluations on foreground, flush on background
5. **Include platform info** - Use Platform API for better targeting (iOS/Android)
6. **Handle network connectivity** - Refresh when network is restored
7. **Update user attributes** when user context changes (login, plan changes, etc.)
8. **Use TypeScript** for better type safety with the provided types (see [Re-exported from React SDK](#re-exported-from-react-sdk) section)
9. **Optimize for mobile** - Use appropriate polling intervals for battery life
10. **Avoid frequent manual fetches** - let automatic polling handle updates

## Troubleshooting

### AsyncStorage Issues

If you encounter AsyncStorage related errors:

1. Ensure AsyncStorage is properly installed
2. For bare React Native, run `pod install` after installation
3. For Expo, try `npx expo install @react-native-async-storage/async-storage`

### Metro Bundler Issues

If you see bundling errors, try:

```sh
npx react-native start --reset-cache
```

For complete API reference and advanced features, see the [React SDK documentation](/sdk/client-side/react) and [JavaScript SDK documentation](/sdk/client-side/javascript).
