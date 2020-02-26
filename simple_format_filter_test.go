package pipeline

import "testing"

func TestSimpleFormatFilter(t *testing.T) {
	pipeline := NewPipeline([]Filter{
		SimpleFormatFilter{},
	})

	raw := "Hello world"
	expected := "<p>Hello world</p>"
	assertCall(t, pipeline, expected, raw)

	raw = "Hello world\n\nThis is a document.\n哈哈哈"
	expected = "<p>Hello world</p><p>This is a document.\n<br/>哈哈哈</p>"
	assertCall(t, pipeline, expected, raw)
}

func BenchmarkSimpleFormatFilter(b *testing.B) {
	raw := "Hello world\n\nThis is a document.\n哈哈哈"
	pipe := NewPipeline([]Filter{
		SimpleFormatFilter{},
	})

	for i := 0; i < b.N; i++ {
		// 8595 ns/op
		pipe.Call(raw)
	}
}
