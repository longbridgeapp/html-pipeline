package pipeline

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type TestFilter struct{}

func (f TestFilter) Call(doc *goquery.Document) (err error) {
	doc.Find("img").Each(func(i int, node *goquery.Selection) {
		node.SetAttr("style", "max-width: 100%")
	})

	return
}

func ExamplePipeline_customFilter() {
	/*
		type TestFilter struct{}

		func (f TestFilter) Call(doc *goquery.Document) (err error) {
			doc.Find("img").Each(func(i int, node *goquery.Selection) {
				node.SetAttr("style", "max-width: 100%")
			})

			return
		}
	*/
	pipe := NewPipeline([]Filter{
		SanitizationFilter{},
		TestFilter{},
	})

	html := `<img onclick="javascript:alert" src="https://google.com/foo.jpg"/>`

	out, _ := pipe.Call(html)
	fmt.Println(out)
	// Output:
	// <img src="https://google.com/foo.jpg" style="max-width: 100%"/>
}
