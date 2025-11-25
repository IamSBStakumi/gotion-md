# gotion-md
Go-powered CLI and library for exporting Notion pates as Markdown.

## Installation (CLI)

```
go install github.com/IamSBStakumi/gotion-md/cmd/gotion-md@latest
```

## Usage (CLI)

```
export NOTION_TOKEN="secret_xxx"

gotion-md -o out.md
```

## Usage (Library)

```go
client := gotion.New(os.Getenv("NOTION_TOKEN"))
md, _ := gotion.ConvertPage(context.Background(), "your_page_id")
fmt.Println(md)
```
