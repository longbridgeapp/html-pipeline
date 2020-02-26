package pipeline

import (
	"github.com/PuerkitoBio/goquery"
	autocorrect "github.com/huacnlee/go-auto-correct"
)

// AutoCorrectFilter Automatically add whitespace between CJK and
// half-width characters (alphabetical letters, numerical digits and symbols).
type AutoCorrectFilter struct{}

// Call AutoCorrectFilter
func (f AutoCorrectFilter) Call(doc *goquery.Document) (err error) {
	html, err := doc.Find("body").Html()
	if err != nil {
		return err
	}

	html = autocorrect.Format(html)

	doc.Find("body").SetHtml(html)

	return
}
