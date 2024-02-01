---
title: iOS
slug: /sdk/client-side/ios
toc_max_heading_level: 4
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

This category contains topics explaining how to configure Bucketeer's iOS SDK.

:::tip iOS SDK Version (Stable)

Bucketeer iOS SDK has reached the production stage, offering you a stable and reliable experience.

:::

:::info Compatibility

The Bucketeer iOS SDK is compatible with iOS versions 11.0 and higher.

:::

## Getting started

Before starting, ensure that you follow the [Getting Started](/getting-started) guide.

### Implementing dependency

To find the latest version, refer to the [SDK releases page](https://github.com/bucketeer-io/ios-client-sdk/releases).

<Tabs>
<TabItem value="swift" label="Cocoapods">

Implement the SDK in your **Podfile** file.

```swift showLineNumbers
use_frameworks!

target 'YOUR_TARGET_NAME' do
  pod 'Bucketeer', 'LATEST_VERSION'
end
```

</TabItem>
<TabItem value="spm" label="Swift Package Manager">

Implement the SDK as a dependency in your **Package.swift** file or through Xcode.

```swift showLineNumbers
// Package.swift
dependencies: [
  .package(url: "https://github.com/bucketeer-io/ios-client-sdk.git", exact: "LATEST_VERSION"),
],
targets: [
  .target(
    name: "YOUR_TARGET_NAME",
    dependencies: [.product(name: "Bucketeer", package: "ios-client-sdk")],
  )
],
```

To include a package dependency in your Xcode project, follow these steps:

Go to **File** -> **Swift Packages** -> **Add Package Dependency**. Next, enter the clone URL of the [iOS SDK repository](https://github.com/bucketeer-io/ios-client-sdk), and specify the version you desire.

</TabItem>
<TabItem value="carthage" label="Carthage">

Implement the SDK in your **Cartfile** file.

```swift showLineNumbers
github "bucketeer-io/ios-client-sdk" ~> LATEST_VERSION
```

</TabItem>
</Tabs>

### Importing client

Import the Bucketeer client into your application code.

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
import Bucketeer
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
<TabItem value="swift" label="Swift">

```swift showLineNumbers
do {
  // SDK configuration
  let config = try BKTConfig(
    apiKey: "YOUR_API_KEY",
    apiEndpoint: "YOUR_API_ENDPOINT",
    featureTag: "YOUR_FEATURE_TAG",
    appVersion: Bundle.main.infoDictionary?["CFBundleShortVersionString"] as! String
  )
  // User configuration
  let user = try BKTUser(id: "USER_ID")
} catch {
  // Error handling
}
```

</TabItem>
</Tabs>

:::info Custom configuration

Depending on your use, you may want to change the optional configurations available in the **BKTConfig**.

- **pollingInterval** (Minimum 60 seconds. Default is 10 minutes)
- **backgroundPollingInterval** (Minimum 20 minutes. Default is 1 hour)
- **eventsFlushInterval** (Minimum 60 seconds. Default is 60 seconds)
- **eventsMaxQueueSize** (Default is 50 events)

:::

:::note

The Bucketeer SDK doesn't save the user data. The Application must save and set it when initializing the client SDK.

:::

### Initializing client

Initialize the client by passing the configurations in the previous step.

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
BKTClient.initialize(config: config, user: user)
```

</TabItem>
</Tabs>

:::info Default timeout

The initialize process default timeout is 5 seconds.

Once initialization is finished, all the requests in the SDK use a timeout of 30 seconds.

:::

If you want to use the feature flag on Splash or Main views, the SDK cache may be old or not exist and may not have enough time to fetch the variations from the Bucketeer server. For this case, we recommend using the callback in the initialize method. In addition, you can define a custom timeout.

:::caution Initialization Timeout error

During the initialization process, errors **are not** related to the initialization itself. Instead, they arise from a timeout request, indicating the variations data from the server weren't received. Therefore, the SDK will work as usual and update the variations in the next [polling](ios#polling) request.

:::


<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
let timeout: Int64 = 2000 // Default is 5 seconds

BKTClient.initialize(
  config: config,
  user: user,
  timeoutMillis: timeout
) { error in
  guard error == nil else {
    // Handle the error when there is no cache or the cache is not updated
    return
  }
  let client = BKTClient.shared
  let showNewFeature = client.boolVariation(featureId: "YOUR_FEATURE_FLAG_ID", defaultValue: false)
  if (showNewFeature) {
    // The Application code to show the new feature
  } else {
    // The code to run when the feature is off
  }
}
```

</TabItem>
</Tabs>


#### Polling

The initialize process immediately starts polling the latest evaluations from the Bucketeer server in the background using the interval `pollingInterval` configuration while the application is in the **foreground state**. When the application changes to the **background state**, it will use the `backgroundPollingInterval` configuration when the [Background fetch](/sdk/client-side/ios#background-fetch) is configured.

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
To check which variation a specific user will receive, you can use the client like below.

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
let client = BKTClient.shared
let showNewFeature = client.boolVariation(featureId: "YOUR_FEATURE_FLAG_ID", defaultValue: false)
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
<TabItem value="swift" label="Swift">

```swift showLineNumbers
func boolVariation(featureId: String, defaultValue: Bool) -> Bool

func stringVariation(featureId: String, defaultValue: String) -> String

func intVariation(featureId: String, defaultValue: Int) -> Int

func doubleVariation(featureId: String, defaultValue: Double) -> Double

func jsonVariation(featureId: String, defaultValue: [String: AnyHashable]) -> [String: AnyHashable]
```

</TabItem>
</Tabs>

### Updating user evaluations

Depending on the use case, you may need to ensure the evaluations in the SDK are up to date before requesting the variation.

The fetch method uses the following parameters.

- **Completion callback**
- **Timeout** (Default is 30 seconds)

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
let timeout: Int64 = 5000
let client = BKTClient.shared

client.fetchEvaluations(timeoutMillis: timeout) { error in
  guard error == nil else {
    // The code to run when there is an error while initializing the SDK
    return
  }
  let showNewFeature = client.boolVariation(featureId: "YOUR_FEATURE_FLAG_ID", defaultValue: false)
  if (showNewFeature) {
      // The Application code to show the new feature
  } else {
      // The code to run when the feature is off
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

Assuming you already have the FCM implementation in your application, the following example shows how to handle silent push notifications.

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
// Receiving notification in background
func application(
  _ application: UIApplication,
  didReceiveRemoteNotification userInfo: [AnyHashable: Any]
) async -> UIBackgroundFetchResult {
  let flag = userInfo["bucketeer_feature_flag_updated"] as? String
  if flag == "true" {
    let client = BKTClient.shared
    let showNewFeature = client.boolVariation(featureId: "YOUR_FEATURE_FLAG_ID", defaultValue: false)
    if (showNewFeature) {
        // The Application code to show the new feature
    } else {
        // The code to run when the feature is off
    }
  }
  return UIBackgroundFetchResult.newData
}
```

</TabItem>
</Tabs>

### Reporting custom events

This method lets you save user actions in your application as events. You can connect these events to metrics in the experiments console UI.

In addition, you can pass a double value to the goal event. These values will sum and show as <br />`Value total` on the experiments console UI. This is useful if you have a goal event for tracking how much a user spent on your application buying items, etc.

The default track value is 0.0.

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
client.track(goalId: "YOUR_GOAL_ID", value: 10.50)
```

</TabItem>
</Tabs>

### Flushing events

This method will send all pending analytics events to the Bucketeer server immediately. This process is asynchronous, so it returns before it is complete.

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
client.flush()
```

</TabItem>
</Tabs>

:::note

In regular use, you don't need to call the flush method because the events are sent every 30 seconds in the background.

:::

### User attributes configuration

This feature will give you robust and granular control over what users can see on your application. You can add rules using these attributes on the console UI's feature flag's targeting tab. [See more](#).

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
do {
  // SDK configuration
  let config = try BKTConfig(
    apiKey: "YOUR_API_KEY",
    apiEndpoint: "YOUR_API_ENDPOINT",
    featureTag: "ios",
    appVersion: Bundle.main.infoDictionary?["CFBundleShortVersionString"] as! String
  )
  // User attributes configuration
  let attributes = [
    "app_version": "1.0.0",
    "os_version": "11.0.0",
    "device_model": "iphone-12",
    "language": "english",
    "genre": "female"
  ]

  let user = try BKTUser(id: "USER_ID", attributes: attributes)
  BKTClient.initialize(config: config, user: user)
  let client = BKTClient.shared
} catch {
  // Error handling
}
```

</TabItem>
</Tabs>

### Updating user attributes

This method will update all the current user attributes. This is useful in case the user attributes update dynamically on the application after initializing the SDK.

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
let attributes = [
  "app_version": "1.0.1",
  "os_version": "11.0.0",
  "device_model": "iphone-12",
  "language": "english",
  "genre": "female",
  "country": "japan"
]

client.updateUserAttributes(attributes: attributes)
```

</TabItem>
</Tabs>

:::caution

This updating method will override the current data.

:::

### Getting user information

This method will return the current user configured in the SDK. This is useful when you want to check the current user id and attributes before updating them through [updateUserAttributes](#getting-user-information).

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
let user = client.currentUser()
```

</TabItem>
</Tabs>

### Getting evaluation details

This method will return the evaluation details for a specific feature flag. This is useful if you need to know the variation reason or send this data to your own analysis database.
It will return null if the feature flag is missing in the SDK.

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
let evaluationDetails = client.evaluationDetails(featureId: "YOUR_FEATURE_FLAG_ID")
```

:::caution

Do not call this method without calling the [Evaluating user method](#evaluating-user). The Evaluating user method must always be called because it generates analytics events that will be sent to the server.

:::

</TabItem>
</Tabs>

### Listening to evaluation updates

BKTClient can notify when the evaluation is updated.
The listener can detect both automatic polling and manual fetching.

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
class EvaluationUpdateListenerImpl: EvaluationUpdateListener {
  func onUpdate() {
    let client = BKTClient.shared
    let showNewFeature = client.boolVariation(featureId: "YOUR_FEATURE_FLAG_ID", defaultValue: false)
    if (showNewFeature) {
      // The Application code to show the new feature
    } else {
      // The code to run when the feature is off
    }
  }
}
let listener = EvaluationUpdateListenerImpl()
// The returned key value is used to remove the listener
let key = client.addEvaluationUpdateListener(listener: listener)

// Remove a listener associated with the key
client.removeEvaluationUpdateListener(key: key)

// Remove all listeners
client.clearEvaluationUpdateListeners()
```

</TabItem>
</Tabs>

### Background mode

This feature is optional but can be used by configuring the Background Mode in your application.<br />
The default setting is **1 hour** and a minimum of **20 minutes** for the `backgroundPollingInterval` and **1 minute** for the `eventsFlushInterval` settings.

#### Configuration
**1.** Open the XCode project setting.<br />
**2.** Select your app target.<br />
**3.** Select the `Signing & Capabilities` settings tab.<br />
**4.** Enable the `Background processing` option in the `Background Modes` section.<br />
**5.** Register the identifier by adding the `Permitted background task scheduler identifiers` option in the **Info.plist**, and set the following task IDs as the values.

-  **io.bucketeer.background.fetch.evaluations** (Background task to fetch the latest evaluations from the server).
-  **io.bucketeer.background.flush.events** (Background task to flush events generated by the client).

**6.** Add the code to enable the background task.
<Tabs>
<TabItem value="Storyboard" label="Storyboard">

Open the `AppDelegate.swift` and add the following code to the `didFinishLaunchingWithOptions` function.

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

</TabItem>

<TabItem value="SwiftUI" label="SwiftUI">

Open the your `App.swift` file and add the following code.

```swift showLineNumbers
import SwiftUI
import Bucketeer

@main
struct ExampleSwiftUIApp: App {
    var body: some Scene {
        WindowGroup {
            ContentView()
        }
        // Add the code to enable background tasks
        .enableBKTBackgroundTask()
    }
}

```
</TabItem>

</Tabs>

**7.** If needed, change the background polling default interval by adding the `backgroundPollingInterval` setting in the **BKTConfig** as follows:

```swift showLineNumbers
do {
  // SDK configuration
  let config = try BKTConfig(
    apiKey: "YOUR_API_KEY",
    apiEndpoint: "YOUR_API_ENDPOINT",
    featureTag: "YOUR_FEATURE_TAG",
    appVersion: Bundle.main.infoDictionary?["CFBundleShortVersionString"] as! String,
    backgroundPollingInterval: 1800 // 30 minutes
  )
} catch {
  // Error handling
}
```

**8.** Please check the [iOS Documentation](https://developer.apple.com/documentation/uikit/app_and_environment/scenes/preparing_your_ui_to_run_in_the_background/using_background_tasks_to_update_your_app) for more details.
