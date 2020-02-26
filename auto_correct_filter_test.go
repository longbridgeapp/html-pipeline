package pipeline

import "fmt"

func ExampleAutoCorrectFilter() {
	raw := "演示html-pipeline实现自动修正空格"

	pipe := NewPipeline([]Filter{
		AutoCorrectFilter{},
	})

	out, _ := pipe.Call(raw)
	fmt.Println(out)
	// Output:
	// 演示 html-pipeline 实现自动修正空格
}
