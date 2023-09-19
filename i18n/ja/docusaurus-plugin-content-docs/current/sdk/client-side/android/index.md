---
title: Androidリファレンス
slug: /sdk/client-side/android
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

このカテゴリには、BucketeerのAndroid SDKの設定方法に関する内容が含まれています。

:::info 互換性

Bucketeer SDKは、Androidバージョン21以降に対応しています(Android 5.0, Lollipop)。

:::

## はじめに

始める前に、必ず[はじめに](/getting-started)ガイドに従ってください。


### 依存関係の実装

依存関係をGradleファイルに実装します。最新バージョンについては、[SDKリリースページ](https://github.com/bucketeer-io/android-client-sdk/releases)をご参照ください。

<Tabs>
<TabItem value="gradle" label="Gradle">

```groovy showLineNumbers
dependencies {
  implementation 'io.bucketeer:android-client-sdk:LATEST_VERSION'
}
```

</TabItem>
</Tabs>

### クライアントのインポート

Bucketeerクライアントをアプリケーションコードにインポートします。

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
import io.bucketeer.sdk.android.*
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

:::info カスタム構成

使用目的に応じて、**BKTConfig.Builder**で使用できるオプション構成を変更することができます。

- **pollingInterval** (最低60秒。デフォルトは10分)
- **backgroundPollingInterval** (最低20分。デフォルトは1時間)
- **eventsFlushInterval** (最低60秒。デフォルトは60秒)
- **eventsMaxQueueSize** (デフォルトは50イベント)

:::

:::note

Bucketeer SDKはユーザーデータを保存しません。アプリケーションは、クライアントSDKを初期化する際にユーザーデータを保存して設定する必要があります。

:::

### クライアントの初期化

前のステップで構成を渡してクライアントを初期化します。

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
BKTClient.initialize(this.application, config, user)
val client = BKTClient.getInstance()
```

</TabItem>
</Tabs>

:::note

初期化処理は、アプリケーションが**フォアグラウンド状態**にある間、`pollingInterval`設定のインターバルを使用して、バックグラウンドでBucketeerから最新のエバリュエーションをすぐにポーリングし始めます。アプリケーションが**バックグラウンド状態**に変わると、`backgroundPollingInterval`設定を使用します。

:::

スプラッシュやメイン画面でフィーチャーフラグを使用したい場合、ユーザーが初めてアプリを開くと、Bucketeerサーバーからバリエーションを取得するのに十分な時間が取れない場合があります。

このケースでは、initializeメソッドから返される`Future<BKTException?>`を使用することをお勧めします。

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
// The callback will return without waiting until the fetching variation process finishes
val timeout = 1000 // Default is 5 seconds

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
    // Handle error
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

SDKにフィーチャーフラグがない場合、バリエーションメソッドはデフォルト値を返します。

:::

</TabItem>
</Tabs>

### 変更種別

Bucketeer SDKは、以下の変更種別をサポートしています。

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
fun booleanVariation(featureId: String, defaultValue: Boolean): Boolean

fun stringVariation(featureId: String, defaultValue: String): String

fun intVariation(featureId: String, defaultValue: Int): Int

fun doubleVariation(featureId: String, defaultValue: Double): Double

fun jsonVariation(featureId: String, defaultValue: JSONObject): JSONObject
```

</TabItem>
</Tabs>

### ユーザーエバリュエーションの更新

ユースケースによっては、バリエーションをリクエストする前に、SDK内のエバリュエーションが最新であることを確認する必要がある場合があります。

fetchメソッドは次のパラメータを使用し`Future<BKTExeptIon?>`を返します。

- **タイムアウト** (デフォルトは30秒)

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
val timeout = 5000
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

アプリケーションにFCM実装が既にあると仮定します。

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
override fun onMessageReceived(remoteMessage: RemoteMessage?) {
  remoteMessage?.data?.also { data ->
    val isFeatureFlagUpdated = data["bucketeer_feature_flag_updated"]
    if (isFeatureFlagUpdated) {
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

### カスタムイベントの報告

このメソッドを使用すると、アプリケーション内のユーザー操作をイベントとして保存できます。これらのイベントは、エクスペリメントコンソールUIの指標に接続できます。

さらに、ゴールイベントにダブル値を渡すことができます。これらの値は合計され、エクスペリメントコンソールUIの<br />`Value total`として表示されます。これは、ユーザーがアプリケーションでアイテムを購入した金額などを追跡するためのゴールイベントがある場合に役立ちます。

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
client.track("YOUR_GOAL_ID", 10.50)
```

</TabItem>
</Tabs>

### イベントのフラッシュ

このメソッドは、保留中のすべての分析イベントを可能な限り早くBucketeerサーバーに送信します。このプロセスは非同期ですが、何か他のことをする前に完了を待つ必要がある場合は、メソッドは`Future<BKTExeptIon?>`を返します。

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
val future = client.flush()
```

</TabItem>
</Tabs>

:::note

通常の使用では、イベントはバックグラウンドで30秒ごとに送信されるため、flushメソッド呼び出す必要はありません。

:::

### ユーザー属性の設定

この機能を使用すると、ユーザーがアプリケーションで表示できるコンテンツを堅牢かつ詳細に制御できます。これらの属性を使用して、コンソールUIのフィーチャーフラグのターゲティングタブでルールを追加できます。[詳細を見る](#)

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

### ユーザー属性の更新

このメソッドは、現在のユーザー属性をすべて更新します。これは、SDKの初期化後にアプリケーションでユーザー属性が動的に更新される場合に役立ちます。

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

この更新メソッドは、現在のデータを上書きします。

:::

### ユーザー情報の取得


このメソッドは、SDKで設定されている現在のユーザーを返します。これは、[updateUserAttributes](#updating-user-attributes)を使用して現在のユーザーIDと属性を更新する前に確認したい場合に役立ちます。

<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
val user = client.currentUser()
```

</TabItem>
</Tabs>

### エバリュエーションの詳細を取得

このメソッドは、特定のフィーチャーフラグのエバリュエーションの詳細を返します。これは、バリエーションの理由を知ったり、このデータを他の場所に送信する必要がある場合に役立ちます。
SDKにフィーチャーフラグがない場合はnullを返します。


<Tabs>
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
val evaluationDetails = client.evaluationDetails("YOUR_FEATURE_FLAG_ID")
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
<TabItem value="kt" label="Kotlin">

```kotlin showLineNumbers
// Returned value is used when you want to remove listener
val key = client.addEvaluationUpdateListener {
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
