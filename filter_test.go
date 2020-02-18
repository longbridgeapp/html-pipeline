package pipeline

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
)

type TestFilter struct{}

func (f TestFilter) Call(doc *goquery.Document) (err error) {
	doc.Find("img").Each(func(i int, node *goquery.Selection) {
		node.SetAttr("style", "max-width: 100%")
	})

	return
}

func TestCustomFilter(t *testing.T) {
	pipe := NewPipeline([]Filter{
		SanitizationFilter{},
		TestFilter{},
	})

	html := `<img onclick="javascript:alert" src="https://google.com/foo.jpg"/>`

	assertCall(t, pipe, `<img src="https://google.com/foo.jpg" style="max-width: 100%"/>`, html)
}
