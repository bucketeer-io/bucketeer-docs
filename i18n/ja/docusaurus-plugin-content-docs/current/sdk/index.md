---
id: overview
title: SDKの概要
slug: /sdk
---

このカテゴリーでは、**クライアント**と**サーバーサイド**のSDKの違いを説明するガイドを見つけることができます。Bucketeerサーバーに接続するには、プログラミング言語とAPIキーに対応したSDKが必要です。

:::note

APIキーは、Bucketeerダッシュボードの**APIキー**メニューから生成されます。

:::

## Bucketeer SDKの選択

Bucketeer SDKをコードと統合する場合、次の2つの要素を考慮することが重要です。

1. 最初に、サーバーサイドSDKが必要かクライアントサイドSDKが必要かを判断するために、特定の要件を評価する必要があります。通常、アプリケーションまたはサービスごとに1つのSDKのみが必要になります。ただし、製品が複数の言語で記述されたアプリケーションまたはサービスで構成されている場合は、複数のSDKを使用できます。
2. その後、システムで通常使用されるプログラミング言語を選択する必要があります。


:::tip

アプリケーションでBucketeer SDKを使用する前に、[クイックスタート](/getting-started/quickstart)ガイドをご覧いただくことを強くお勧めします。

:::

### クライアントサイドSDK

クライアントサイドSDKは、シングルユーザーのデスクトップ、モバイル、および埋め込みアプリケーション向けに設計されています。パーソナルコンピュータやモバイルデバイスなどの、潜在的にセキュリティの低い環境での使用を目的としています。これらのSDKは通常、エンドユーザーデバイスで実行されるため、モバイルアプリを解凍してSDKバイトコードを調べたり、ブラウザの開発者ツールを使用して内部サイトデータを確認するユーザーによって侵害される可能性があります。したがって、クライアントサイドSDKまたはモバイルアプリケーションでサーバーサイドSDKキーを使用しないことは不可欠です。


クライアントサイドのサポートされているSDK：


<div className="row" style={{maxWidth: '500px'}}>

  <div className="col--3 text--center">
    <a href="/sdk/client-side/android" className="brand-link">
      <i className="android-icon brand-icon"></i>
      <span>Android</span>
    </a>
  </div>

  <div className="col--3 text--center">
    <a href="/sdk/client-side/ios" className="brand-link">
      <i className="ios-icon brand-icon"></i>
      <span>iOS</span>
    </a>
  </div>

  <div className="col--3 text--center">
    <a href="/sdk/client-side/flutter" className="brand-link">
      <i className="flutter-icon brand-icon"></i>
      <span>Flutter</span>
    </a>
  </div>

  <div className="col--3 text--center">
    <a href="/sdk/client-side/javascript" className="brand-link">
      <i className="javascript-icon brand-icon"></i>
      <span>JavaScript</span>
    </a>
  </div>

</div>

### サーバーサイドSDK

サーバーサイドSDKは、マルチユーザーシステム向けに設計されており、企業ネットワークやWebサーバなどの信頼できる環境での使用を目的としています。インフラストラクチャまたは信頼できるクラウドベースのインフラストラクチャ上で実行されるサーバーアーキテクチャのアプリケーション内で動作します。これらの場所は、エンドユーザーから直接アクセスすることはできません。サーバーベースのアプリケーションのアクセスが制限されているため、サーバーサイドSDKは、機密情報をフィルタリングする必要なしに、フラグデータとルールセットを安全に受信できます。この設定により、サーバーサイドSDKは、ルールとセグメントを含むフラグデータを、機密情報をフィルタリングすることなく安全に受信することが可能です。

サーバーサイドのサポートされているSDK：

<div className="row" style={{maxWidth: '500px'}}>

  <div className="col--3 text--center">
    <a href="/sdk/server-side/go" className="brand-link">
      <i className="golang-icon brand-icon"></i>
      <span>Go</span>
    </a>
  </div>

  <div className="col--3 text--center">
    <a href="/sdk/server-side/node-js" className="brand-link">
      <i className="nodejs-icon brand-icon"></i>
      <span>Node.js</span>
    </a>
  </div>

</div>
