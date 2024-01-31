---
title: Node JS reference
sidebar_position: 1
slug: /sdk/server-side/node-js
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

This category contains topics explaining how to configure Bucketeer's Node.js SDK.

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

### Configuring client

Configure the SDK config and user configuration.

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

:::info Custom configuration

Depending on your use, you may want to change the optional configurations available.

- **pollingIntervalForRegisterEvents** (Default is 1 minutes)
- **logger** (Default is `logger.DefaultLogger`)

:::

### Initializing client

Initialize the client by passing the configurations in the previous step.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const client = initialize(config);
```

</TabItem>
</Tabs>

If you want to use the feature flag on Splash or Main views, and the user opens your application for the first time, it may not have enough time to fetch the variations from the Bucketeer server.

For this case, we recommend using the `Promise` returned from the initialize method. The Promise rejects with `BKTException` when something goes wrong.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const flag = await client.getBoolVariation(
	{ id: 'uid', data: {} }, // user info
	'YOUR_FEATURE_FLAG_ID', // feature id
	false, // default value
);
```

</TabItem>
</Tabs>

## Supported features

### Evaluating user

The variation method determines whether or not a feature flag is enabled for a specific user.<br />
To check which variation a specific user will receive, you can use the client like below.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const showNewFeature = client.getBoolVariation(
	{ id: 'uid', data: {} },
	'YOUR_FEATURE_FLAG_ID',
	false,
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

The variation method will return the default value if the feature flag is missing in the SDK.

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

This method lets you save user actions in your application as events. You can connect these events to metrics in the experiments console UI.

In addition, you can pass a number value to the goal event. These values will sum and show as <br />`Value total` on the experiments console UI. This is useful if you have a goal event for tracking how much a user spent on your application buying items, etc.

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
client.track({ id: 'uid', data: {} }, "YOUR_GOAL_ID", 10.50);
```

</TabItem>
</Tabs>

### Finalize client

This method will send all pending analytics events to the Bucketeer server as soon as possible and stop workers. The application should call destroy before the application stops otherwise, the remaining events can be lost.

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
}
});
const flag = await client.getBoolVariation(user, 'YOUR_FEATURE_FLAG_ID', false);
```

</TabItem>
</Tabs>
