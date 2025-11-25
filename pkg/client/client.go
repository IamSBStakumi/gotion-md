package client

import (
	"context"
	"errors"
)

type Client struct{
	token string
}

func New(token string) *Client{
	return &Client{token: token}
}

// ConvertPage is the main public API.
// TODO: implement real Notion API fetch + conversion.
func (c *Client) ConvertPage(ctx context.Context, pageID string)(string, error){
	if c.token == "" {
		return "", errors.New("notion token is empty")
	}

	if pageID == "" {
		return "", errors.New("pageID is empty")
	}

	// TODO: call Notion API -> get blocks -> pass to converter
	md := "# TODO: markdown output\n"

	return md, nil
}