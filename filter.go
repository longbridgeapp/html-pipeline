package pipeline

import "github.com/PuerkitoBio/goquery"

// Filter base filter interface
type Filter interface {
	Call(doc *goquery.Document) (err error)
}
