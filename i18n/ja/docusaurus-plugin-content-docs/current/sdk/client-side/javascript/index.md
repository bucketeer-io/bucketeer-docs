---
title: JavaScriptリファレンス
slug: /sdk/client-side/javascript
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

このカテゴリには、BucketeerのJavaScript SDKの設定方法に関する内容が含まれています。

## はじめに

始める前に、必ず[はじめに](/getting-started)ガイドに従ってください。

### 依存関係のインストール

アプリケーションに依存関係をインストールします。

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

### クライアントのインポート

Bucketeerクライアントをアプリケーションコードにインポートします。

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
import { BKTClient, getBKTClient, defineBKTConfig, defineBKTUser, initializeBKTClient } from '@bucketeer/sdk';
```

</TabItem>
</Tabs>

### クライアント構成

SDK 設定とユーザー設定を構成します。

:::info

**フィーチャータグ**設定は、フィーチャーフラグを作成する際に構成するタグです。設定されていない場合は、環境内のすべてのフィーチャーフラグを評価します。エバリュエーションプロセスを高速化し、クライアントのキャッシュサイズを削減するために、タグを使用することを**強くお勧めします**。

:::

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

:::info カスタム構成

使用目的に応じて、オプション構成を変更することができます。

- **pollingInterval** (最低5分。デフォルトは10分)
- **eventsFlushInterval** (デフォルトは30秒)
- **eventsMaxQueueSize** (デフォルトは50イベント)
- **storageKeyPrefix** (デフォルトは空)
- **userAgent** (デフォルトは`window.navigator.userAgent`)
- **fetch** (デフォルトは`window.fetch`)

:::

:::note

Bucketeer SDKはユーザーデータを保存しません。アプリケーションは、クライアントSDKを初期化する際にユーザーデータを保存して設定する必要があります。

:::

### クライアントの初期化

前のステップで構成を渡してクライアントを初期化します。

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
await initializeBKTClient(config, user);
const client = getBKTClient()
```

</TabItem>
</Tabs>

:::note

初期化処理では、`pollingInterval`の設定により、バックグラウンドでBucketeerから最新のエバリュエーションを取得します。JavaScript SDKはBackground fetchに**対応していません**。

:::

スプラッシュやメイン画面でフィーチャーフラグを使用したい場合、ユーザーが初めてアプリを開くと、Bucketeerサーバーからバリエーションを取得するのに十分な時間が取れない場合があります。

このケースでは、初期化メソッドから返される`Promise`を使用することをお勧めします。何か問題が発生した場合、Promiseは`BKTException`で拒否されます。

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
// The callback will return without waiting until the fetching variation process finishes
const timeout = 1000 // Default is 5 seconds

await initializeBKTClient(config, user, timeout);
const client = getBKTClient()
```

</TabItem>
</Tabs>

### 例外処理

多くの場合、エラーは内部的に処理されますが、一部のメソッドは何か問題が発生すると `BKTException` をスローします。 
それらのメソッドは以下の通りです：

- **`initializeBKTClient()`**
- **`BKTClient#fetchEvaluations()`**
- **`BKTClient#flush()`**

これらのメソッドは `Promise` を返しますが、`BKTException` で拒否される可能性があるため、必ずエラーを検出する必要があります。

## サポートされている機能

### ユーザーエバリュエーション

バリエーションメソッドは、フィーチャーフラグが特定のユーザーに対して有効であるかどうかを決定します。<br />
特定のユーザーがどのバリエーションを受け取るかを確認するには、以下のようなクライアントを使用します。


<Tabs>
<TabItem value="js" label="JavaScript">

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

SDKにフィーチャーフラグがない場合、バリエーションメソッドはデフォルト値を返します。

:::

### バリエーションの種類

Bucketeer SDKは、以下のバリエーションの種類をサポートしています。

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
boolVariation(featureId: string, defaultValue: boolean): Promise<boolean>;

stringVariation(featureId: string, defaultValue: string): Promise<string>;

