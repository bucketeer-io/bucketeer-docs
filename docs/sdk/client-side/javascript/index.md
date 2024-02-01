---
title: JavaScript
slug: /sdk/client-side/javascript
toc_max_heading_level: 4
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

This category contains topics explaining how to configure Bucketeer's JavaScript SDK.

:::tip JavaScript SDK Version (Stable)

Bucketeer JavaScript SDK has reached the production stage, offering you a stable and reliable experience.

:::

## Getting started

Before starting, ensure that you follow the [Getting Started](/getting-started) guide.

### Installing dependency

Install the dependency in your application.

<Tabs>
<TabItem value="npm" label="npm">

```sh showLineNumbers
npm install @bucketeer/js-client-sdk
```

</TabItem>
<TabItem value="yarn" label="Yarn">

```sh showLineNumbers
yarn add @bucketeer/js-client-sdk
```

</TabItem>
</Tabs>

### Importing client

Import the Bucketeer client into your application code.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
import { BKTClient, getBKTClient, defineBKTConfig, defineBKTUser, initializeBKTClient } from '@bucketeer/js-client-sdk';
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
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const config = defineBKTConfig({
  apiKey: 'YOUR_API_KEY',
  apiEndpoint: 'YOUR_API_URL',
  featureTag: 'YOUR_FEATURE_TAG', // Optional
  appVersion: 'YOUR_APP_VERSION',
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
- **fetch** (Default is `window.fetch`)

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
const client = getBKTClient();
```

</TabItem>
</Tabs>

:::info Default timeout

The initialize process default timeout is 5 seconds.

Once initialization is finished, all the requests in the SDK use a timeout of 30 seconds.

:::

If you want to use the feature flag on Splash or Main views, the SDK cache may be old or not exist and may not have enough time to fetch the variations from the Bucketeer server. For this case, we recommend using the `Promise` returned from the initialize method. The Promise rejects with `BKTException` when something goes wrong.

:::caution Initialization Timeout error

During the initialization process, errors **are not** related to the initialization itself. Instead, they arise from a timeout request, indicating the variations data from the server weren't received. Therefore, the SDK will work as usual and update the variations in the next [polling](javascript#polling) request.

:::

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const timeout = 1000 // Default is 5 seconds
const initialFetchPromise = initializeBKTClient(config, user, timeout);
initialFetchPromise
  .then(() => {
    const client = getBKTClient();
    const showNewFeature = client?.booleanVariation('YOUR_FEATURE_FLAG_ID', false);
    if (showNewFeature) {
        // The Application code to show the new feature
    } else {
        // The code to run when the feature is off
    }
  })
  .catch((error) => {
    // Handle the error when there is no cache or the cache is not updated
  });
```

</TabItem>
</Tabs>

#### Polling

The initialize process starts polling right away the latest evaluations from the Bucketeer server in the background using the interval `pollingInterval` configuration. JavaScript SDK **does not support** Background fetch.

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

These methods return `Promise` and might reject with `BKTException``, so you should make sure to catch the error.

## Supported features

### Evaluating user

The variation method determines whether or not a feature flag is enabled for a specific user.<br />
To check which variation a specific user will receive, you can use the client like below.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const showNewFeature = client?.booleanVariation('YOUR_FEATURE_FLAG_ID', false);
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
booleanVariation(featureId: string, defaultValue: boolean): Promise<boolean>;

stringVariation(featureId: string, defaultValue: string): Promise<string>;

numberVariation(featureId: string, defaultValue: number): Promise<number>;

jsonVariation(featureId: string, defaultValue: object): Promise<object>;
```

</TabItem>
</Tabs>

### Updating user evaluations

Depending on the use case, you may need to ensure the evaluations in the SDK are up to date before requesting the variation.

The fetch method uses the following parameter. Make sure to wait for its completion.

- **Timeout** (Default is 30 seconds)

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const timeout = 5000;

await client?.fetchEvaluations(timeout);
```

</TabItem>
</Tabs>

### Reporting custom events

This method lets you save user actions in your application as events. You can connect these events to metrics in the experiments console UI.

In addition, you can pass a double value to the goal event. These values will sum and show as <br />`Value total` on the experiments console UI. This is useful if you have a goal event for tracking how much a user spent on your application buying items, etc.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
client?.track("YOUR_GOAL_ID", 10.50);
```

</TabItem>
</Tabs>

### Flushing events

This method will send all pending analytics events to the Bucketeer server as soon as possible. This process is asynchronous, so it returns before it is complete.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
await client?.flush();
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
};

client?.updateUserAttributes(attributes);
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
const user = client?.currentUser();
```

</TabItem>
</Tabs>

### Getting evaluation details

This method will return the evaluation details for a specific feature flag. This is useful if you need to know the variation reason or send this data elsewhere.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const evaluationDetails = client?.evaluationDetails("YOUR_FEATURE_FLAG_ID");
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
const key = client?.addEvaluationUpdateListener(() => {
  const showNewFeature = client?.booleanVariation("YOUR_FEATURE_FLAG_ID", false)
  if (showNewFeature) {
      // The Application code to show the new feature
  } else {
      // The code to run when the feature is off
  }
});

// Remove a listener associated with the key
client?.removeEvaluationUpdateListener(key);

// Remove all listeners
client?.clearEvaluationUpdateListeners();
```

</TabItem>
</Tabs>
