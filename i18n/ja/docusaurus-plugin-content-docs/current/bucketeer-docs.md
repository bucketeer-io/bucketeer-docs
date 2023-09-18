---
title: Bucketeer Docs
sidebar_position: 1
slug: /
description: Describes what the Bucketeer is and its solution. In addition, the page also provides an overview of the main sections covered in the documentation.
tags: ['home', 'guide', 'presentation', 'overview', 'contact']
---

import Button from '@site/src/components/button/Button';
import ButtonShelf from '@site/src/components/button-shelf/ButtonShelf';

# Bucketeerドキュメントへようこそ

Bucketeerドキュメントへようこそ。ここは、Bucketeerプラットフォーム、インテグレーション、SDKに関するあらゆる情報を得ることができるサイトです。Bucketeerソリューションを効果的に活用するために必要なすべての情報がここにあります。以下の便利なナビゲーションボタンを使用することで、目的のセクションにすばやくアクセスすることができます。さらにこのページでは、Bucketeerの主な特徴や機能を詳しくご紹介します。

<ButtonShelf>
  <Button
    redirect="../getting-started"
    title="はじめに"
    info="Bucketeerソリューションのクイックスタートガイドを提供します。このガイドでは、フラグの作成とシステムへの統合について説明します。"
  />
  <Button
    redirect="../feature-flags"
    title="フィーチャーフラグ"
    info="フィーチャーフラグの作成と管理、およびそれらを使用してテストを実行する方法について説明します。"
  />
  <Button
    redirect="../sdk"
    title="SDK"
    info="Bucketeer SDKをシステムに統合する方法を、サーバーアプリケーションとクライアントアプリケーションを含めて詳細に説明します。"
  />
  <Button
    redirect="../contribution-guide/contributing"
    title="貢献ガイド"
    info="Bucketeerシステムへの貢献方法を提示し、ドキュメントを作成する人向けのスタイルガイドを含みます。"
  />
</ButtonShelf>

:::info

管理コンソールの更新に伴い、現在ドキュメンテーションサイトも更新しています。

毎週更新していきますので、どうぞご期待ください！

:::

## Bucketeerとは

Bucketeerは、強力なツールを提供することでソフトウェア開発を最適化する機能管理プラットフォームです。コントロールされた機能リリースのためのフィーチャーフラグを提供し、簡単な実験、A/Bテスト、およびターゲットユーザーセグメンテーションを可能にします。このプラットフォームは、ベータテストプログラムをサポートし、動的な構成変更を容易にし、ダークローンチと特定のユーザーグループへのリリースのロールアウトを可能にします。また、Bucketeerを使用することで、チームは機能ライフサイクルを効率的に管理し、貴重なユーザーフィードバックを収集し、データドリブンの意思決定を行って、全体的なユーザーエクスペリエンスを向上させることができます。

## Bucketeerソリューションでできること

### A/Bテスト

Bucketeerは、表面的な変更を超えて、大幅な機能のテストを可能にする強力なA/Bテストソリューションを提供します。この機能管理プラットフォームを使用することで、組織は目標を設定し、さまざまな機能の影響を効果的に測定することができます。

Bucketeerで[A/Bテスト](./feature-flags/testing-with-flags/experiments)を作成する方法を確認してください。

### ユーザーターゲティング

高度にカスタマイズされたメタデータのポテンシャルを引き出して、正確で効果的なユーザーターゲティングを実現します。地域、年齢、メールなどの属性を活用することで、特定の基準に合わせたセグメントやグループを作成できます。このレベルの詳細度では、誰が何を表示するかを完全に制御できるため、パーソナライズされた体験を提供することができます。

Bucketeer内でのターゲティングの動作について調査してください。[こちら](./feature-flags/creating-feature-flags/targeting)をご覧ください。

### 前提条件

