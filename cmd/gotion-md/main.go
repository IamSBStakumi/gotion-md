package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/IamSBStakumi/gotion-md/pkg/client"
)

func main(){
	out := flag.String("o", "out.md", "output markdown file")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: gotion-md <page-id> [-o out.md]")
		os.Exit(1)
	}

	pageID := flag.Arg(0)
	token := os.Getenv("NOTION_TOKEN")
	if token == "" {
		fmt.Println("ERROR: NOTION_TOKEN is not set")
		os.Exit(1)
	}

	c:=client.New(token)
	md, err := c.ConvertPage(context.Background(), pageID)
	if err != nil {
		fmt.Println("convert error:", err)
		os.Exit(1)
	}

	if err := os.WriteFile(*out, []byte(md), 0644); err != nil {
		fmt.Println("write error:", err)
		os.Exit((1))
	}

	fmt.Println("wrote: ", *out)
}