package html

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
)

// SanitizationFilter use bluemonday default UGCPolicy to sanitize html
type SanitizationFilter struct {
	Rule *bluemonday.Policy
}

func (f SanitizationFilter) Call(doc *goquery.Document) (err error) {
	html, err := doc.Html()
	if err != nil {
		return
	}
	rule := f.Rule
	if rule == nil {
		rule = bluemonday.UGCPolicy()
	}
	doc.ReplaceWithHtml(rule.Sanitize(html))

	return
}
