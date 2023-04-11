---
title: Javascript reference
slug: /sdk/client-side/javascript
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

This category contains topics explaining how to configure Bucketeer's Javascript SDK.

## Getting started

Before starting, ensure that you follow the [Getting started](/) guide.

### Installing dependency

Install the dependency in your application.

<Tabs>
<TabItem value="npm" label="npm">

```sh showLineNumbers
npm install @bucketeer/sdk
```

</TabItem>
<TabItem value="yarn" label="Yarn">

```sh showLineNumbers
yarn add @bucketeer/sdk
```

</TabItem>
</Tabs>

### Importing client

Import the Bucketeer client into your application code.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
import { BKTClient, getBKTClient, defineBKTConfig, defineBKTUser, initializeBKTClient } from '@bucketeer/sdk';
```

</TabItem>
</Tabs>

### Configuring client

Configure the SDK config and user configuration.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const config = defineBKTConfig({
  apiKey: 'YOUR_API_KEY',
  apiEndpoint: 'YOUR_API_URL',
  featureTag: 'YOUR_FEATURE_TAG',
  appVersion: 'YOUR_APP_VERSION',
  fetch: window.fetch // your fetch implementation
});

const user = defineBKTUser({
  id: 'USER_ID'
});
```

</TabItem>
</Tabs>

:::info Custom configuration

Depending on your use, you may want to change the optional configurations available.

- **pollingInterval** (Minimum 5 minutes. Default is 10 minutes)
- **eventsFlushInterval** (Default is 30 seconds)
- **eventsMaxQueueSize** (Default is 50 events)
- **storageKeyPrefix** (Default is empty) 
- **userAgent** (Default is `window.navigator.userAgent`)

:::

:::note

The Bucketeer SDK doesn't save the user data. The Application must save and set it when initializing the client SDK.

:::

### Initializing client

Initialize the client by passing the configurations in the previous step.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
await initializeBKTClient(config, user);
const client = getBKTClient()
```

</TabItem>
</Tabs>

:::note

The initialize process starts polling the latest variations from Bucketeer in the background using the interval `pollingInterval` configuration.

:::

If you want to use the feature flag on Splash or Main views, and the user opens your application for the first time, it may not have enough time to fetch the variations from the Bucketeer server.

For this case, we recommend using the `Promise` returned from the initialize method. The Promise rejects with `BKTException` when something goes wrong.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
// The callback will return without waiting until the fetching variation process finishes
const timeout = 1000 // Default is 5 seconds

await initializeBKTClient(config, user, timeout);
const client = getBKTClient()
```

</TabItem>
</Tabs>

### Handling exceptions

While most of the time error is handled internally, some methods throw `BKTException` when something goes wrong.  
Those methods are:

- **`initializeBKTClient()`**
- **`BKTClient#fetchEvaluations()`**
- **`BKTClient#flush()`**

These methods return `Promise` and might rejects with `BKTException`, so you should make sure to catch the error.

## Supported features

### Evaluating user

The variation method determines whether or not a feature flag is enabled for a specific user.<br />
To check which variation a specific user will receive, you can use the client like below.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const showNewFeature = client.boolVariation('YOUR_FEATURE_FLAG_ID', false);
if (showNewFeature) {
    // The Application code to show the new feature
} else {
    // The code to run when the feature is off
}
```

</TabItem>
</Tabs>

:::note

The variation method will return the default value if the feature flag is missing in the SDK.

:::

### Variation types

The Bucketeer SDK supports the following variation types.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
boolVariation(featureId: string, defaultValue: boolean): Promise<boolean>;

stringVariation(featureId: string, defaultValue: string): Promise<string>;

numberVariation(featureId: string, defaultValue: number): Promise<number>;

jsonVariation(featureId: string, defaultValue: object): Promise<object>;
```

</TabItem>
</Tabs>

### Updating user variations

Sometimes depending on your use, you may need to ensure the variations in the SDK are up to date before evaluating a user.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
// It will unlock without waiting until the fetching variation process finishes
val timeout = 1000 // Default is 5 seconds

await client.fetchEvaluations(timeout);
```

</TabItem>
</Tabs>

### Reporting custom events

This method lets you save user actions in your application as events. You can connect these events to metrics in the experiments console UI.

In addition, you can pass a double value to the goal event. These values will sum and show as <br />`Value total` on the experiments console UI. This is useful if you have a goal event for tracking how much a user spent on your application buying items, etc.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
client.track("YOUR_GOAL_ID", 10.50);
```

</TabItem>
</Tabs>

### Flushing events

This method will send all pending analytics events to the Bucketeer server as soon as possible. This process is asynchronous, so it returns before it is complete.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
await client.flush();
```

</TabItem>
</Tabs>

:::note

In regular use, you don't need to call the flush method because the events are sent every 30 seconds in the background.

:::

### User attributes configuration

This feature will give you robust and granular control over what users can see on your application. You can add rules using these attributes on the console UI's feature flag's targeting tab. [See more](#).

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const attributes = {
  app_version: '1.0.0',
  os_version: '11.0.0',
  device_model: 'pixel-5',
  language: 'english',
  genre: 'female'
};

const user = defineBKTUser({
  id: 'USER_ID',
  attributes: attributes
});

await initializeBKTClient(config, user);
```

</TabItem>
</Tabs>

### Updating user attributes

This method will update all the current user attributes. This is useful in case the user attributes update dynamically on the application after initializing the SDK.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const attributes = {
  app_version: '1.0.0',
  os_version: '11.0.0',
  device_model: 'pixel-5',
  language: 'english',
  genre: 'female'
}

client.updateUserAttributes(attributes);
```

</TabItem>
</Tabs>

:::caution

This updating method will override the current data.

:::

### Getting user information

This method will return the current user configured in the SDK. This is useful when you want to check the current user id and attributes before updating them through [updateUserAttributes](#getting-user-information).

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const user = client.currentUser();
```

</TabItem>
</Tabs>

### Getting evaluation details

This method will return the evaluation details for a specific feature flag. This is useful if you need to know the variation reason or send this data elsewhere.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const evaluationDetails = client.evaluationDetails("YOUR_FEATURE_FLAG_ID");
```

:::note

This method will return null if the feature flag is missing in the SDK.

:::

</TabItem>
</Tabs>

### Listening to evaluation updates

BKTClient can notify when the evaluation is updated.  
The listener can detect both automatic polling and manual fetching.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
// Returned value is used when you want to remove listener
val key = client.addEvaluationUpdateListener(() => {
  val showNewFeature = client.booleanVariation("YOUR_FEATURE_FLAG_ID", false)
  if (showNewFeature) {
      // The Application code to show the new feature
  } else {
      // The code to run when the feature is off
  }
})

// Remove a listener associated with the key
client.removeEvaluationUpdateListener(key)

// Remove all listeners
client.clearEvaluationUpdateListeners()
```

</TabItem>
</Tabs>
