package pipeline

import (
	"fmt"
	"testing"
)

func ExampleHTMLEscapeFilter() {
	raw := `<div>Hello "Foo's Bar"</div>`

	pipe := NewPipeline([]Filter{
		HTMLEscapeFilter{},
	})

	out, _ := pipe.Call(raw)
	fmt.Printf(out)
	// Output:
	// &lt;div&gt;Hello &#34;Foo&#39;s Bar&#34;&lt;/div&gt;
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
