package pipeline

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
)

// MarkdownFilter render Markdown with blackfriday
type MarkdownFilter struct {
	Opts []blackfriday.Option
}

// Call render
func (f MarkdownFilter) Call(doc *goquery.Document) (err error) {
	text := doc.Find("body").Text()
	html := blackfriday.Run([]byte(text), f.Opts...)
	doc.Find("body").SetHtml(string(html))
	return
}
