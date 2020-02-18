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
		pipeline.SanitizationFilter{},
		ImageMaxWidthFilter{},
	})

	html := `<img onclick="javascript:alert" src="https://google.com/foo.jpg"/>`
	out, _ := pipe.Call(html)
	fmt.Println(out)
	// <img src="https://google.com/foo.jpg" style="max-width: 100%"/>
}
```

https://play.golang.org/p/teBIIhyFNug

## Built-in filters

- SanitizationFilter - Use bluemonday default UGCPolicy to sanitize html
- SimpleFormatFilter - Format plain text for covert `\n\n` into paragraph, like Rails [simple_format](https://api.rubyonrails.org/classes/ActionView/Helpers/TextHelper.html#method-i-simple_format).

## License

MIT License
