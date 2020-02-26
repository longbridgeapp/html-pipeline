package pipeline

import (
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
