---
id: contributing
title: 貢献
slug: /contribution-guide/contributing
---

# 貢献ガイド

Bucketeerをより良いサービスにするために、是非ご協力をお願いします！どなたでも利用し、改善し、お楽しみいただけます！<br />
ここでお探しの回答が見つからない場合は、お気軽に[お問い合わせ](https://app.slack.com/client/T08PSQ7BQ/C043026BME1)ください。


## 貢献者ライセンス契約

このプロジェクトへの貢献には、[bucketeer-io/bucketeer/master/CLA.md](https://github.com/bucketeer-io/bucketeer/blob/master/CLA.md)に記述されている貢献者ライセンス契約(CLA)が必要です。<br />
あなた(またはその雇用主)があなたの貢献に対する著作権を保持するものとします。これは、プロジェクトの一部としてあなたの貢献を使用し、再配布する許可を私たちに与えるものです。

通常、CLAへの署名は一度だけ必要ですので、既に署名している場合は、再度署名する必要はおそらくありません。

CLAへの署名がまだの場合は、[bucketeer-io/bucketeer](https://github.com/bucketeer-io/bucketeer)リポジトリに最初のプルリクエストを送る際に、[bucketeer-bot](https://github.com/bucketeer-bot)がCLAへの署名を案内します。

## 課題の作成

課題を見つけた場合は、[bucketeer-io/bucketeer](https://github.com/bucketeer-io/bucketeer/issues) リポジトリで課題を作成してください。


## プルリクエストの作成中

リポジトリをフォークしてください。

最初のプルリクエストに最適な課題を見つけるには、[ヘルプが必要な課題](https://github.com/bucketeer-io/bucketeer/issues?q=is%3Aissue+is%3Aopen+label%3A"help+wanted")や[最初に最適な課題](https://github.com/bucketeer-io/bucketeer/issues?q=is%3Aissue+is%3Aopen+label%3A"good+first+issue")を参考にしてください。


:::info

レビュー依頼後に PRブランチに強制的にプッシュするのはやめてください。何が変更されたかわからないため、PR全体を再度レビューすることを余儀なくされます。

:::

## コミットメッセージの形式

私たちは[Conventional Commits 1.0.0](https://www.conventionalcommits.org/en/v1.0.0)のメッセージ規約に従っています。<br />
この形式はコミット履歴を閲覧しやすくし、変更履歴の生成や[セマンティック・バージョニング](https://semver.org)に準拠することを可能にします。


:::tip

コミットメッセージは私たちのリリースの変更履歴に使用されます。ユーザーの視点からわかりやすいメッセージを書くようお願いします。

:::

### タイプ

以下のいずれかに該当する必要があります：

- **build:** ビルドシステムまたは外部ライブラリの依存関係に影響を与える変更
- **chore:** srcファイルまたはテストファイルを変更しないその他の変更
- **ci:** CI設定ファイルおよびスクリプトの変更
- **docs:** ドキュメントのみの変更
- **feat:** 新しい機能または既存の機能の更新
- **fix:** バグ修正
- **perf:** パフォーマンス向上のためのコード変更
- **refactor:** バグ修正でも機能追加でもないコード変更
- **revert:** 直前のコミットの取り消し
- **test:** 不足しているテストの追加または既存のテストの修正


:::info

破壊的変更コミットの場合、APIの破壊的変更(セマンティック・バージョニングのメジャーバージョンと関連)を導入するタイプの後に `!` を追加する必要があります。<br />
例：`feat!: new API interface 2.0`

:::

### サブジェクト

サブジェクトには変更の説明が含まれます：

- "changed"や "changes"ではなく命令形、現在形"change"を使用する
- 最初の文字を大文字にしない
- 最後にドット(.)を付けない

### 例

```
feat: add webhook feature
^--^  ^-----------------^
|     |
|     +-> Subject in present tense. Not capitalized.
|
+-------> Type: build|chore|ci|docs|feat|fix|perf|refactor|revert|test
```

## コードレビュー

すべての提出、プロジェクトメンバーによる提出も含め、レビューが必要となります。この目的のためにGitHubのプルリクエストを使用しています。プルリクエストの使い方については、[GitHub Help](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/about-pull-requests) をご参照ください。
