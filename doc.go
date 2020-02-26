/*
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
	pipeline "github.com/huacnlee/html-pipeline"
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
		pipeline.MentionFilter{
			Prefix: "#",
			Format: func(name string) string {
				return fmt.Sprintf(`<a href="https://github.com/topic/%s">#%s</a>`, name, name)
			},
		},
		pipeline.MentionFilter{
			Prefix: "@",
			Format: func(name string) string {
				return fmt.Sprintf(`<a href="https://github.com/%s">@%s</a>`, name, name)
			},
		},
	})

	markdown := `# Hello world

![](javascript:alert) [Click me](javascript:alert)

This is #html-pipeline example, @huacnlee created.`
	out, _ := pipe.Call(markdown)
	fmt.Println(out)
}
```

https://play.golang.org/p/zB0T7KczdB4

## Built-in filters

- [SanitizationFilter](https://github.com/huacnlee/html-pipeline/blob/master/sanitization_filter.go) - Use [bluemonday](github.com/microcosm-cc/bluemonday) default UGCPolicy to sanitize html
- [MarkdownFilter](https://github.com/huacnlee/html-pipeline/blob/master/markdown_filter.go) - Use [blackfriday](https://github.com/russross/blackfriday) to covert Markdown to HTML.
- [MentionFilter](https://github.com/huacnlee/html-pipeline/blob/master/mention_filter.go) - Match Mention or HashTag like Twitter.
- [HTMLEscapeFilter](https://github.com/huacnlee/html-pipeline/blob/master/html_escape_filter.go) - HTML Escape for plain text.
- [SimpleFormatFilter](https://github.com/huacnlee/html-pipeline/blob/master/simple_format_filter.go) - Format plain text for covert `\n\n` into paragraph, like Rails [simple_format](https://api.rubyonrails.org/classes/ActionView/Helpers/TextHelper.html#method-i-simple_format).
*/

package pipeline
