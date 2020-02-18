# HTML Pipeline for Go

This is go version of [html-pipeline](https://github.com/jch/html-pipeline)

## Other versions

- [html-pipeline](https://github.com/jch/html-pipeline) - Ruby
- [html-pipeline.cr](https://github.com/huacnlee/html-pipeline.cr) - Crystal

## Usage

```go
package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/huacnlee/html-pipeline"
)

// ImageMaxWidthFilter a custom filter example
type ImageMaxWidthFilter struct{}

func (f ImageMaxWidthFilter) Call(doc *goquery.Document) (err error) {
	doc.Find("img").Each(func(i int, node *goquery.Selection) {
		node.SetAttr("style", `max-width: 100%`)
	})

	return
}

func main() {
	pipe := pipeline.NewPipeline([]pipeline.Filter{
		pipeline.MarkdownFilter{},
		pipeline.SanitizationFilter{},
		ImageMaxWidthFilter{},
	})

	markdown := `# Hello world

	<img onclick="javascript:alert" src="https://google.com/foo.jpg"/>`
	out, _ := pipe.Call(html)
	fmt.Println(out)
	// <h1>Hello world</h1>
	// <p><img src="https://google.com/foo.jpg" style="max-width: 100%"/></p>
}
```

https://play.golang.org/p/RoyEXqx8gui

## Built-in filters

- [SanitizationFilter](https://github.com/huacnlee/html-pipeline/blob/master/sanitization_filter.go) - Use [bluemonday](github.com/microcosm-cc/bluemonday) default UGCPolicy to sanitize html
- [MarkdownFilter](https://github.com/huacnlee/html-pipeline/blob/master/markdown_filter.go) - Use [blackfriday](https://github.com/russross/blackfriday) to covert Markdown to HTML.
- [SimpleFormatFilter](https://github.com/huacnlee/html-pipeline/blob/master/simple_format_filter.go) - Format plain text for covert `\n\n` into paragraph, like Rails [simple_format](https://api.rubyonrails.org/classes/ActionView/Helpers/TextHelper.html#method-i-simple_format).

## License

MIT License
