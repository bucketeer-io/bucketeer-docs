---
title: Node.js
sidebar_position: 1
slug: /sdk/server-side/node-js
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

This category contains topics explaining how to configure Bucketeer's Node.js SDK.

:::caution

The Node.js SDK is currently in its Beta stage.
If you find any issues or have suggestions for improvement, feel free to open an [issue](https://github.com/bucketeer-io/node-server-sdk/issues).

:::

## Getting started

Before starting, ensure that you follow the [Getting Started](/getting-started) guide.

### Installing dependency

Install the dependency in your application.

<Tabs>
<TabItem value="npm" label="npm">

```sh showLineNumbers
npm install @bucketeer/node-server-sdk
```

</TabItem>
<TabItem value="yarn" label="Yarn">

```sh showLineNumbers
yarn add @bucketeer/node-server-sdk
```

</TabItem>
</Tabs>

### Importing SDK

Import the Bucketeer SDK into your application code.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
import { initialize } from '@bucketeer/node-server-sdk';
```

</TabItem>
</Tabs>

### Initializing client

The SDK supports local and remote evaluations.

- [Local evaluation](#evaluating-users-within-sdk-locally): Evaluates end users within SDK locally.
- [Remote evaluation](#remote-evaluation): Evaluate end users on the server.

Configure the SDK config and user configuration.
:::info

The **tag** setting is the tag you configure when creating a Feature Flag. It will evaluate all the Feature Flags in the environment when it is not configured.<br />
**We strongly recommend** using tags to speed up the evaluation process and reduce the response latency.

:::
<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const config = {
  host: 'YOUR_API_URL',
  token: 'YOUR_API_KEY',
  tag: 'YOUR_TAG',
};
```

</TabItem>
</Tabs>

#### Custom configuration

:::info Custom configuration

Depending on your use, you may want to change the optional configurations available.

- **pollingIntervalForRegisterEvents** (Default is 1 minute - specify in milliseconds)
- **logger** (Default is `logger.DefaultLogger`)
- **enableLocalEvaluation** (Default is false)
- **cachePollingInterval** (Default is 1 minute - specify in milliseconds)

For more information, please check the Option implementation [here](https://github.com/bucketeer-io/node-server-sdk/blob/master/src/config.ts).

:::

#### Remote evaluation

To evaluate users on the server side you must create an API Key using the `Client SDK` role.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
  const config = {
    host: HOST,
    token: TOKEN,
    tag: FEATURE_TAG,
  }
  const client = initialize(config);
```

</TabItem>
</Tabs>

Once the SDK is configured, please check this [section](#evaluating-user) to learn how to get the variation for a user.

#### Evaluating users within SDK locally

By evaluating users locally you can improve response time significantly.<br />
To evaluate them you must create an API Key using the `Server SDK` role.

The SDK will poll the Feature Flags and Segment Users from the server, and cache them in memory.

:::caution

The Server SDK API Key has access to all Feature Flags and Segment Users in the environment.<br />
Keep in mind that it might contain sensitive information, so be careful when sharing the key with others.

:::

When initializing the SDK you must enable the local evaluation setting.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
  const config = {
    host: HOST,
    token: TOKEN,
    tag: FEATURE_TAG,
    enableLocalEvaluation: true,
  }
  const client = initialize(config);
```

</TabItem>
</Tabs>

Once the SDK is configured, please check this [section](#evaluating-user) to learn how to get the variation for a user.


## Supported features

### Evaluating user

The variation method determines whether or not a feature flag is enabled for a specific user.<br />
To check which variation a specific user will receive, you can use the client like below.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const showNewFeature = await client.getBoolVariation(
  user: User({
    id: 'USER_ID',
    data: {}, // The user attributes are optional
  }),
  featureId: 'YOUR_FEATURE_FLAG_ID',
  defaultValue: false
);
if (showNewFeature) {
  // The Application code to show the new feature
} else {
  // The code to run when the feature is off
}
```

</TabItem>
</Tabs>

:::note

The variation method will return the default value if the feature flag doesn't exist, or is archived, or if the request fails.

:::

### Variation types

The Bucketeer SDK supports the following variation types.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
getStringVariation(user: User, featureId: string, defaultValue: string): Promise<string>;

getBoolVariation(user: User, featureId: string, defaultValue: boolean): Promise<boolean>;

getNumberVariation(user: User, featureId: string, defaultValue: number): Promise<number>;

getJsonVariation(user: User, featureId: string, defaultValue: object): Promise<object>;
```

</TabItem>
</Tabs>

### Reporting custom events

This method lets you save user actions in your application as events. You can connect these events to metrics in the experiments or in the kill switch (auto operation) on the console UI.

In addition, you can pass a number value to the goal event. These values will sum and show as <br />`Value total` on the experiments console UI. This is useful if you have a goal event for tracking how much a user spent on your application buying items, etc.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
client.track(
  user: {
    id: 'USER_ID', 
    data: {}, // The user attributes are optional
  },
  goalId: 'YOUR_GOAL_ID', 
  value: 10.50,
);
```

</TabItem>
</Tabs>

### Destroy client

This method will send all pending analytics events to the Bucketeer server as soon as possible and stop the workers. The application should call this before the application stops. Otherwise, the remaining events can be lost.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
await client.destroy();
```

</TabItem>
</Tabs>

### User attributes configuration

This feature will give you robust and granular control over what users can see on your application. You can add rules using these attributes on the console UI's feature flag's targeting tab. [See more](/feature-flags/creating-feature-flags/targeting#user-attributes).

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
import { User } from '@bucketeer/node-server-sdk';

const user = User({
  id: 'USER_ID',
  data: {
    app_version: '1.0.0',
    os_version: '11.0.0',
    device_model: 'pixel-5',
    language: 'english',
    genre: 'female'
  },
});
const showNewFeature = await client.getBoolVariation(user, 'YOUR_FEATURE_FLAG_ID', false);
```

</TabItem>
</Tabs>
