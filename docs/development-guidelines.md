# Development Guidelines

## コーディング規約

### Basic

- コード内のコメント、コミットメッセージ、ドキュメントには日本語を使用する
- ユーザー向けのメッセージは日本語を使用する
- 開発者向け・ログ用のメッセージは英語を使用する

### Backend

- google の [Go Style Guide](https://google.github.io/styleguide/go/guide) に準拠する
- google の [Go Style Decisions](https://google.github.io/styleguide/go/decisions) に準拠する

### Frontend

- [Google HTML/CSS Style Guide](https://google.github.io/styleguide/htmlcssguide.html) に準拠する
- MDN の [HTML: A good basis for accessibility](https://developer.mozilla.org/en-US/docs/Learn_web_development/Core/Accessibility/HTML) に準拠する
- MDN の [CSS and JavaScript accessibility best practices](https://developer.mozilla.org/en-US/docs/Learn_web_development/Core/Accessibility/CSS_and_JavaScript) に準拠する

## 更新禁止のディレクトリおよびファイル

```shell
./
├── data/*
├── docs/*
├── AGENTS.md
├── LICENSE.md
└── README.md
```
