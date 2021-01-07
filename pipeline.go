package pipeline

import (
	"bytes"
	"regexp"
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

var (
	stripHTMLTagRe = regexp.MustCompile(`<.+?>`)
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
	if p.mode == modeHTML {
		return p.callWithHTML(raw)
	} else {
		return p.callWithPlain(raw)
	}
}

// Call to Render with Pipleline
func (p Pipeline) callWithHTML(raw string) (out string, err error) {
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

	out, err = doc.Find("body").Html()
	if err != nil {
		return
	}

	if !hasEscapeFilter {
		out = unescapeSingleQuote(out)
	}

	return
}

// CacallWithPlain render plain text
func (p Pipeline) callWithPlain(raw string) (out string, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(raw))
	if err != nil {
		return
	}

	for _, filter := range p.Filters {
		switch filter.(type) {
		case HTMLEscapeFilter:
			continue
		}

		err = filter.Call(doc)
		if err != nil {
			return
		}
	}

	out = getRawHTML(doc.Find("body"))

	// Ensure to remove HTML Tag for avoid XSS
	// Because Plain mode has limited not supports any HTML Tag, so here we can make sure to remove all of them.
	out = stripHTMLTagRe.ReplaceAllString(out, "")

	out = strings.TrimSpace(out)
	return
}

func unescapeSingleQuote(in string) (out string) {
	return strings.ReplaceAll(in, "&#39;", "'")
}

// Text gets the combined text contents of each element in the set of matched
// elements, including their descendants.
// https://github.com/PuerkitoBio/goquery/blob/v1.6.0/property.go#L62
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
