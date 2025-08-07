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

## Installation

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

### AsyncStorage Dependency

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

## Setup

### 1. Configure and Initialize Client

First, import the necessary components and configure the SDK:

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

Configure the SDK and user settings:

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

Initialize the client in your root component:

<Tabs>
<TabItem value="tsx" label="Initialization">

```tsx showLineNumbers
export default function App() {
  const [client, setClient] = useState<BKTClient | null>(null);

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
          return;
        }
      }
    };

    init();
    return () => destroyBKTClient();
  }, []);

  if (!client) {
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

### 2. Use Feature Flag Hooks

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

## API Reference

This SDK re-exports all APIs from the [React SDK](/sdk/client-side/react) and [JavaScript SDK](/sdk/client-side/javascript). All hooks, components, and methods work identically.

## Best Practices

1. **Use `defineBKTConfigForReactNative`** - Essential for proper AsyncStorage integration
2. **Handle app state changes** - Refresh evaluations on foreground, flush on background
3. **Include platform info** - Use Platform API for better targeting
4. **Handle network connectivity** - Refresh when network is restored
5. **Optimize for mobile** - Use appropriate polling intervals for battery life

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
