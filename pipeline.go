package pipeline

import (
	"html"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Mode for render
type Mode int

const (
	// ModeHTML use HTML output
	ModeHTML Mode = iota
	// ModePlain use Plain text output
	ModePlain
)

// Pipeline stuct
type Pipeline struct {
	Filters []Filter
	Mode    Mode
}

// NewPipeline create pipeline with HTML mode
func NewPipeline(filters []Filter) Pipeline {
	return Pipeline{
		Filters: filters,
		Mode:    ModeHTML,
	}
}

// NewPlainPipeline create pipeline with Plain mode
func NewPlainPipeline(filters []Filter) Pipeline {
	return Pipeline{
		Filters: filters,
		Mode:    ModePlain,
	}
}

// Call to Render with Pipleline
func (p Pipeline) Call(raw string) (out string, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(raw))
	if err != nil {
		return
	}

	var hasEscapeFilter = false

	for _, filter := range p.Filters {
		switch filter.(type) {
		case HTMLEscapeFilter:
			hasEscapeFilter = true
		}

		err = filter.Call(doc)
		if err != nil {
			return
		}
	}

	if p.Mode == ModeHTML {
		out, err = doc.Find("body").Html()
		if err != nil {
			return
		}

		if !hasEscapeFilter {
			out = unescapeSingleQuote(out)
		}
	} else {
		out, err = doc.Find("body").Html()
		if err != nil {
			return
		}

		out = html.UnescapeString(out)
		out = strings.TrimSpace(out)
	}

	return
}

func unescapeSingleQuote(in string) (out string) {
	return strings.ReplaceAll(in, "&#39;", "'")
}
