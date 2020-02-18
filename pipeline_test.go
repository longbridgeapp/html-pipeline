package pipeline

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	htmlSpaceRe = regexp.MustCompile(`>[\s]+<`)
)

func assertHTMLEqual(t *testing.T, exptected, actual string) {
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
