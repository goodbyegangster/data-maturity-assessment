# Repository Structure

## ディレクトリ構造

> tree -F --dirsfirst -L 3 -a

```shell
./
├── .github/
│   └── workflows/
│       └── docker-publish.yml  # コンテナビルド用ワークフロー
├── assets/
│   └── input.css               # Tailwindcss cli のインプットファイル
├── bin/
├── cmd/
│   └── app/
│       └── main.go             # エントリーポイント
├── data/
│   ├── source/                 # IPA オリジナルデータ
│   ├── au_data_maturity.schema.json      # JSONSchema (au_data_maturity.yml)
│   ├── au_data_maturity.yaml             # 豪州政府 data maturity assessment データ
│   ├── custom_data_maturity.schema.json  # JSONSchema (custom_data_maturity.yml)
│   ├── custom_data_maturity.yaml         # 独自作成 data maturity assessment データ
│   ├── uk_data_maturity.schema.json      # JSONSchema (uk_data_maturity.yml)
│   └── uk_data_maturity.yaml             # 英国政府 data maturity assessment データ
├── docs/
│   ├── architecture.md
│   ├── development-guidelines.md
│   ├── functional-design.md
│   ├── product-requirements.md
│   └── repository-structure.md
├── internal/
│   ├── handler/                # HTTPハンドラ (chi)
│   ├── markdown/
│   ├── middleware/             # ミドルウェア
│   ├── model/                  # データ構造体
│   └── service/                # ドメインロジック
├── node_modules/
├── public/                     # ブラウザからアクセス可能な静的資産
│   ├── css/
│   │   └── style.css           # Tailwindcss cli のアウトプットファイル
│   └── js/
│       ├── chart.umd.js        # Chart.js ライブラリ
│       └── htmx.min.js         # htmx ライブラリ
├── tmp/
├── views/                      # templ コンポーネント
│   ├── components/             # 再利用可能な部品
│   ├── layout/                 # 共通レイアウト
│   ├── pages/                  # 各ページ
│   └── shared/
│       └── level_color.go      # 配色ルール（レベル別 Tailwind クラス / RGB カラー値）
├── .air.toml
├── .dockerignore
├── .gitignore
├── .node-version
├── AGENTS.md
├── Dockerfile
├── LICENSE.md
├── Makefile
├── README.md
├── go-version
├── go.mod
├── go.sum
├── package.json
└── pnpm-lock.yaml
```
