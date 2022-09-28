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
<TabItem value="js" label="Javascript">

```js showLineNumbers
import * as BKTClient from '@bucketeer/sdk';
```

</TabItem>
</Tabs>

### Configuring client

Configure the SDK config and user configuration.

<Tabs>
<TabItem value="js" label="Javascript">

```js showLineNumbers
const config = {
  apiKey: 'YOUR_API_KEY',
  apiURL: 'YOUR_API_URL',
  featureTag: 'YOUR_FEATURE_TAG'
};

const user = {id: 'USER_ID'};
```

</TabItem>
</Tabs>

:::info Custom configuration

Depending on your use, you may want to change the optional configurations available.

- **pollingInterval** (Minimum 5 minutes. Default is 10 minutes)
- **backgroundPollingInterval** (Minimum 20 minutes. Default is 1 hour)
- **eventsFlushInterval** (Default is 30 seconds)
- **eventsMaxQueueSize** (Default is 50 events)

:::

:::note

The Bucketeer SDK doesn't save the user data. The Application must save and set it when initializing the client SDK.

:::

### Initializing client

Initialize the client by passing the configurations in the previous step.

<Tabs>
<TabItem value="js" label="Javascript">

```js showLineNumbers
const client = BKTClient.initialize(config, user);
```

</TabItem>
</Tabs>

:::note

The initialize process starts polling the latest variations from Bucketeer in the background using the interval `pollingInterval` configuration. When your application moves to the background state, it will use the `backgroundPollingInterval` configuration.

:::

If you want to use the feature flag on Splash or Main views, and the user opens your application for the first time, it may not have enough time to fetch the variations from the Bucketeer server.

For this case, we recommend using the callback in the initialize method.

<Tabs>
<TabItem value="js" label="Javascript">

```js showLineNumbers
// The callback will return without waiting until the fetching variation process finishes
const timeout = 1000 // Default is 5 seconds

const client = await BKTClient.initialize(config, user, timeout);
```

</TabItem>
</Tabs>

## Supported features

### Evaluating user

The variation method determines whether or not a feature flag is enabled for a specific user.<br />
To check which variation a specific user will receive, you can use the client like below.

<Tabs>
<TabItem value="js" label="Javascript">

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
<TabItem value="js" label="Javascript">

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
<TabItem value="js" label="Javascript">

```js showLineNumbers
// It will unlock without waiting until the fetching variation process finishes
val timeout = 1000 // Default is 5 seconds

await client.fetchEvaluations(timeout);
```

</TabItem>
</Tabs>

### Updating user variations in real-time

The Bucketeer SDK supports FCM ([Firebase Cloud Messaging](https://firebase.google.com/docs/cloud-messaging)).
Every time you change some feature flag, Bucketeer will send notifications using the FCM API to notify the client so that you can update the variations in real-time.

Assuming you already have the FCM implementation in your application.

<Tabs>
<TabItem value="js" label="Javascript">

```js showLineNumbers
// TODO
```

</TabItem>
</Tabs>

:::note

You need to register your FCM API Key on the console UI. [See more](#).

:::

### Reporting custom events

This method lets you save user actions in your application as events. You can connect these events to metrics in the experiments console UI.

In addition, you can pass a double value to the goal event. These values will sum and show as <br />`Value total` on the experiments console UI. This is useful if you have a goal event for tracking how much a user spent on your application buying items, etc.

<Tabs>
<TabItem value="js" label="Javascript">

```js showLineNumbers
client.track("YOUR_GOAL_ID", 10.50);
```

</TabItem>
</Tabs>

### Flushing events

This method will send all pending analytics events to the Bucketeer server as soon as possible. This process is asynchronous, so it returns before it is complete.

<Tabs>
<TabItem value="js" label="Javascript">

```js showLineNumbers
client.flush();
```

</TabItem>
</Tabs>

:::note

In regular use, you don't need to call the flush method because the events are sent every 30 seconds in the background.

:::

### User attributes configuration

This feature will give you robust and granular control over what users can see on your application. You can add rules using these attributes on the console UI's feature flag's targeting tab. [See more](#).

<Tabs>
<TabItem value="js" label="Javascript">

```js showLineNumbers
const attributes = new Map([
  ['app_version', '1.0.0'],
  ['os_version', '11.0.0'],
  ['device_model', 'pixel-5'],
  ['language', 'english'],
  ['genre', 'female']
]);

const user = {
  id: 'USER_ID',
  attributes: attributes
};

const client = BKTClient.initialize(config, user);
```

</TabItem>
</Tabs>

### Updating user attributes

This method will update all the current user attributes. This is useful in case the user attributes update dynamically on the application after initializing the SDK.

<Tabs>
<TabItem value="js" label="Javascript">

```js showLineNumbers
const attributes = new Map([
  ['app_version', '1.0.0'],
  ['os_version', '11.0.0'],
  ['device_model', 'pixel-5'],
  ['language', 'english'],
  ['genre', 'female']
]);

await client.updateUserAttributes(attributes);
```

</TabItem>
</Tabs>

:::caution

This updating method will override the current data.

:::

### Getting user information

This method will return the current user configured in the SDK. This is useful when you want to check the current user id and attributes before updating them through [updateUserAttributes](#getting-user-information).

<Tabs>
<TabItem value="js" label="Javascript">

```js showLineNumbers
const user = await client.currentUser();
```

</TabItem>
</Tabs>

### Getting evaluation details

This method will return the evaluation details for a specific feature flag. This is useful if you need to know the variation reason or send this data elsewhere.

<Tabs>
<TabItem value="js" label="Javascript">

```js showLineNumbers
const evaluationDetails = await client.evaluationDetails("YOUR_FEATURE_FLAG_ID");
```

:::note

This method will return null if the feature flag is missing in the SDK.

:::

</TabItem>
</Tabs>