numberVariation(featureId: string, defaultValue: number): Promise<number>;

jsonVariation(featureId: string, defaultValue: object): Promise<object>;
```

</TabItem>
</Tabs>

### ユーザーエバリュエーションの更新

ユースケースによっては、バリエーションをリクエストする前に、SDK内のエバリュエーションが最新であることを確認する必要がある場合があります。

fetchメソッドは次のパラメータを使用します。完了するまで必ずお待ちください。

- **タイムアウト** (デフォルトは30秒)

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
val timeout = 5000

await client.fetchEvaluations(timeout);
```

</TabItem>
</Tabs>

### カスタムイベントの報告

このメソッドを使用すると、アプリケーション内のユーザー操作をイベントとして保存できます。これらのイベントは、エクスペリメントコンソールUIの指標に接続できます。

さらに、ゴールイベントにダブル値を渡すことができます。これらの値は合計され、エクスペリメントコンソールUIの<br />`Value total`として表示されます。これは、ユーザーがアプリケーションでアイテムを購入した金額などを追跡するためのゴールイベントがある場合に便利です。

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
client.track("YOUR_GOAL_ID", 10.50);
```

</TabItem>
</Tabs>

### イベントのフラッシュ

このメソッドは、保留中のすべての分析イベントを可能な限り早くBucketeerサーバーに送信します。このプロセスは非同期であるため、完了前に返されます。

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
await client.flush();
```

</TabItem>
</Tabs>

:::note

通常の使用では、イベントはバックグラウンドで30秒ごとに送信されるため、Flushメソッド呼び出す必要はありません。

:::

### ユーザー属性の設定

この機能を使用すると、ユーザーがアプリケーションで表示できるコンテンツを堅牢かつ詳細に制御できます。これらの属性を使用して、コンソールUIのフィーチャーフラグのターゲティングタブでルールを追加できます。[詳細を見る](#)

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

### ユーザー属性の更新

このメソッドは、現在のユーザー属性をすべて更新します。これは、SDKの初期化後にアプリケーションでユーザー属性が動的に更新される場合に役立ちます。

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const attributes = {
  app_version: '1.0.0',
  os_version: '11.0.0',
  device_model: 'pixel-5',
  language: 'english',
  genre: 'female'
}

client.updateUserAttributes(attributes);
```

</TabItem>
</Tabs>

:::caution

この更新メソッドは、現在のデータを上書きします。

:::

### ユーザー情報の取得

このメソッドは、SDKで設定されている現在のユーザーを返します。これは、[updateUserAttributes](#getting-user-information)を使用して現在のユーザーIDと属性を更新する前に確認したい場合に役立ちます。


<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const user = client.currentUser();
```

</TabItem>
</Tabs>

### エバリュエーションの詳細を取得

このメソッドは、特定のフィーチャーフラグのエバリュエーションの詳細を返します。バリエーションの理由を知ったり、このデータを他の場所に送信する必要がある場合に役立ちます。

<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
const evaluationDetails = client.evaluationDetails("YOUR_FEATURE_FLAG_ID");
```

:::note

このメソッドは、SDKにフィーチャーフラグがない場合は null を返します。

:::

</TabItem>
</Tabs>

### エバリュエーション更新のリスニング

BKTClientは、エバリュエーションが更新された際に通知することができます。
リスナーは、自動ポーリングと手動フェッチングの両方を検出できます。


<Tabs>
<TabItem value="js" label="JavaScript">

```js showLineNumbers
// Returned value is used when you want to remove listener
val key = client.addEvaluationUpdateListener(() => {
  val showNewFeature = client.booleanVariation("YOUR_FEATURE_FLAG_ID", false)
  if (showNewFeature) {
      // The Application code to show the new feature
  } else {
      // The code to run when the feature is off
  }
})

// Remove a listener associated with the key
client.removeEvaluationUpdateListener(key)

// Remove all listeners
client.clearEvaluationUpdateListeners()
```

</TabItem>
</Tabs>