フラグとユーザーターゲティングを向上させるために前提条件を使用し、フラグ間の依存関係を確立します。フラグの評価すべきか、エンドユーザーに提供すべき変化を決定する条件や依存関係を定義します。

Bucketeerの前提条件機能について詳しく学ぶには、[こちら](./feature-flags/creating-feature-flags/targeting#prerequisites)をご覧ください。

### キルスイッチ

機能リリースを元に戻すためにコードをリバーとする必要はありません。Bucketeerを使用すると、誰でもいつでもすぐに機能を無効化できます。フィーチャーフラグをキルスイッチとして使用することは、機能リリースに関連するリスクを軽減するための一般的な方法です。この機能により、製品チームとマーケティングチームは、エンジニアリングサポートに大きく依存することなく、機能テストとリリースに参加できます。本番環境でテストを実施しているか、カナリアリリースを実装しているか、または機能の廃止を準備しているかに関係なく、スイッチを押すだけという簡単な操作で行うことができるので安心です。


Bucketeerにおけるキルスイッチの動作について確認してください。[こちら](./feature-flags/creating-feature-flags/auto-operation#how-auto-operation-works)をご覧ください。

### ベータテスト

Bucketeerの機能管理プラットフォームは、大規模なベータテストプログラムを効率的に管理するために広く使用されています。ターゲティングルールに基づいて特定のグループを作成すると、必要に応じてグループ全体またはセグメントをテストに含めることができます。ベータグループからのフィードバックは非常に貴重であり、新しい機能を検証し、すべてのユーザーベースにロールアウトする前にバグを特定して修正することができます。


Bucketeerを使用してターゲティングする際に、[ユーザーグループ](./feature-flags/creating-feature-flags/targeting#targeting)の使い方を探索してください。

### ダイナミック設定

リリース済みの製品の設定を変更することは、気が重くなる作業になる可能性があります。しかし、フィーチャーフラグを使用すると、ユーザーフレンドリーなインターフェイスを介してアプリケーションを簡単に構成し、値を簡単に変更できます。この機能は、エンジニアだけに限定されるものではありません。ビジネスチームにも拡張され、要件に応じて変更を加えることができます。


### ダークローンチ

Bucketeerを使用すると、フィーチャーフラグをOFFにすることで、何も影響を与えずに本番環境に新機能を導入できます。時期が来たら、これらの機能を簡単に有効化して、スムーズな移行を確保できます。


### ロールアウトリリース

専用のベータテスターグループなどのユーザーサブセットに新機能を簡単にロールアウトし、実際の使用シナリオから貴重なフィードバックとバグレポートを収集します。これにより、新機能をより広いユーザーにリリースする前に、ユーザーエクスペリエンスに基づいて新機能を改善および最適化できます。


"Bucketeerで[ロールアウトリリース](./feature-flags/creating-feature-flags/targeting#rollout-percentage)を作成する方法を学んでください。"

### トランクベース開発

フィーチャーフラグをOFFにしてフィーチャーブランチをより頻繁にマージすることで、チームは進行中のブランチの数を減らすことができます。このアプローチは、レビュアーの「ビッグバンリリース」に対するストレスと不安を軽減します。

### 廃止された機能

時間の経過とともに、古い機能は新しい機能と競合したり、時代遅れになったりすることがあります。Bucketeerのプラットフォームは、どの機能がまだ使用されており、誰が使用しているかを可視化することで、チームはコードベースを効果的に管理し、何を保持すべきかを判断することができます。


### 反応モニタリング

Bucketeerの機能管理プラットフォームは、チームがリアルタイムで問題を解決するのに役立ちます。アプリケーション内にフィーチャーフラグを実装すると、チームはBucketeerの監視ログを使用して問題を引き起こした変更を特定し、迅速な解決と継続的な改善を可能にします。


## Bucketeer チームへのお問い合わせ

お探しの回答が見つからない場合は、お気軽に[お問い合わせ](https://app.slack.com/client/T08PSQ7BQ/C043026BME1)ください。  
