package pipeline

import "fmt"

func ExampleAutoCorrectFilter() {
	raw := "<p>演示html-pipeline实现自动修正空格</p><p>这是第2个段落</p>"

	pipe := NewPipeline([]Filter{
		AutoCorrectFilter{},
	})

	out, _ := pipe.Call(raw)
	fmt.Println(out)
	// Output:
	// <p>演示 html-pipeline 实现自动修正空格</p><p>这是第 2 个段落</p>
}

func ExampleAutoCorrectFilter1() {
	raw := "演示html-pipeline实现自动修正空格"

	pipe := NewPipeline([]Filter{
		AutoCorrectFilter{},
	})

	out, _ := pipe.Call(raw)
	fmt.Println(out)
	// Output:
	// 演示 html-pipeline 实现自动修正空格
}
