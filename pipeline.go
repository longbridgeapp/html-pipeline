package pipeline

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Pipeline struct {
	Filters []Filter
}

func NewPipeline(filters []Filter) Pipeline {
	return Pipeline{
		Filters: filters,
	}
}

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
