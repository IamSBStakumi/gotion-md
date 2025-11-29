package convert

import (
	"fmt"
	"strings"
)

type HeadingRenderer struct{}

func(r HeadingRenderer) EnableRender(block Block) bool {
	return block.Type == "heading_1" || block.Type == "heading_2" || block.Type == "heading_3"
}

func(r HeadingRenderer) Render(block Block) string{
	level := 0
	
	switch block.Type{
	case "heading_1":
		level = 1
	case "heading_2":
		level = 2
	case "heading_3":
		level = 3
	default:
		// EnableRenderで除外されるので基本的にここには来ない
		return ""
	}

	prefix := strings.Repeat("#", level)
	text := strings.Join(block.Text, "")

	return fmt.Sprintf("%s %s", prefix, text)
}