package html

import (
	"github.com/PuerkitoBio/goquery"
)

// Filter for html content
type Filter interface {
	Call(doc *goquery.Document) (err error)
}

// TextFilter for plain text
type TextFilter interface {
	Call(text *string) (err error)
}
