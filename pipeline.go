package pipeline

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// Mode for render
type mode int

const (
	// ModeHTML use HTML output
	modeHTML mode = iota
	// ModePlain use Plain text output
	modePlain
)

// Pipeline stuct
type Pipeline struct {
	Filters []Filter
	mode    mode
}

// NewPipeline create pipeline with HTML mode
func NewPipeline(filters []Filter) Pipeline {
	return Pipeline{
		Filters: filters,
		mode:    modeHTML,
	}
}

// NewPlainPipeline create pipeline with Plain mode (HTML tags will remove)
func NewPlainPipeline(filters []Filter) Pipeline {
	return Pipeline{
		Filters: filters,
		mode:    modePlain,
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
			// Skip HTMLEscapeFilter in plain mode
			if p.mode == modePlain {
				continue
			}
		}

		err = filter.Call(doc)
		if err != nil {
			return
		}
	}

	if p.mode == modeHTML {
		out, err = doc.Find("body").Html()
		if err != nil {
			return
		}

		if !hasEscapeFilter {
			out = unescapeSingleQuote(out)
		}
	} else {
		out = getRawHTML(doc.Find("body"))
		out = strings.TrimSpace(out)
	}

	return
}

func unescapeSingleQuote(in string) (out string) {
	return strings.ReplaceAll(in, "&#39;", "'")
}

// Text gets the combined text contents of each element in the set of matched
// elements, including their descendants.
func getRawHTML(s *goquery.Selection) string {
	var buf bytes.Buffer

	// Slightly optimized vs calling Each: no single selection object created
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode || n.Type == html.RawNode {
			// Keep newlines and spaces, like jQuery
			buf.WriteString(n.Data)
		}

		if n.FirstChild != nil {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	for _, n := range s.Nodes {
		f(n)
	}

	return buf.String()
}
