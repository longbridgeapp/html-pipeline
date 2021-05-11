package pipeline

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ignoreHosts = []string{"*.ruby-china.com", "ruby-china.org", "localhost"}
)

func Test_ImageURLFilter_isHost(t *testing.T) {
	hostsWillIgnore := []string{
		"https://ruby-china.com/foo.jpg",
		"https://www.ruby-china.com/foo.jpg",
		"https://www.Ruby-china.com/foo.jpg",
		"https://Ruby-china.org/foo.jpg",
		"https://www.Ruby-china.org/foo.jpg",
		"https://localhost/foo.jpg",
		"https://localhost:3000/foo.jpg",
	}
	for _, host := range hostsWillIgnore {
		if !isHost(ignoreHosts, host) {
			t.Errorf("%s not match", host)
		}
	}

	hostsWillNotIgnore := []string{
		"https://baidu.com/foo.jpg",
		"https://aaa.com/foo.jpg",
	}
	for _, host := range hostsWillNotIgnore {
		if isHost(ignoreHosts, host) {
			t.Errorf("%s matched, but not expected", host)
		}
	}
}

func Test_ImageURLFilter(t *testing.T) {
	pipe := NewPipeline([]Filter{
		ImageURLFilter{
			IgnoreHosts: ignoreHosts,
			Format: func(src string) string {
				return fmt.Sprintf("https://imageproxy.ruby-china.com/%s", src)
			},
		},
	})

	html := `
	<p>Hello <img src="https://ruby-china.com/test/image.jpg"/></p>
	<p>Hello <img src="https://www.ruby-china.com/test/image.jpg"/></p>
	<p>Hello <img src="https://fooo.ruby-china.com/test/image.jpg"/></p>
	<p>Hello <img src="https://ruby-china.org/test/image.jpg"/></p>
	<p>Hello <img src="https://www.ruby-china.org/test/image.jpg"/></p>
	<p>Hello <img src="https://l.ruby-china.org/test/image.jpg"/></p>
	<p>Hello <img src="https://localhost/test/image.jpg"/></p>
	<p>Hello <img src="https://localhost:3000/test/image.jpg"/></p>
	<p>Hello <img src="https://file.github.com/test/image.jpg"/></p>
	<p>Hello <img src="https://f.google.com/test/image.jpg"/></p>
	`

	expected := `
	<p>Hello <img src="https://ruby-china.com/test/image.jpg"/></p>
	<p>Hello <img src="https://www.ruby-china.com/test/image.jpg"/></p>
	<p>Hello <img src="https://fooo.ruby-china.com/test/image.jpg"/></p>
	<p>Hello <img src="https://ruby-china.org/test/image.jpg"/></p>
	<p>Hello <img src="https://www.ruby-china.org/test/image.jpg"/></p>
	<p>Hello <img src="https://l.ruby-china.org/test/image.jpg"/></p>
	<p>Hello <img src="https://localhost/test/image.jpg"/></p>
	<p>Hello <img src="https://localhost:3000/test/image.jpg"/></p>
	<p>Hello <img src="https://imageproxy.ruby-china.com/https://file.github.com/test/image.jpg"/></p>
	<p>Hello <img src="https://imageproxy.ruby-china.com/https://f.google.com/test/image.jpg"/></p>
	`

	out, err := pipe.Call(html)
	assert.NoError(t, err)
	assertHTMLEqual(t, expected, out)

	// Match and ignore
	pipe = NewPipeline([]Filter{
		ImageURLFilter{
			MatchHosts: []string{"github.com", "google.com"},
			Format: func(src string) string {
				return fmt.Sprintf("https://imageproxy.ruby-china.com/%s", src)
			},
		},
	})

	out, err = pipe.Call(html)
	assert.NoError(t, err)
	assertHTMLEqual(t, expected, out)

}
