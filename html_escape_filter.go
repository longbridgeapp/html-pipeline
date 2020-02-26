package pipeline

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// HTMLEscapeFilter HTML escape for Plain text
type HTMLEscapeFilter struct{}

// Call HTMLEscapeFilter
func (f HTMLEscapeFilter) Call(doc *goquery.Document) (err error) {
	html, err := doc.Find("body").Html()
	if err != nil {
		return err
	}

	html = strings.ReplaceAll(html, ">", "&gt;")
	html = strings.ReplaceAll(html, "<", "&lt;")

	doc.Find("body").SetHtml(html)

	return
}
