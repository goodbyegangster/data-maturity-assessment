# Architecture

## 概要

`htmx` を利用した Server Side Rendering にて、Web サービスを作成する。

ただし、一部機能（レーダーチャート描画、PDF 出力）については、ブラウザ側に処理を委ねる方針とする。

## Technical Specifications

### Backend

#### Language

`Go` を採用する。

- 利用バージョンは [go.mod](../go.mod) を参照

#### Routing

[go-chi/chi](https://github.com/go-chi/chi) を採用する。

- 利用バージョンは [go.mod](../go.mod) を参照
- `main.go` に全てのルートを書かず、`internal/handler/` 配下に関心事単位で分離する
- ミドルウェアを活用し、ハンドラ内での共通処理（htmx 判定など）を重複させない

#### Templating Engine

[a-h/templ](https://github.com/a-h/templ) を採用する。

- 利用バージョンは [go.mod](../go.mod) を参照
- コンポーネントは `views/` ディレクトリ内で機能単位に整理する
- `templ` は受け取ったデータの表示のみに徹する
  - ビジネスロジック、複雑なデータ加工は Go 側の Service 層で行う
- 表示する成熟度モデルの YAML データは、markdown 記法で記載されている
  - Web に表示する際には [yuin/goldmark](https://github.com/yuin/goldmark) を利用した変換を実施する

#### Database

利用しない。オンメモリで完結させる。

### Frontend

#### Markup

`HTML5` を採用する。

#### Styling

[Tailwind CSS](https://tailwindcss.com/) を採用する。

- CDN 経由で使用するのではなく、`public/css/style.css` を利用する
- 5 回以上繰り返されるスタイルは、CSS の `@apply` を使わず、`templ` の小さなコンポーネントとして共通化する

#### Interactivity

[htmx](https://htmx.org/) を採用する。

- CDN 経由で使用するのではなく、`public/js/htmx.min.js` を利用する
- `Locality of Behavior (行動の局所性)`に従う
  - 要素の挙動（htmx）と見た目（Tailwind CSS）は、可能な限りその HTML 要素内に記述する
- `HX-Request` ヘッダーをチェックし、フルレンダリングと部分（Partial）レンダリングを適切に切り分ける
- 予期せぬ書き換えを防ぐため、`hx-target` は可能な限り明示的に指定する
- ページの状態が変わる操作では、`hx-push-url="true"` を検討し、ブラウザの「戻る」ボタンに対応させる (URL 同期)
- エラー時は `HX-Trigger` ヘッダー等を利用して、フロントエンドに通知する

#### Radar Chart Visualization

[Chart.js](https://www.chartjs.org/) を採用する。

- レーダーチャートの描画に使用
- CDN 経由で使用するのではなく、`public/js/chart.umd.js` を利用する

#### PDF Print Dialog

ブラウザの印刷機能 `(window.print())` を使用する。

- Tailwind CSS の `print:` バリアントを使用して、印刷時に不要な要素（ナビゲーション・ボタン等）を非表示にする
