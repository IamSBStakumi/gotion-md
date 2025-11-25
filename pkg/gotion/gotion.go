package gotion

import (
	"context"
	"errors"
)

type Gotion struct{
	token string
}

func New(token string) *Gotion{
	return &Gotion{token: token}
}

// ConvertPage is the main public API.
// TODO: implement real Notion API fetch + conversion.
func (g *Gotion) ConvertPage(ctx context.Context, pageID string)(string, error){
	if g.token == "" {
		return "", errors.New("notion token is empty")
	}

	if pageID == "" {
		return "", errors.New("pageID is empty")
	}

	// TODO: call Notion API -> get blocks -> pass to converter
	md := "# TODO: markdown output\n"

	return md, nil
}