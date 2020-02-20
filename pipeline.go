package pipeline

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type Pipeline struct {
	Filters []Filter
}

// NewPipeline create a new pipeline
func NewPipeline(filters []Filter) Pipeline {
	return Pipeline{
		Filters: filters,
	}
}

// Call to Render with Pipleline
func (p Pipeline) Call(html string) (out string, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return
	}

	for _, filter := range p.Filters {
		err = filter.Call(doc)
		if err != nil {
			return
		}
	}

	out, err = doc.Find("body").Html()
	if err != nil {
		return
	}

	return
}

// TraverseTextNodes map nested node to find all text node
func TraverseTextNodes(node *html.Node, fn func(*html.Node)) {
	if node == nil {
		return
	}
	if node.Type == html.TextNode {
		fn(node)
	}

	cur := node.FirstChild

	for cur != nil {
		next := cur.NextSibling
		TraverseTextNodes(cur, fn)
		cur = next
	}
}
