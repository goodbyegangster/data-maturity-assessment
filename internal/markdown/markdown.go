package markdown

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

// converter はアプリケーション全体で共有する Markdown → HTML 変換器。
// 入力は自己管理 YAML ファイルのみを想定しているため、unsafe レンダリングを許可する。
var converter = goldmark.New(
	goldmark.WithRendererOptions(html.WithUnsafe()),
)

// ToHTML は Markdown テキストを HTML 文字列に変換して返す。
// 変換に失敗した場合は入力をそのまま返す。
func ToHTML(src string) string {
	var buf bytes.Buffer
	if err := converter.Convert([]byte(src), &buf); err != nil {
		return src
	}
	return buf.String()
}
