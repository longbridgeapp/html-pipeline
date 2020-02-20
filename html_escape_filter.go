package pipeline

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// HTMLEscapeFilter HTML escape for Plain text
type HTMLEscapeFilter struct{}

func (f HTMLEscapeFilter) Call(doc *goquery.Document) (err error) {
	text, err := doc.Find("body").Html()
	if err != nil {
		return err
	}

	text = strings.ReplaceAll(text, ">", "&gt;")
	text = strings.ReplaceAll(text, "<", "&lt;")

	doc.Find("body").SetHtml(text)

	return
}
