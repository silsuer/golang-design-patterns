package main

import "fmt"

// 示例结构体
type Example struct {
	Content string
}

func (e *Example) Clone() *Example {
	res := *e
	return &res
}

func main() {
	r1 := new(Example)
	r1.Content = "this is example 1"
	r2 := r1.Clone()
	r2.Content = "this is example 2"

	fmt.Println(r1)
	fmt.Println(r2)

}
