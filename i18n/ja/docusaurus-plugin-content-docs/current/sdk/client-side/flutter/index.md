---
title: Flutterリファレンス
slug: /sdk/client-side/flutter
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

このカテゴリには、BucketeerのFlutter SDKの設定方法に関する内容が含まれています。

## はじめに

始める前に、必ず[はじめに](/getting-started)ガイドに従ってください。

### 依存関係の実装

依存関係を Pubspec ファイルに実装します。最新バージョンについては、[SDKリリースページ](https://github.com/bucketeer-io/flutter-client-sdk/releases)をご参照ください。

<Tabs>
<TabItem value="yaml" label="Pubspec">

```yaml showLineNumbers
dependencies:
  bucketeer_flutter_client_sdk: LATEST_VERSION
```

</TabItem>
</Tabs>

### クライアントのインポート

Bucketeerクライアントをアプリケーションコードにインポートします。

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
import 'package:bucketeer_flutter_client_sdk/bucketeer_flutter_client_sdk.dart';
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
<TabItem value="dart" label="Dart">

```dart showLineNumbers
final config = BKTConfigBuilder()
  .apiKey("YOUR_API_KEY")
  .apiURL("YOUR_API_URL")
  .featureTag("YOUR_FEATURE_TAG")
  .build();

final user = BKTUserBuilder
  .id("USER_ID")
  .build();
```

</TabItem>
</Tabs>

:::info カスタム構成

使用目的に応じて、**BKTConfig.Builder**で使用できるオプション構成を変更することができます。

- **pollingInterval** (最低5分。デフォルトは10分)
- **backgroundPollingInterval** (最低20分。デフォルトは1時間)
- **eventsFlushInterval** (デフォルトは30秒)
- **eventsMaxQueueSize** (デフォルトは50イベント)

:::

:::note

Bucketeer SDKはユーザーデータを保存しません。アプリケーションは、クライアントSDKを初期化する際にユーザーデータを保存して設定する必要があります。

:::

### 「クライアントの初期化」

前のステップで設定を渡して、クライアントを初期化します。

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
final client = await BKTClient.initialize(config, user);
```

</TabItem>
</Tabs>

:::note

初期化処理は、アプリケーションが**フォアグラウンド状態**にある間、`pollingInterval`設定のインターバルを使用して、バックグラウンドでBucketeerから最新のエバリュエーションをすぐにポーリングし始めます。アプリケーションが**バックグラウンド状態**に変わると、`backgroundPollingInterval`設定を使用します。

:::

スプラッシュやメイン画面でフィーチャーフラグを使用したい場合、ユーザーが初めてアプリを開くと、Bucketeerサーバーからバリエーションを取得するのに十分な時間が取れない場合があります。

このケースでは、initializeメソッドの`timeout`オプションを使用することをお勧めします。これにより、SDKが最新のユーザーバリエーションを受信するまでブロックされます。

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
// It will unlock without waiting until the fetching variation process finishes
int timeout = 1000;

final client = await BKTClient.initialize(config, user, timeout);
```

</TabItem>
</Tabs>

## サポートされている機能

### ユーザーエバリュエーション

バリエーションメソッドは、フィーチャーフラグが特定のユーザーに対して有効であるかどうかを決定します。<br />
特定のユーザーがどのバリエーションを受け取るかを確認するには、以下のようなクライアントを使用します。


<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
final showNewFeature = await client.boolVariation("YOUR_FEATURE_FLAG_ID", false);
if (showNewFeature) {
    // The Application code to show the new feature
} else {
    // The code to run if the feature is off
}
```

</TabItem>
</Tabs>

:::caution

SDKにフィーチャーフラグがない場合、デフォルト値が返されます。

:::

### バリエーションの種類

Bucketeer SDKは、以下のバリエーションの種類をサポートしています。

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
static Future<bool> boolVariation(String featureId, bool defaultValue);

static Future<String> stringVariation(String featureId, String defaultValue);

static Future<int> intVariation(String featureId, int defaultValue);

static Future<double> doubleVariation(String featureId, double defaultValue);

static Future<Map<String, dynamic>> jsonVariation(String featureId, LDValue defaultValue);
```

</TabItem>
</Tabs>

### ユーザーエバリュエーションの更新

ユースケースによっては、バリエーションをリクエストする前に、SDK内のエバリュエーションが最新であることを確認する必要がある場合があります。

fetchメソッドは次のパラメータを使用します。完了するまで必ずお待ちください。

- **タイムアウト** (デフォルトは30秒)

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
int timeout = 5000;

final result = await client.fetchEvaluations(timeout);
if (result.isSuccess) {
  final showNewFeature = await client.boolVariation("YOUR_FEATURE_FLAG_ID", false);
  if (showNewFeature) {
      // The Application code to show the new feature
  } else {
      // The code to run if the feature is off
  }
} else {
   // The code to run if the feature is off
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
<TabItem value="dart" label="Dart">

```dart showLineNumbers
// TODO
```

</TabItem>
</Tabs>

### カスタムイベントの報告

このメソッドを使用すると、アプリケーション内のユーザー操作をイベントとして保存できます。これらのイベントは、エクスペリメントコンソールUIの指標に接続できます。

さらに、目標イベントにダブル値を渡すことができます。これらの値は合計され、エクスペリメントコンソールUIの<br />`Value total`として表示されます。これは、ユーザーがアプリケーションでアイテムを購入した金額などを追跡するための目標イベントがある場合に役立ちます。

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
await client.track("YOUR_GOAL_ID", 10.50);
```

</TabItem>
</Tabs>

### イベントのフラッシュ

このメソッドは、保留中のすべての分析イベントを可能な限り早くBucketeerサーバーに送信します。このプロセスは非同期であるため、完了前に返されます。

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
client.flush();
```

</TabItem>
</Tabs>

:::note

通常の使用では、イベントはバックグラウンドで30秒ごとに送信されるため、flushメソッド呼び出す必要はありません

:::

### ユーザー属性の設定

この機能を使用すると、ユーザーがアプリケーションで表示できるコンテンツを堅牢かつ詳細に制御できます。これらの属性を使用して、コンソールUIのフィーチャーフラグのターゲティングタブでルールを追加できます。[詳細を見る](#)
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

final client = await BKTClient.initialize(config, user);
```

</TabItem>
</Tabs>

### ユーザー属性の更新

このメソッドは、現在のユーザー属性をすべて更新します。これは、SDKの初期化後にアプリケーションでユーザー属性が動的に更新される場合に役立ちます。

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

await client.updateUserAttributes(attributes);
```

</TabItem>
</Tabs>

:::caution

この更新メソッドは、現在のデータを上書きします。

:::

### ユーザー情報の取得

このメソッドは、SDKで設定されている現在のユーザーを返します。これは、[updateUserAttributes](#getting-user-information)を使用して現在のユーザーIDと属性を更新する前に確認したい場合に役立ちます。


<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
final user = await client.currentUser();
```

</TabItem>
</Tabs>

### エバリュエーションの詳細を取得

このメソッドは、特定のフィーチャーフラグのエバリュエーションの詳細を返します。これは、バリエーションの理由を知ったり、このデータを他の場所に送信する必要がある場合に役立ちます。

<Tabs>
<TabItem value="dart" label="Dart">

```dart showLineNumbers
final evaluationDetails = await client.evaluationDetails("YOUR_FEATURE_FLAG_ID");
```

:::note

このメソッドは、SDKにフィーチャーフラグがない場合はnullを返します。

:::

</TabItem>
</Tabs>
