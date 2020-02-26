package pipeline

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	simpleFormatAllBreak      = regexp.MustCompile("\r\n?")
	simpleFormatParagraphMark = regexp.MustCompile("\n\n+")
)

// SimpleFormatFilter covnert simple plain text into breakable html
type SimpleFormatFilter struct {
}

func (f SimpleFormatFilter) Call(doc *goquery.Document) (err error) {
	html, err := doc.Find("body").Html()
	if err != nil {
		return
	}

	outs := []string{}
	for _, paragraph := range f.splitParagraphs(html) {
		paragraph = strings.ReplaceAll(paragraph, "\n", "<br/>")
		outs = append(outs, "<p>"+paragraph+"</p>")
	}
	html = strings.Join(outs, "")

	doc.Find("body").SetHtml(html)

	return
}

func (f SimpleFormatFilter) splitParagraphs(html string) (paragraphs []string) {
	html = simpleFormatAllBreak.ReplaceAllString(html, "\n")
	paragraphs = simpleFormatParagraphMark.Split(html, -1)
	return
}
