package pipeline

import (
	"testing"

	"github.com/russross/blackfriday"
)

func TestMarkdownFilter(t *testing.T) {
	pipe := NewPipeline([]Filter{
		MarkdownFilter{
			Opts: []blackfriday.Option{
				blackfriday.WithExtensions(blackfriday.AutoHeadingIDs),
			},
		},
		SanitizationFilter{},
	})

	raw := `# Hello world
<script>alert;</script>
<style>body {}</style>

This is [html-pipeline](https://github.com/huacnlee/html-pipeline) Markdown filter.`

	out := `<h1 id="hello-world">Hello world</h1>

<p>alert;
body {}</p>

<p>This is <a href="https://github.com/huacnlee/html-pipeline" rel="nofollow">html-pipeline</a> Markdown filter.</p>
`

	assertCall(t, pipe, out, raw)
}
