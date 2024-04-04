---
title: Flutter
slug: /sdk/client-side/flutter
toc_max_heading_level: 4
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

This category contains topics explaining how to configure Bucketeer's Flutter SDK.

:::caution Flutter SDK Version (Beta)

The Flutter SDK is currently in its Beta stage.

If you find any issues or have suggestions for improvement, feel free to open an [issue](https://github.com/bucketeer-io/flutter-client-sdk/issues).<br />
The SDK doesn't support Flutter Web yet. All contributions are welcome!

:::

## Getting started

Before starting, ensure that you follow the [Getting Started](/getting-started) guide.

### Implementing dependency

Implement the dependency in your Pubspec file. Please refer to the [SDK releases page](https://github.com/bucketeer-io/flutter-client-sdk/releases) to find the latest version.

<Tabs>
<TabItem value="yaml" label="Pubspec">

```yaml showLineNumbers
dependencies:
  bucketeer_flutter_client_sdk: LATEST_VERSION
```

</TabItem>
</Tabs>

### Importing client

Import the Bucketeer client into your application code.

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
import 'package:bucketeer_flutter_client_sdk/bucketeer_flutter_client_sdk.dart';
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
<TabItem value="dart" label="Dart">

```dart showLineNumbers
final config = BKTConfigBuilder()
  .apiKey("YOUR_API_KEY")
  .apiEndpoint("YOUR_API_URL")
  .featureTag("YOUR_FEATURE_TAG")
  .appVersion("YOUR_APP_VERSION")
  .build();

final user = BKTUserBuilder
  .id("USER_ID")
  .customAttributes(YOUR_USER_ATTRIBUTES) /// optional Map<String, String>
  .build();
```

</TabItem>
</Tabs>

:::info Custom configuration

Depending on your use, you may want to change the optional configurations available in the **BKTConfigBuilder**.

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
<TabItem value="dart" label="Dart">

```dart showLineNumbers
await BKTClient.initialize(config: config, user: user);
```

</TabItem>
</Tabs>

:::info Default timeout

The initialize process default timeout is **5 seconds**.<br />
Once initialization is finished, all the requests in the SDK use a timeout of **30 seconds**.

:::

If you want to use the feature flag on Splash or Main views, the SDK cache may be old or not exist and may not have enough time to fetch the variations from the Bucketeer server. In this case, we recommend using `await` in the initialize method. In addition, you can define a custom timeout.

:::info Initialization Timeout error

During the initialization process, errors **are not** related to the initialization itself. Instead, they arise from a timeout request, indicating the variations data from the server weren't received. Therefore, the SDK will work as usual and update the variations in the next [polling](flutter#polling) request.

:::

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
/// It will unlock without waiting until the fetching variation process finishes
const int timeout = 2000; /// Default is 5 seconds (In milliseconds)
final result = await BKTClient.initialize(
    config: config, user: user, timeoutMillis: timeout);
if (result.isSuccess) {
  const client = BKTClient.instance;
  final showNewFeature = await client
      .boolVariation("YOUR_FEATURE_FLAG_ID", defaultValue: false);
  if (showNewFeature) {
    /// The Application code to show the new feature
  } else {
    /// The code to run if the feature is off
  }
} else {
  /// Handle the error when there is no cache or the cache is not updated
  if (result.asFailure.exception is BKTTimeoutException) {
    /// Handle timeout error
  } else {
    /// Anything else
  }
}
```

</TabItem>
</Tabs>

#### Polling

The initialize process immediately starts polling the latest evaluations from the Bucketeer server in the background using the interval `pollingInterval` configuration while the application is in the **foreground state**.
When the application changes to the **background state**, it will use the `backgroundPollingInterval` configuration.

#### Polling retry behavior

The Bucketeer SDK regularly polls the latest evaluations from the server based on the `pollingInterval` parameter. By default, the `pollingInterval` is set to 10 minutes, but you can adjust it to suit your needs.

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
To check which variation a specific user will receive, you can use the client like below.

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
final showNewFeature =
    await client.boolVariation("YOUR_FEATURE_FLAG_ID", false);
if (showNewFeature) {
  /// The Application code to show the new feature
} else {
  /// The code to run if the feature is off
}
```

</TabItem>
</Tabs>

:::caution

In case the feature flag is missing in the SDK, it will return the default value.

:::

### Variation types

The Bucketeer SDK supports the following variation types.

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
Future<bool> boolVariation(String featureId, { required bool defaultValue });

Future<String> stringVariation(String featureId, { required String defaultValue });

Future<int> intVariation(String featureId, { required int defaultValue });

Future<double> doubleVariation(String featureId, { required double defaultValue });

Future<Map<String, dynamic>> jsonVariation(String featureId, { required Map<String, dynamic> defaultValue });
```

</TabItem>
</Tabs>

### Updating user evaluations

Depending on the use case, you may need to ensure the evaluations in the SDK are up to date before requesting the variation.

The fetch method uses the following parameter. Make sure to wait for its completion.

- **Timeout** - Default is 30 seconds (In milliseconds)

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
/// It will unlock without waiting until the fetching variation process finishes
int timeout = 5000; /// Optional. Default is 30 seconds (In milliseconds)

final result = await client.fetchEvaluations(timeoutMillis: timeout);
if (result.isSuccess) {
  final showNewFeature = await client
      .boolVariation("YOUR_FEATURE_FLAG_ID", defaultValue: false);
  if (showNewFeature) {
    /// The Application code to show the new feature
  } else {
    /// The code to run if the feature is off
  }
} else {
  /// Handle the error when the cache is not updated
  if (result.asFailure.exception is BKTTimeoutException) {
    /// Handle timeout error
  } else {
    /// Anything else
  }
}
```

</TabItem>
</Tabs>

:::caution

Depending on the client network, it could take a few seconds until the SDK fetches the data from the server, so use this carefully.

You don't need to call this method manually in regular use because the SDK is polling the latest variations in the background.

:::

### Updating user evaluations in real-time

The Bucketeer SDK supports FCM ([Firebase Cloud Messaging](https://firebase.google.com/docs/cloud-messaging)).<br />
Every time you change a feature flag on the admin console, Bucketeer will send notifications using the FCM API to notify the client so that you can update the evaluations in real-time.

:::note

1. You need to register your FCM API Key on the admin console. Check the [Pushes](/integration/pushes) section to learn how to do it.
2. This feature may not work if the user has the notification disabled.

:::

Assuming you already have the FCM implementation in your application.

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
FirebaseMessaging.onMessage.listen((RemoteMessage message) async {
  final isFeatureFlagUpdated = message.data["bucketeer_feature_flag_updated"]
    if (isFeatureFlagUpdated) {
      int timeout = 1000;
      const client = BKTClient.instance;
      final result = await client.fetchEvaluations(timeoutMillis: timeout);
      if (result.isSuccess) {
        final showNewFeature = await client
            .boolVariation("YOUR_FEATURE_FLAG_ID", defaultValue: false);
        if (showNewFeature) {
          /// The Application code to show the new feature
        } else {
          /// The code to run if the feature is off
        }
      } else {
        /// Handle the error when the cache is not updated
        if (result.asFailure.exception is BKTTimeoutException) {
          /// Handle timeout error
        } else {
          /// Anything else
        }
      }
    }
});
```

</TabItem>
</Tabs>

### Reporting custom events

This method lets you save user actions in your application as events. You can connect these events to metrics in the experiments console UI.

In addition, you can pass a double value to the goal event. These values will sum and show as <br />`Value total` on the experiments console UI. This is useful if you have a goal event for tracking how much a user spent on your application buying items, etc.

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
final result = await client.track("YOUR_GOAL_ID", value: 10.50);
if (result.isSuccess) {
  /// The Application code to run after track custom event
} else {
  /// Handle the error
}
```

</TabItem>
</Tabs>

### Flushing events

This method will send all pending analytics events to the Bucketeer server immediately. This process is asynchronous, so it returns before it is complete.

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
await client.flush();
```

</TabItem>
</Tabs>

:::note

In regular use, you don't need to call the flush method because the events are sent every **60 seconds** in the background.

:::

### User attributes configuration

This feature will give you robust and granular control over what users can see on your application. You can add rules using these attributes on the console UI's feature flag's targeting tab. [See more](#).

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
final attributes = {
  'app_version': '1.0.0',
  'os_version': '11.0.0',
  'device_model': 'pixel-5'
  'language': 'english',
  'genre': 'female'
};

final user = BKTUserBuilder()
  .id("USER_ID")
  .customAttributes(attributes)
  .build();

await BKTClient.initialize(config: config, user: user);
```

</TabItem>
</Tabs>

### Updating user attributes

This method will update all the current user attributes. This is useful in case the user attributes update dynamically on the application after initializing the SDK.

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
final attributes = {
  'app_version': '1.0.1',
  'os_version': '11.0.0',
  'device_model': 'pixel-5'
  'language': 'english',
  'genre': 'female'
  'country': 'japan'
};

final result = await client.updateUserAttributes(attributes);
if (result.isSuccess) {
  /// The Application code to run after update the user attributes
} else {
   /// Handle the error
}
```

</TabItem>
</Tabs>

:::caution

This updating method will override the current data.

:::

### Getting user information

This method will return the current user configured in the SDK. This is useful when you want to check the current user id and attributes before updating them through [updateUserAttributes](#getting-user-information).

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
final userResult = await client.currentUser();
if (userResult.isSuccess) {
  final user = userResult.asSuccess.data;
  /// The Application code to run after get the user
} else {
  /// Handle the error
}
```

</TabItem>
</Tabs>

### Getting evaluation details

This method will return the **evaluation details** for a specific feature flag or will return **null** if the feature flag is missing in the SDK's cache.

This is useful if you use another A/B Test solution with Bucketeer and need to know the variation name, reason, and other information.

<details>
  <summary><strong>Evaluation details</strong></summary>
  <Tabs>
  <TabItem value="dart" label="Dart">

  ```dart showLineNumbers
  class BKTEvaluation {
    final String id;
    final String featureId;
    final int featureVersion;
    final String userId;
    final String variationId;
    final String variationName;
    final String variationValue;
    final String reason;
  }
  ```

  </TabItem>
  </Tabs>
</details>

:::caution

Do not call this method without calling the [Evaluating user method](#evaluating-user). The Evaluating user method must always be called because it generates analytics events that will be sent to the server.

:::

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
final evaluationDetails = await client.evaluationDetails("YOUR_FEATURE_FLAG_ID");
```

</TabItem>
</Tabs>


### Listening to evaluation updates

The SDK can notify when the evaluation is updated.
The listener can detect both automatic polling and manual fetching.

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
class EvaluationUpdateListenerImpl implements BKTEvaluationUpdateListener {
  @override
  void onUpdate() async {
    const client = BKTClient.instance;
    final showNewFeature =
        await client.boolVariation("YOUR_FEATURE_FLAG_ID", defaultValue: false);
    if (showNewFeature) {
      /// The Application code to show the new feature
    } else {
      /// The code to run when the feature is off
    }
  }
}

final listener = EvaluationUpdateListenerImpl();
const client = BKTClient.instance;
/// The returned key value is used to remove the listener
final key = await client.addEvaluationUpdateListener(listener);
/// Remove a listener associated with the key
client.removeEvaluationUpdateListener(key);
/// Remove all listeners
client.clearEvaluationUpdateListeners();
```

</TabItem>
</Tabs>

### Background mode

For Android, the background mode works without any configuration by default.<br />
For iOS, this feature is optional, but it can be used by configuring the background mode in the Xcode project.<br />
The default setting is **1 hour** and a minimum of **20 minutes** for the `backgroundPollingInterval`, and **1 minute** for the `eventsFlushInterval` settings.

#### Configuration for iOS
**1.** Open the iOS folder using Xcode IDE to open the file `Runner.xcworkspace`.<br />
**2.** Open the XCode project setting.<br />
**3.** Select the `Runner` target.<br />
**4.** Select the `Signing & Capabilities` settings tab.<br />
**5.** Enable the `Background processing` option in the `Background Modes` section.<br />
**6.** Register the identifier by adding the `Permitted background task scheduler identifiers` option in the **Info.plist**, and set the following task IDs as the values.

-  **io.bucketeer.background.fetch.evaluations** (Background task to fetch the latest evaluations from the server).
-  **io.bucketeer.background.flush.events** (Background task to flush events generated by the client).

**7.** Open the `AppDelegate.swift` and add the following code to the `didFinishLaunchingWithOptions` function.

```swift showLineNumbers
import UIKit
import Bucketeer

class AppDelegate: UIResponder, UIApplicationDelegate {

    func application(_ application: UIApplication, didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]?) -> Bool {
        // Add the code to enable background tasks
        if #available(iOS 13.0, tvOS 13.0, *) {
            BKTBackgroundTask.enable()
        }
        // Your app logic code
        return true
    }
```

**8.** If needed, change the background polling default interval by adding the `backgroundPollingInterval` setting in the **BKTConfigBuilder** as follows:

```dart showLineNumbers
final config = BKTConfigBuilder()
  .apiKey("YOUR_API_KEY")
  .apiEndpoint("YOUR_API_URL")
  .featureTag("YOUR_FEATURE_TAG")
  .appVersion("YOUR_APP_VERSION").
  .eventsFlushInterval(60) /// 1 minute
  .backgroundPollingInterval(1800)
  .build(); /// 30 minutes
```

**9.** Please check the [iOS Documentation](https://developer.apple.com/documentation/uikit/app_and_environment/scenes/preparing_your_ui_to_run_in_the_background/using_background_tasks_to_update_your_app) for more details.

### Destroying client

There are cases you might want to switch the user ID or reduce resources when the application is in the background.<br />
For those cases, you can call the destroy interface, which will clear the client instance.

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
await client.destroy()
```

</TabItem>
</Tabs>

:::tip

If you want to switch the user ID, please call the [flush](#flushing-events) interface before calling the destroy, so that all the pending events can be sent before clearing the client instance, then call the [initialize](#initializing-client) interface with the new user information.

:::
