package pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExternalLInkFilter(t *testing.T) {
	pipe := NewPipeline([]Filter{
		ExternalLinkFilter{
			IgnoreHosts: []string{"ruby-china.org", "*.ruby-china.com"},
		},
	})

	html := `
	<p>Hello <a href="https://ruby-china.com/foo/bar">Link</a></p>
	<p>Hello <a href="https://www.ruby-china.com/foo/bar">Link</a></p>
	<p>Hello <a href="https://fooo.ruby-china.com/foo/bar">Link</a></p>
	<p>Hello <a href="https://ruby-china.org/foo/bar">Link</a></p>
	<p>Hello <a href="https://www.ruby-china.org/foo/bar">Link</a></p>
	<p>Hello <a href="https://l.ruby-china.org/foo/bar">Link</a></p>
	<p>Hello <a href="https://file.github.com/foo/bar">Link</a></p>
	<p>Hello <a href="https://f.google.com/foo/bar">Link</a></p>
	`

	expected := `
	<p>Hello <a href="https://ruby-china.com/foo/bar">Link</a></p>
	<p>Hello <a href="https://www.ruby-china.com/foo/bar">Link</a></p>
	<p>Hello <a href="https://fooo.ruby-china.com/foo/bar">Link</a></p>
	<p>Hello <a href="https://ruby-china.org/foo/bar">Link</a></p>
	<p>Hello <a href="https://www.ruby-china.org/foo/bar">Link</a></p>
	<p>Hello <a href="https://l.ruby-china.org/foo/bar">Link</a></p>
	<p>Hello <a href="https://file.github.com/foo/bar" rel="nofollow" target="_blank">Link</a></p>
	<p>Hello <a href="https://f.google.com/foo/bar" rel="nofollow" target="_blank">Link</a></p>
	`

	out, err := pipe.Call(html)
	assert.NoError(t, err)
	assertHTMLEqual(t, expected, out)

}
