# Product Requirements Document

## 概要

データ成熟度評価モデルに基づいた、Web ベースの自己測定ツール。

ユーザーは質問に回答することで、組織のデータ成熟度を可視化（レーダーチャート）でき、且つ結果画面を PDF ファイルとしてい出力できる。

## 利用するデータ成熟度評価モデル

以下 3 種類を利用する。

### 英国政府 Data Maturity Assessment

- Source（日本語訳）
  - [英国政府データマチュリティアセスメント日本語仮訳](https://www.ipa.go.jp/digital/data/f55m8k0000005msd-att/hm-government_data-maturity-assessment-framework_tentative-translation.xlsx)
  - 独立行政法人情報処理推進機構 (IPA) にて日本語訳された資料
- Original (原典)
  - [Data maturity assessment for government: framework](https://www.gov.uk/government/publications/data-maturity-assessment-for-government-framework)
- Data File
  - 本 Web サービスでは、Source の情報を以下の YAML ファイルに移植したデータを利用する
  - YAML: [data/uk_data_maturity.yaml](../data/uk_data_maturity.yaml)
  - JSONSchema: [data/uk_data_maturity.schema.json](../data/uk_data_maturity.schema.json)

### 豪州政府 Data Maturity Assessment

- Source（日本語訳）
  - [豪州政府データマチュリティアセスメント日本語仮訳](https://www.ipa.go.jp/digital/data/f55m8k0000005msd-att/australian-government_data-maturity-assessment-tool_tentative-translation.xlsx)
  - 独立行政法人情報処理推進機構 (IPA) にて日本語訳された資料
- Original (原典)
  - [Data Maturity Assessment Tool](https://www.finance.gov.au/government/public-data/public-data-policy/data-maturity-assessment-tool)
- Data File
  - 本 Web サービスでは、Source の情報を以下の YAML ファイルに移植したデータを利用する
  - YAML: [data/au_data_maturity.yaml](../data/au_data_maturity.yaml)
  - JSONSchema: [data/au_data_maturity.schema.json](../data/au_data_maturity.schema.json)

### 独自作成 Data Maturity Assessment

- Data File
  - YAML: [data/custom_data_maturity.yaml](../data/custom_data_maturity.yaml)
  - JSONSchema: [data/custom_data_maturity.schema.json](../data/custom_data_maturity.schema.json)

## 想定利用ユーザー

- CDO (Chief Data Officer)
- CoE (Center of Excellence) メンバー
- Data Steward
- Data Owner
- Data Architect
- Data Engineer
- Data Analyst
- Data Scientist

## 主要機能

### アセスメント機能

- ユーザーはアセスメントに利用するデータ成熟度評価モデルを、以下より選択できる
  - 英国政府 Data Maturity Assessment
  - 豪州政府 Data Maturity Assessment
  - 独自作成 Data Maturity Assessment
- ユーザーはさらに質問カテゴリを選択することで、表示される質問を絞ることが可能となる
- 質問の内容は YAML の Data File として定義済みのもので、質問より複数の選択肢（成熟度）を選ぶものとなる

### 結果表示機能

- 回答後、ユーザーは以下の結果画面を受け取ることができる
  - 質問カテゴリ毎のスコアをレーダーチャートした画像
  - 質問内容と回答結果の一覧

### 結果画面の PDF 取得機能

- 回答結果画面は、ブラウザの印刷ダイアログを利用して、PDF ファイルとして取得が可能である
