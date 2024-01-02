---
title: エクスペリメントの結果
# sidebar_position: 
slug: /feature-flags/testing-with-flags/experiment-results
description: Describes the experiments tab on the feature flag and how to link feature flags to experiments.
tags: ['experiments', 'feature-flag', 'test']
---

import CenteredImg from '@site/src/components/centered-img/CenteredImg';

エクスペリメントを作成すると、自動的にダッシュボードの**エクスペリメント**タブに表示されます。エクスペリメントには、以下のいずれかの状態が示されます：

- **待機中**: 定義された開始日に達していない状態を指します。エクスペリメント結果は表示されません。
- **実行中**: エクスペリメントは、リストに記載されたゴールに従ってユーザーの行動に関するデータを収集しています。エクスペリメント結果ページには、既にデータが記載されています。 
- **完了**: エクスペリメントが終了日に達し、正常に完了している状態を指します。最終結果は結果ページに記載されています

エクスペリメントの状態に関わらず、**結果**tボタンをクリックすればいつでもエクスペリメントの結果にアクセスすることができます。以下の画像は、エクスペリメントのリストの一例を示しています。

<CenteredImg
  imgURL="img/feature-flags/using-experiments/experiments-list.png"
  alt="List of experiments"
  wSize="550px"
/>

フラグの詳細にアクセスすると、**エクスペリメント**タブを選択してエクスペリメント結果を確認することもできます。

## エクスペリメント結果の分析

エクスペリメント結果ページに移動すると、すぐにエクスペリメントの現在の状態とエバリュエーション期間に関する情報が上部に表示されます。例えば、次の例は**停止中*のエクスペリメントを示しています。

<CenteredImg
  imgURL="img/feature-flags/using-experiments/stopped-experiment.png"
  alt="Stopped experiment."
/>

エクスペリメントの状態に応じて、以下のコンポーネントが表示されます：

- **ゴールセレクター**: この機能を使用すると、エクスペリメントに関連付けられているすべてのゴールから選択することができます。

- **バリエーションデータテーブル**: 選択したゴール内のすべてのバリエーションのデータを包括的に表示するテーブル（表）です。このテーブルには、各バリエーションのパフォーマンスに関する関連する指標と洞察が示されており、それらの効果を比較および評価することができます。

- **バリエーションデータ履歴チャート**: 選択したゴールとその関連するバリエーションの履歴データを可視化します。このグラフは、時間の経過に伴うパフォーマンスの傾向とパターンを示し、データの変化を観察することができます。このグラフは、テーブルで利用可能なすべての列に関連するデータを表示できます。


以下の画像は、エクスペリメントの例のデータセットを示しており、**ゴール-1**に焦点を当てています。グラフは、4つの既存のバリエーションすべてのコンバージョン率データを示しており、それらのパフォーマンスを評価および比較することができます。

<CenteredImg
  imgURL="img/feature-flags/using-experiments/experiments-1.png"
  alt="First portion of experiment panel"
  borderWidth="1px"
/>

### テーブルデータ

エクスペリメント結果のバリエーションデータテーブルは、分析に役立つ貴重な情報を提供します。それには以下の列が含まれています：

- **エバリュエーションユーザー**: この列は、SDKリクエスト後にサーバーからバリエーションを受け取ったユニークユーザー数を表します。これは、特定のフィーチャーフラグバリエーションに割り当てられたユーザーの実際の数を示しています。

- **ゴール合計**: クライアントによって発生したゴールイベントの総数を表示します。同じユーザーによる複数回のトリガーを含む、ゴールイベントのすべての発生をカウントします。

- **ゴールユーザー**: ゴールイベントを達成したユニークユーザー数を示します。ゴール合計とは異なり、このカウントは同じユーザーがゴールイベントを複数回トリガーしても増加しません。ゴールを達成したユーザーの個別のカウントを表します。

- **コンバージョン率**: コンバージョン率の列は、ゴールイベントを発生させたユニークユーザー数を、バリエーションが返されたユニークユーザー数で割って計算されます。割り当てられたバリエーションに基づいて、ゴールを正常に完了したユーザーの割合に関する洞察を提供します。

- **合計値**: 合計値の列は、ゴールイベントに割り当てられた値の合計を表します。この列は、変動ごとに異なる値を持つ場合があります。例えば、各バリエーションに基づいてユーザーが費やした金額を測定することができます。ゴールがユーザーが特定のタスクを実行したかどうかを確認することである場合 (例: ボタンをクリックする)、この列には NULL 値のみが含まれる場合があります。

- **値/ユーザー**: 値/ユーザーの列は、ユーザーごとのゴールイベントに割り当てられた値の平均を計算します。これは、ゴールイベントに割り当てられた値の合計を、ゴールイベントを発火したユニークユーザーの数で割った値です。


### ベイズ推定に基づく最善なバリエーション

画面の下部には、Bucketeerシステムが、エクスペリメントで最もパフォーマンスの高いバリエーションに関する提案を提供します。最善なバリエーションを決定するために、Bucketeerは[ベイズ推定](https://en.wikipedia.org/wiki/Bayesian_inference)を活用しています。このアプローチにより、データサイエンスの深い理解を必要とせずに、上位のパフォーマンスを発揮するバリエーションに関する予測を行うことができます。

ベイズ推定を使用することで、Bucketeerは、新しいデータが取得されるたびに、最善なバリエーションを選択する確率を更新できます。Bucketeerは、各バリエーションのパフォーマンスをベースラインと比較します。ベースラインは、通常、エクスペリメントの作成時に設定されたコントロールグループまたは基準点です。

エクスペリメント結果には、既存のテーブルに追加情報が記載されているため、最もパフォーマンスの高いバリエーションの選択に役立ちます。:


- **改善率**: この指標は、各バリエーションがベースラインと比較して達成した改善を定量化します。これは、バリエーションで観測された値の範囲とベースラインで観測された値の範囲を比較することで計算されます。改善率が高いほど、ベースラインと比較してより良好なパフォーマンスを示していることを示します。

- **ベースラインを上回る確率**: この推定尤度は、バリエーションがベースラインのパフォーマンスを上回る確率を表します。これは、バリエーションがベースラインを上回る可能性を評価するのに役立ちます。また、各バリエーションの相対的な有効性に関する洞察を提供します。

- **最善である確率**: この指標は、バリエーションが最もパフォーマンスの高いオプションである確立を示します。これは、バリエーションが他のすべてのバリエーションを上回り、最も成功したバリエーションである可能性を表します。


以下の画像は、前回提供された一連の結果から続くベイズ推定を使用して取得された結果の例を示しています。

<CenteredImg
  imgURL="img/feature-flags/using-experiments/experiments-2.png"
  alt="Second portion of experiment panel"
  borderWidth="1px"
/>

:::tip 最善なバリエーションを選ぶ方法とは？
Bucketeerチームは、ベースラインを上回る確率と最善である確率の両方が最低95％の信頼水準を満たすバリエーションを選択することを推奨しています。
:::

Bucketeer の推奨値に基づいて、現在のテストは、最善のバリエーションを選択する前に、より多くのデータの取得を継続する必要があります。推奨される信頼値に達するバリエーションがない場合、テストは実行され続けます。