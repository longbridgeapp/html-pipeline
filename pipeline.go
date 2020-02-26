package pipeline

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Pipeline stuct
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
