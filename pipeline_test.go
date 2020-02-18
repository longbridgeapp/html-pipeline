package html

import (
	"testing"

	"gotest.tools/assert"
)

func TestNewPipeline(t *testing.T) {
	pipeline := NewPipeline([]Filter{
		SanitizationFilter{},
	})

	assert.Equal(t, 1, len(pipeline.Filters))
}
