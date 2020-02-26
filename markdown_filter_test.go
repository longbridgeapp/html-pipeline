package pipeline

import (
	"fmt"

	"github.com/russross/blackfriday"
)

func ExampleMarkdownFilter() {
	// Custom blackfriday HTML render options
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: blackfriday.UseXHTML |
			blackfriday.NofollowLinks |
			blackfriday.CompletePage,
	})

	// Custom blackfriday extensions
	extensions := blackfriday.Tables |
		blackfriday.FencedCode |
		blackfriday.Autolink |
		blackfriday.Strikethrough |
		blackfriday.SpaceHeadings |
		blackfriday.HardLineBreak |
		blackfriday.NoEmptyLineBeforeBlock

	pipe := NewPipeline([]Filter{
		MarkdownFilter{
			Opts: []blackfriday.Option{
				blackfriday.WithRenderer(renderer),
				blackfriday.WithExtensions(extensions),
			},
		},
		SanitizationFilter{},
	})

	raw := `# Hello world
<script>alert;</script>
<style>body {}</style>

| Name | Location |
| ---- | --- |
| Jason | Chengdu |

This is [html-pipeline](https://github.com/huacnlee/html-pipeline) Markdown filter.`

	out, _ := pipe.Call(raw)
	fmt.Println(out)
	// Output:
	// <h1>Hello world</h1>
	//
	// <p>alert;<br/>
	// body {}</p>
	//
	// <table>
	// <thead>
	// <tr>
	// <th>Name</th>
	// <th>Location</th>
	// </tr>
	// </thead>
	//
	// <tbody>
	// <tr>
	// <td>Jason</td>
	// <td>Chengdu</td>
	// </tr>
	// </tbody>
	// </table>
	// <p>This is <a href="https://github.com/huacnlee/html-pipeline" rel="nofollow">html-pipeline</a> Markdown filter.</p>
}
