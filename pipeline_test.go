package pipeline

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	htmlSpaceRe = regexp.MustCompile(`>[\s]+<`)
)

func assertHTMLEqual(t *testing.T, exptected, actual string) {
	exptected = strings.TrimSpace(exptected)
	actual = strings.TrimSpace(actual)

	if htmlSpaceRe.ReplaceAllString(exptected, "><") != htmlSpaceRe.ReplaceAllString(actual, "><") {
		t.Errorf("\nexptected:\n%s\nactual   :\n%s", exptected, actual)
	}
}

func assertCall(t *testing.T, pipeline Pipeline, exptected, raw string) {
	out, err := pipeline.Call(raw)
	assert.NoError(t, err)
	assert.Equal(t, exptected, out)
}

func TestNewPipeline(t *testing.T) {
	pipeline := NewPipeline([]Filter{
		&SanitizationFilter{},
	})

	assert.Equal(t, 1, len(pipeline.Filters))

	out, err := pipeline.Call("<p>Hello world<script>alert</script></p>")
	assert.NoError(t, err)
	assert.Equal(t, "<p>Hello world</p>", out)
}

func BenchmarkMultiplePiplelines(b *testing.B) {
	raw := `#huacnlee This is a #test_huacn-lee of some cool #中文名称 features that #mi_asd be
	#use-ful but #don't. look at this email#address.com. #bla! I like #nylas but I don't
	like to go to this apple.com?a#url. I also don't like the comment blocks.
	But #msft is cool. #huacnlee This is a #test_huacn-lee of some cool #中文名称 features that #mi_asd be
	#use-ful but #don't. look at this email#address.com. #bla! I like #nylas but I don't
	like to go to this apple.com?a#url. I also don't like the comment blocks.
	But #msft is cool. #huacnlee This is a #test_huacn-lee of some cool #中文名称 features that #mi_asd be
	#use-ful but #don't. look at this email#address.com. #bla! I like #nylas but I don't
	like to go to this apple.com?a#url. I also don't like the comment blocks.
	But #msft is cool.`

	pipe := NewPipeline([]Filter{
		HTMLEscapeFilter{},
		SimpleFormatFilter{},
		MentionFilter{},
	})

	for i := 0; i < b.N; i++ {
		// 41251 ns/op
		pipe.Call(raw)
	}
}

func ExamplePipeline() {
	pipe := NewPipeline([]Filter{
		MarkdownFilter{},
		SanitizationFilter{},
		MentionFilter{
			Prefix: "#",
			Format: func(name string) string {
				return fmt.Sprintf(`<a href="https://github.com/topic/%s">#%s</a>`, name, name)
			},
		},
		MentionFilter{
			Prefix: "@",
			Format: func(name string) string {
				return fmt.Sprintf(`<a href="https://github.com/%s">@%s</a>`, name, name)
			},
		},
	})

	markdown := `# Hello world

![](javascript:alert) [Click me](javascript:alert)

This is #html-pipeline example, @huacnlee created.`
	out, _ := pipe.Call(markdown)
	fmt.Printf(out)
	// Output:
	// <h1>Hello world</h1>
	//
	// <p><img alt=""/> Click me</p>
	//
	// <p>This is <a href="https://github.com/topic/html-pipeline">#html-pipeline</a> example, <a href="https://github.com/huacnlee">@huacnlee</a> created.</p>
}

func TestHTMLUnescape(t *testing.T) {
	raw := "We don't like 'escape' and 'unescape'."
	pipe := NewPipeline([]Filter{})
	out, _ := pipe.Call(raw)
	assert.Equal(t, raw, out)

	raw = "<object props=\"{&quot;url&quot;: &quot;https://example.com/a.jpg&quot;}\">We don't like 'escape'</object>"
	out, _ = pipe.Call(raw)
	assert.Equal(t, `<object props="{&#34;url&#34;: &#34;https://example.com/a.jpg&#34;}">We don't like 'escape'</object>`, out)
}

func TestRenderPlainText(t *testing.T) {
	raw := `
	[tag value="qcc"]Foo#QuantumScape & Corporation Class A[/tag] @huacnlee with 'code',
	<script>alert()</script> &lt;style&gt;style tag&lt;/style&gt; <mark value="Bar">HTML mark will not remove</mark> recenty pressed "News".
	https://www.google.com/search?newwindow=1&sxsrf=ALeKk01IaJz5BXWn2_C3_AFY3m_NL0c0pQ%3A1609989927324&ei=J3_2X76rE66zmAWLxIfoCg&q=Complex+url+%E4%B8%AD%E6%96%87&oq=Complex+url+%E4%B8%AD%E6%96%87&gs_lcp=CgZwc3ktYWIQAzoFCAAQywE6BAgAEB46BggAEAUQHjoGCAAQCBAeOgUIIRCgAVDJO1jORGDyRWgAcAB4AoAB_AGIAeQNkgEFMC4zLjaYAQCgAQGqAQdnd3Mtd2l6wAEB&sclient=psy-ab&ved=0ahUKEwj-2tbt74juAhWuGaYKHQviAa0Q4dUDCA0&uact=5
	`

	expected := `
	[tag value="qcc"]Foo#QuantumScape & Corporation Class A[/tag] [mt value="huacnlee"]@huacnlee[/mt] with 'code',
	alert() style tag HTML mark will not remove recenty pressed "News".
	https://www.google.com/search?newwindow=1&sxsrf=ALeKk01IaJz5BXWn2_C3_AFY3m_NL0c0pQ%3A1609989927324&ei=J3_2X76rE66zmAWLxIfoCg&q=Complex+url+%E4%B8%AD%E6%96%87&oq=Complex+url+%E4%B8%AD%E6%96%87&gs_lcp=CgZwc3ktYWIQAzoFCAAQywE6BAgAEB46BggAEAUQHjoGCAAQCBAeOgUIIRCgAVDJO1jORGDyRWgAcAB4AoAB_AGIAeQNkgEFMC4zLjaYAQCgAQGqAQdnd3Mtd2l6wAEB&sclient=psy-ab&ved=0ahUKEwj-2tbt74juAhWuGaYKHQviAa0Q4dUDCA0&uact=5
	`

	// With HTMLEscapeFilter (will ignore)
	pipe := NewPlainPipeline([]Filter{
		HTMLEscapeFilter{},
		MentionFilter{
			Format: func(name string) string {
				return fmt.Sprintf(`[mt value="%s"]@%s[/mt]`, name, name)
			},
		},
	})
	out, _ := pipe.Call(raw)
	assert.Equal(t, strings.TrimSpace(expected), out)

	// Without HTMLEscapeFilter
	pipe = NewPlainPipeline([]Filter{
		MentionFilter{
			Format: func(name string) string {
				return fmt.Sprintf(`[mt value="%s"]@%s[/mt]`, name, name)
			},
		},
	})
	out, _ = pipe.Call(raw)
	assert.Equal(t, strings.TrimSpace(expected), out)
}
