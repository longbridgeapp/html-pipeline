package pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTMLEscapeFilter(t *testing.T) {
	raw := `<div>Hello "Foo's Bar"</div>`

	pipe := NewPipeline([]Filter{
		HTMLEscapeFilter{},
	})

	out, err := pipe.Call(raw)
	assert.NoError(t, err)
	assert.Equal(t, "&lt;div&gt;Hello &#34;Foo&#39;s Bar&#34;&lt;/div&gt;", out)
}

func BenchmarkHTMLEscapeFilter(b *testing.B) {
	raw := `<div>Hello "Foo's Bar"</div>`
	pipe := NewPipeline([]Filter{
		HTMLEscapeFilter{},
	})

	for i := 0; i < b.N; i++ {
		// 7001 ns/op
		pipe.Call(raw)
	}
}
