package pipeline

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
)

// SanitizationFilter use bluemonday default UGCPolicy to sanitize html
type SanitizationFilter struct {
	Policy *bluemonday.Policy
}

func (f SanitizationFilter) Type() string {
	return "string"
}

func (f SanitizationFilter) PolicyWithDefault() *bluemonday.Policy {
	rule := f.Policy
	if rule == nil {
		rule = bluemonday.UGCPolicy()
	}
	return rule
}

func (f SanitizationFilter) Call(doc *goquery.Document) (err error) {
	html, err := doc.Find("body").Html()
	if err != nil {
		return
	}

	doc.Find("body").SetHtml(f.PolicyWithDefault().Sanitize(html))
	return
}
