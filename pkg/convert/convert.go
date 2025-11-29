package convert

import "strings"

// Block -> Markdown の変換ロジックをここに定義していく

type Block struct {
	Type string
	Text []string
}

// BlockRendererは任意のブロックをMarkdownに変換するインターフェース
type BlockRenderer interface {
	Render(block Block) string
	EnableRender(block Block) bool
}

func BlocksToMarkdown(blocks []Block) string {
	renderers := []BlockRenderer{
		HeadingRenderer{},
		// ParagraphRenderer, ListRendererなどを追加
	}

	var sb strings.Builder

	for _, block := range blocks {
		for _, r := range renderers {
			if r.EnableRender(block){
				sb.WriteString(r.Render(block))
				sb.WriteString("\n\n")
				break
			}
		}
	}

	return sb.String()
}