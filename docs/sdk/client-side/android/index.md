---
title: Android
slug: /sdk/client-side/android
toc_max_heading_level: 4
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

This category contains topics explaining how to configure Bucketeer's Android SDK.

:::tip Android SDK Version (Stable)

Bucketeer Android SDK has reached the production stage, offering you a stable and reliable experience.

:::

:::info Compatibility

The Bucketeer SDK is compatible with Android SDK versions 21 and higher (Android 5.0, Lollipop).

:::

## Getting started

Before starting, ensure that you follow the [Getting Started](/getting-started) guide.

:::info Using ProGuard or R8

The **aar artifact** will automatically include the configuration. If you don't use ProGuard or R8, you must include the configuration from [proguard-rules.pro](https://github.com/bucketeer-io/android-client-sdk/blob/main/bucketeer/proguard-rules.pro) into your proguard file.

:::

### Implementing dependency

Implement the dependency in your Gradle file. Please refer to the [SDK releases page](https://github.com/bucketeer-io/android-client-sdk/releases) to find the latest version.

<Tabs>
<TabItem value="gradle" label="Gradle">

```groovy showLineNumbers
dependencies {
  implementation 'io.bucketeer:android-client-sdk:LATEST_VERSION'
}
```

</TabItem>
</Tabs>

### Importing client

Import the Bucketeer client into your application code.

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
import io.bucketeer.sdk.android.*
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
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
val config = BKTConfig.builder()
  .apiKey("YOUR_API_KEY")
  .apiEndpoint("YOUR_API_ENDPOINT")
  .featureTag("YOUR_FEATURE_TAG")
  .build()

val user = BKTUser.builder()
  .id("USER_ID")
  .build()
```

</TabItem>
</Tabs>

:::info Custom configuration

Depending on your use, you may want to change the optional configurations available in the **BKTConfig.Builder**.

- **pollingInterval** - Minimum 60 seconds. Default is 10 minutes (In Milliseconds)
- **backgroundPollingInterval** - Minimum 20 minutes. Default is 1 hour (In Milliseconds)
- **eventsFlushInterval** - Minimum 60 seconds. Default is 60 seconds (In Milliseconds)
- **eventsMaxQueueSize** - Default is 50 events

:::

:::note

The Bucketeer SDK doesn't save the user data. The Application must save and set it when initializing the client SDK.

:::

### Initializing client

Initialize the client by passing the configurations in the previous step.

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
BKTClient.initialize(this.application, config, user)
val client = BKTClient.getInstance()
```

</TabItem>
</Tabs>

:::info Default timeout

The initialize process default timeout is **5 seconds**.<br />
Once initialization is finished, all the requests in the SDK use a timeout of **30 seconds**.

:::

If you want to use the feature flag on Splash or Main views, the SDK cache may be old or not exist and may not have enough time to fetch the variations from the Bucketeer server. In this case, we recommend using the `Future<BKTException?>` returned from the initialize method. In addition, you can define a custom timeout.

:::info Initialization Timeout error

During the initialization process, errors **are not** related to the initialization itself. Instead, they arise from a timeout request, indicating the variations data from the server weren't received. Therefore, the SDK will work as usual and update the variations in the next [polling](android#polling) request.

:::

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
val timeout = 2000 // Default is 5 seconds (In milliseconds)

viewLifecycleOwner.lifecycleScope.launch {
  val future = BKTClient.initialize(this.application, config, user, timeout)

  // Future is blocking, so you need to wait on non-main thread.
  val error = withContext(Dispatchers.IO) {
    future.get()
  }

  // Future returns null if BKTClient successfully fetched evaluations.
  if (error == null) {
    val showNewFeature = BKTClient.getInstance().booleanVariation("YOUR_FEATURE_FLAG_ID", false)
    if (showNewFeature) {
        // The Application code to show the new feature
    } else {
        // The code to run when the feature is off
    }
  } else {
    // Handle the error when there is no cache or the cache is not updated
  }
}
```

</TabItem>
</Tabs>

#### Polling

The initialize process immediately starts polling the latest evaluations from the Bucketeer server in the background using the interval `pollingInterval` configuration while the application is in the **foreground state**.
When the application changes to the **background state**, it will use the `backgroundPollingInterval` configuration.

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

## Supported features

### Evaluating user

The variation method determines whether or not a feature flag is enabled for a specific user.<br />
You can use the client like the one below to check which variation a specific user will receive.

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
val showNewFeature = client.booleanVariation("YOUR_FEATURE_FLAG_ID", false)
if (showNewFeature) {
    // The Application code to show the new feature
} else {
    // The code to run when the feature is off
}
```

:::note

The variation method will return the default value if the feature flag is missing in the SDK.

:::

</TabItem>
</Tabs>

### Variation types

The Bucketeer SDK supports the following variation types.

:::caution Deprecated

The `jsonVariation` interface is deprecated. Please use the `objectVariation` instead.

:::

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
fun booleanVariation(featureId: String, defaultValue: Boolean): Boolean

fun stringVariation(featureId: String, defaultValue: String): String

fun intVariation(featureId: String, defaultValue: Int): Int

fun doubleVariation(featureId: String, defaultValue: Double): Double

fun objectVariation(featureId: String, defaultValue: BKTValue): BKTValue
```

</TabItem>
</Tabs>

### Getting evaluation details

The following methods will return the **evaluation details** for a specific feature flag. If the feature flag is missing in the SDK's cache, the variable `reason` value will be `CLIENT`, which means the default value was returned.

This is useful if you use another A/B Test solution with Bucketeer and need to know the variation name, reason, and other information.

:::caution Deprecated

The `evaluationDetails` interface is deprecated. Please use the following [interfaces](#interface).

:::


#### Interface

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
fun boolVariationDetails(
  featureId: String,
  defaultValue: Boolean,
): BKTEvaluationDetails<Boolean>

fun stringVariationDetails(
  featureId: String,
  defaultValue: String,
): BKTEvaluationDetails<String>

fun intVariationDetails(
  featureId: String,
  defaultValue: Int,
): BKTEvaluationDetails<Int>

fun doubleVariationDetails(
  featureId: String,
  defaultValue: Double,
): BKTEvaluationDetails<Double>

fun objectVariationDetails(
  featureId: String,
  defaultValue: BKTValue,
): BKTEvaluationDetails<BKTValue>
```

</TabItem>
</Tabs>

#### Object

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
data class BKTEvaluationDetails<T>(
  val featureId: String,
  val featureVersion: Int,
  val userId: String,
  val variationId: String,
  val variationName: String,
  val variationValue: T,
  val reason: Reason,
) {
  enum class Reason {
    TARGET, // Evaluated using an Individual targeting
    RULE, // Evaluated using a custom Rule targeting
    DEFAULT, // Evaluated using the Default Strategy
    CLIENT, // The flag is missing in the cache. The default value was returned
    OFF_VARIATION, // Evaluated using the Off Variation
    PREREQUISITE, // Evaluated using a Prerequiste targeting

    ;
  }
}
```

</TabItem>
</Tabs>

#### Usage

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
val showNewFeature = client.boolVariationDetails("YOUR_FEATURE_FLAG_ID", false)
if (showNewFeature.variationValue) {
    // The Application code to show the new feature
} else {
    // The code to run when the feature is off
}
```

</TabItem>
</Tabs>

### Updating user evaluations

Depending on the use case, you may need to ensure the evaluations in the SDK are up to date before requesting the variation.

The fetch method uses the following parameter and returns a `Future<BKTExeptIon?>`.

- **Timeout** - Default is 30 seconds (In milliseconds)

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
val timeout = 5000 // Optional. Default is 30 seconds (In milliseconds)
val future = client.fetchEvaluations(timeout)

// Future is blocking, avoid waiting it on the main thread.
val error = future.get()
if (error == null) {
  val showNewFeature = client.booleanVariation("YOUR_FEATURE_FLAG_ID", false)
  if (showNewFeature) {
      // The Application code to show the new feature
  } else {
      // The code to run when the feature is off
  }
} else {
  // Handle the error
}
```

</TabItem>
</Tabs>

:::caution

Depending on the client network, it could take a few seconds until the SDK fetches the data from the server, so use this carefully.

You don't need to call this method manually in regular use because the SDK is polling the latest variations in the background.

:::

### Updating user evaluations in real-time

Bucketeer SDK supports FCM ([Firebase Cloud Messaging](https://firebase.google.com/docs/cloud-messaging)).<br />
Every time you change a feature flag on the admin console, Bucketeer will send a silent notification using the FCM V1 API to notify the client so that you can update the evaluations in real-time.

:::info Before using
1. You need to register your FCM Service Account on the admin console. Check the [Pushes](/organization-settings/pushes) section to learn how to do it.
2. This feature may not work if the end-user has the notification disabled.
:::

:::info Subscription Topic
When using the FCM integration, your applications must subscribe to Bucketeer's topic, so they can receive notifications.<br />

The topic varies depending on the feature flag tag.<br />
E.g.: **bucketeer-\<YOUR_FEATURE_FLAG_TAG\>**

The tag in the topic is the same tag used when initializing the client SDK.<br />
If you have a flag using the `android` tag, the topic will be **bucketeer-android**.

**Please be aware that the tag is case-sensitive.**

**No Tags:** If you're not using tags, subscribe to the topic: **bucketeer-default**
:::

Assuming you already have the FCM implementation in your application.

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
// In order to receive notifications you must subscribe to the topic
private fun subscribeToTopic() {
  val tag = getTag() // The same tag used when initializing the client SDK
  Firebase.messaging
    .subscribeToTopic("bucketeer-$tag")
    .addOnCompleteListener { task ->
      var msg = "Subscribed successfully"
      if (!task.isSuccessful) {
        msg = "Failed to subscribe"
      }
      Log.d(TAG, msg)
    }
}

override fun onMessageReceived(remoteMessage: RemoteMessage?) {
  remoteMessage?.data?.also { data ->
    val isFeatureFlagUpdated = data["bucketeer_feature_flag_updated"]
    if (isFeatureFlagUpdated == "true") {
      // The callback will return without waiting until the fetching variation process finishes
      val timeout = 1000 // Default is 5 seconds

      val future = client.fetchEvaluations(timeout)
      val error = future.get()
      if (error == null) {
        val showNewFeature = client.booleanVariation("YOUR_FEATURE_FLAG_ID", false)
        if (showNewFeature) {
            // The Application code to show the new feature
        } else {
            // The code to run when the feature is off
        }
      } else {
        // Handle the error
      }
    }
  }
}
```

</TabItem>
</Tabs>

### Reporting custom events

This method lets you save user actions in your application as events. You can connect these events to metrics in the experiments console UI.

In addition, you can pass a double value to the goal event. These values will sum and show as <br />`Value total` on the experiments console UI. This is useful if you have a goal event for tracking how much a user spent on your application buying items, etc.

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
client.track("YOUR_GOAL_ID", 10.50)
```

</TabItem>
</Tabs>

### Flushing events

This method will send all pending analytics events to the Bucketeer server immediately. This process is asynchronous, but the method returns `Future<BKTExeptIon?>` if you want to wait for its completion before doing something else.

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
val future = client.flush()
```

</TabItem>
</Tabs>

:::note

In regular use, you don't need to call the flush method because the events are sent every **60 seconds** in the background.

:::

### User attributes configuration

This feature will give you robust and granular control over what users can see on your application. You can add rules using these attributes on the console UI's feature flag's targeting tab. [See more](#).

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
val attributes = mapOf(
  "app_version" to "1.0.0",
  "os_version" to "11.0.0",
  "device_model" to "pixel-5"
  "language" to "english",
  "genre" to "female",
)

val user = BKTUser.builder()
  .id("USER_ID")
  .customAttributes(attributes)
  .build()

BKTClient.initialize(this.application, config, user)
val client = BKTClient.getInstance()
```

</TabItem>
</Tabs>

### Updating user attributes

This method will update all the current user attributes. This is useful in case the user attributes update dynamically on the application after initializing the SDK.

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
val attributes = mapOf(
  "app_version" to "1.0.1",
  "os_version" to "11.0.0",
  "device_model" to "pixel-5"
  "language" to "english",
  "genre" to "female",
  "country" to "japan",
)

client.updateUserAttributes(attributes)
```

</TabItem>
</Tabs>

:::caution

This updating method will override the current data.

:::

### Getting user information

This method will return the current user configured in the SDK. This is useful when you want to check the current user id and attributes before updating them through [updateUserAttributes](#updating-user-attributes).

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
val user = client.currentUser()
```

</TabItem>
</Tabs>


### Listening to evaluation updates

The SDK can notify when the evaluation is updated.
The listener can detect both automatic polling and manual fetching.

:::info

The listener callback is called on the main thread.

:::

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
// Returned value is used when you want to remove listener
val key = client.addEvaluationUpdateListener {
  // The listener callback is called on the main thread
  val showNewFeature = client.booleanVariation("YOUR_FEATURE_FLAG_ID", false)
  if (showNewFeature) {
      // The Application code to show the new feature
  } else {
      // The code to run when the feature is off
  }
}

// Remove a listener associated with the key
client.removeEvaluationUpdateListener(key)

// Remove all listeners
client.clearEvaluationUpdateListeners()
```

</TabItem>
</Tabs>

### Destroying client

There are cases you might want to switch the user ID or reduce resources when the application is in the background.<br />
For those cases, you can call the destroy interface, which will clear the client instance.

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
client.destroy()
```

</TabItem>
</Tabs>

:::tip

If you want to switch the user ID, please call the [flush](#flushing-events) interface before calling the destroy, so that all the pending events can be sent before clearing the client instance, then call the [initialize](#initializing-client) interface with the new user information.

:::
