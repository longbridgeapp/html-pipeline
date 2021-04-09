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

	/*
		<h1>Hello world</h1>

		<p><img alt="" style="max-width: 100%"/> Click me</p>

		<p>This is <a href="https://github.com/topic/html-pipeline">#html-pipeline</a> example, <a href="https://github.com/huacnlee">@huacnlee</a> created.</p>
	*/
}
```

https://play.golang.org/p/zB0T7KczdB4

## Use for Plain Text case

Sometimes, you may want use html-pipeline to manage the Plain Text process.

For example:

- Match mentions, and then send notifications.
- Convert Mention / HashTag or other text into other format.

But in HTML mode, it will escape some chars (`"`, `'`, `&`) ... We don't wants that.

So, there have `NewPlainPipeline` method for you to create a plain mode pipeline without any escape.

> NOTE: For secruity, this pipeline will remove all HTML tags `<.+?>`

```go
package main

import (
	"fmt"
	"github.com/huacnlee/html-pipeline"
)

func main() {
	pipe := pipeline.NewPlainPipeline([]pipeline.Filter{
		pipeline.MentionFilter{
			Prefix: "#",
			Format: func(name string) string {
				return fmt.Sprintf(`[hashtag name="%s"]%s[/hashtag]`, name, name)
			},
		},
		pipeline.MentionFilter{
			Prefix: "@",
			Format: func(name string) string {
				return fmt.Sprintf(`[mention name="%s"]@%s[/mention]`, name, name)
			},
		},
	})

	text := `"Hello" & 'world' this <script>danger</script> is #html-pipeline created by @huacnlee.`
	out, _ := pipe.Call(text)
	fmt.Println(out)
	// "Hello" & 'world' this danger is [hashtag name="html-pipeline"]html-pipeline[/hashtag] created by [mention name="huacnlee"]@huacnlee[/mention].
}
```

https://play.golang.org/p/vxKZU9jJi3u

## Built-in filters

- [SanitizationFilter](https://github.com/huacnlee/html-pipeline/blob/master/sanitization_filter.go) - Use [bluemonday](github.com/microcosm-cc/bluemonday) default UGCPolicy to sanitize html
- [MarkdownFilter](https://github.com/huacnlee/html-pipeline/blob/master/markdown_filter.go) - Use [blackfriday](https://github.com/russross/blackfriday) to covert Markdown to HTML.
- [MentionFilter](https://github.com/huacnlee/html-pipeline/blob/master/mention_filter.go) - Match Mention or HashTag like Twitter.
- [HTMLEscapeFilter](https://github.com/huacnlee/html-pipeline/blob/master/html_escape_filter.go) - HTML Escape for plain text.
- [SimpleFormatFilter](https://github.com/huacnlee/html-pipeline/blob/master/simple_format_filter.go) - Format plain text for covert `\n\n` into paragraph, like Rails [simple_format](https://api.rubyonrails.org/classes/ActionView/Helpers/TextHelper.html#method-i-simple_format).
- [AutoCorrectFilter](https://github.com/huacnlee/html-pipeline/blob/master/auto_correct_filter.go) - Use [go-auto-correct](https://github.com/huacnlee/go-auto-correct) to automatically add spaces between Chinese and English words.
- [ImageProxyFilter](https://github.com/huacnlee/html-pipeline/blob/master/image_proxy_filter.go) - A filter can match all `img` to replace src as proxy url with [imageproxy](https://github.com/willnorris/imageproxy).

## License

MIT License
