---
title: iOSリファレンス
slug: /sdk/client-side/ios
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

このカテゴリには、BucketeerのiOS SDKの設定方法に関する内容が含まれています。

:::info 互換性

Bucketeer iOS SDKは、iOSバージョン11.0以降に対応しています。

:::

## はじめに

始める前に、必ず[はじめに](/getting-started)ガイドに従ってください。

### 依存関係の実装

最新バージョンについては、[SDKリリースページ](https://github.com/bucketeer-io/ios-client-sdk/releases)をご参照ください。

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

SDKを依存関係として**Package.swift**ファイルまたはXcodeを通して実装します。

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

Xcode プロジェクトにパッケージの依存関係を含めるには、以下のステップに従ってください：

**ファイル** -> **Swiftパッケージ** -> **パッケージ依存関係の追加**に移動します。次に、[iOS SDKリポジトリ](https://github.com/bucketeer-io/ios-client-sdk)のクローンURLを入力し、希望のバージョンを指定します。


</TabItem>
<TabItem value="carthage" label="Carthage">

Implement the SDK in your **Cartfile** file.

```swift showLineNumbers
github "bucketeer-io/ios-client-sdk" ~> LATEST_VERSION
```

</TabItem>
</Tabs>

### クライアントのインポート

Bucketeerクライアントをアプリケーションコードにインポートします。

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
import Bucketeer
```

</TabItem>
</Tabs>

### クライアント構成

SDK 設定とユーザー設定を構成します。

:::note

**フィーチャータグ**設定は、フィーチャーフラグを作成する際に構成するタグです。

:::

以下の例の設定はすべて必須となっています。

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

:::info カスタム構成

使用目的に応じて、**BKTConfig**で使用できるオプション構成を変更することができます。

- **pollingInterval** (最低60秒。デフォルトは10分)
- **backgroundPollingInterval** (最低20分。デフォルトは1時間)
- **eventsFlushInterval** (最低60秒。デフォルトは60秒)
- **eventsMaxQueueSize** (デフォルトは50イベント)

:::

:::note

Bucketeer SDKはユーザーデータを保存しません。アプリケーションは、クライアントSDKを初期化する際にユーザーデータを保存して設定する必要があります。

:::

### バックグラウンドモード

この機能は任意ですが、アプリケーションでバックグラウンドモードを設定することで使用できます。<br />
`backgroundPollingInterval`の設定のデフォルトは**1時間**、最低**20分**、`eventsFlushInterval`の設定のデフォルトは1分です。


#### 構成

**1.** `Signing & Capabilities`設定で`Background fetch`を有効にします。<br />
**2.** **Info.plist**に`Permitted background task scheduler identifiers`オプションを追加して識別子を登録し、以下のタスクIDを値として設定します。

-  **io.bucketeer.background.fetch.evaluations** (サーバーから最新のエバリュエーションを取得するためのバックグラウンドタスク).
-  **io.bucketeer.background.flush.events** (クライアントで生成されたイベントをフラッシュするためのバックグラウンドタスク).


**3.** 必要に応じて、**BKTConfig**に`backgroundPollingInterval`設定を追加して、バックグラウンドポーリングのデフォルト間隔を変更します。

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

### クライアントの初期化

前のステップで構成を渡してクライアントを初期化します。

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
BKTClient.initialize(config: config, user: user)
```

</TabItem>
</Tabs>

:::note

初期化プロセスは、アプリケーションが**前景状態**にある間、バックグラウンドで`pollingInterval`の設定間隔でBucketeerから最新の評価をポーリングし始めます。アプリケーションが**バックグラウンド状態**に変更されると、[バックグラウンドフェッチ](/sdk/client-side/ios#background-fetch)が設定されている場合には`backgroundPollingInterval`の設定が使用されます。

:::

スプラッシュやメイン画面でフィーチャーフラグを使用したい場合、ユーザーが初めてアプリを開くと、Bucketeerサーバーからバリエーションを取得するのに十分な時間が取れない場合があります。

このケースでは、initializeメソッドでコールバックを使用することをお勧めします。

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
// The callback will return without waiting until the fetching variation process finishes
let timeout: Int64 = 2000 // Default is 5 seconds

BKTClient.initialize(
  config: config,
  user: user,
  timeoutMillis: timeout
) { error in
    guard error == nil else {
      // The code to run when there is an error while initializing the SDK
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

## サポートされている機能

### ユーザーエバリュエーション

バリエーションメソッドは、フィーチャーフラグが特定のユーザーに対して有効であるかどうかを決定します。<br />
特定のユーザーがどのバリエーションを受け取るかを確認するには、以下のようなクライアントを使用します。


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

SDKにフィーチャーフラグがない場合、バリエーションメソッドはデフォルト値を返します。

:::

### バリエーションの種類

Bucketeer SDKは、以下のバリエーションの種類をサポートしています。

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

### ユーザーエバリュエーションの更新

ユースケースによっては、バリエーションをリクエストする前に、SDK内のエバリュエーションが最新であることを確認する必要がある場合があります。

fetchメソッドは次のパラメータを使用します。

- **完了コールバック**
- **タイムアウト** (デフォルトは30秒)

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

クライアントのネットワークによっては、SDKがサーバーからデータをフェッチするまでに数秒かかる場合がありますので、注意してください。

SDKはバックグラウンドで最新のバリエーションをポーリングしているので、通常の使用ではこのメソッドを手動で呼び出す必要はありません。

:::

### ユーザーエバリュエーションのリアルタイム更新

Bucketeer SDKはFCM ([Firebase Cloud Messaging](https://firebase.google.com/docs/cloud-messaging))をサポートしています。<br />
管理コンソールでフィーチャーフラグを変更するたびに、BucketeerはFCM APIを使用してクライアントに通知を送信し、エバリュエーションをリアルタイムで更新可能にします。

:::note

1. 管理コンソールでFCM APIキーを登録する必要があります。登録方法は、[プッシュ](/integration/pushes)セクションをご参照ください。
2. ユーザーが通知を無効にしている場合、この機能は動作しない場合があります。

:::

アプリケーションにFCM実装が既にあることを仮定し、以下の例では、サイレントプッシュ通知を処理する方法を示します。

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

### カスタムイベントの報告

このメソッドを使用すると、アプリケーション内のユーザー操作をイベントとして保存できます。これらのイベントは、エクスペリメントコンソールUIの指標に接続できます。

さらに、ゴールイベントにダブル値を渡すことができます。これらの値は合計され、エクスペリメントコンソールUIの<br />`Value total`として表示されます。これは、ユーザーがアプリケーションでアイテムを購入した金額などを追跡するためのゴールイベントがある場合に役立ちます。

デフォルトのトラック値は0.0です。

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
client.track(goalId: "YOUR_GOAL_ID", value: 10.50)
```

</TabItem>
</Tabs>

### イベントのフラッシュ

このメソッドは、保留中のすべての分析イベントを可能な限り早くBucketeerサーバーに送信します。このプロセスは非同期であるため、完了前に返されます。

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
client.flush()
```

</TabItem>
</Tabs>

:::note

通常の使用では、イベントはバックグラウンドで30秒ごとに送信されるため、flushメソッド呼び出す必要はありません。

:::

### ユーザー属性の設定

この機能を使用すると、ユーザーがアプリケーションで表示できるコンテンツを堅牢かつ詳細に制御できます。これらの属性を使用して、コンソールUIのフィーチャーフラグのターゲティングタブでルールを追加できます。[詳細を見る](/feature-flags/creating-feature-flags/targeting#ユーザー属性)

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

### ユーザー属性の更新

このメソッドは、現在のユーザー属性をすべて更新します。これは、SDKの初期化後にアプリケーションでユーザー属性が動的に更新される場合に役立ちます。

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

この更新メソッドは、現在のデータを上書きします。

:::

### ユーザー情報の取得

このメソッドは、SDKで設定されている現在のユーザーを返します。これは、[updateUserAttributes](ios#ユーザー属性の更新)を使用して現在のユーザーIDと属性を更新する前に確認したい場合に役立ちます。


<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
let user = client.currentUser()
```

</TabItem>
</Tabs>

### エバリュエーションの詳細を取得

このメソッドは、特定のフィーチャーフラグのエバリュエーションの詳細を返します。これは、バリエーションの理由を知ったり、このデータを独自の分析データベースに送信する必要がある場合に役立ちます。SDKにフィーチャーフラグがない場合はnullを返します。

<Tabs>
<TabItem value="swift" label="Swift">

```swift showLineNumbers
let evaluationDetails = client.evaluationDetails(featureId: "YOUR_FEATURE_FLAG_ID")
```

:::caution

 [ユーザーエバリュエーションメソッド](#evaluating-user)を呼び出さずにこのメソッドを呼び出さないでください。ユーザーエバリュエーションメソッドは、サーバーに送信される分析イベントを生成するため常に呼び出す必要があります。

:::

</TabItem>
</Tabs>

### エバリュエーション更新のリスニング

BKTClientは、エバリュエーションが更新された際に通知することができます。
リスナーは、自動ポーリングと手動フェッチングの両方を検出できます。


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