package convert_test

import (
	"testing"

	"github.com/IamSBStakumi/gotion-md/pkg/convert"
)

func TestHeadingRenderer(t *testing.T){
	renderer := convert.HeadingRenderer{}

	tests := []struct{
		block convert.Block
		expected string
	}{
		{convert.Block{Type: "heading_1", Text: []string{"Title 1"}}, "# Title 1"},
		{convert.Block{Type: "heading_2", Text: []string{"Title 2"}}, "## Title 2"},
		{convert.Block{Type: "heading_3", Text: []string{"Title 3"}}, "### Title 3"},
	}

	for _, tt := range tests{
		if !renderer.EnableRender(tt.block){
			t.Errorf("expected EnableRender true for %v", tt.block.Type)
		}

		md := renderer.Render(tt.block)

		if md != tt.expected {
			t.Errorf("expected '%s', got '%s'", tt.expected, md)
		}
	}
}