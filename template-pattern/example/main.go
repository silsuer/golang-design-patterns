package main

import (
	"fmt"
)

// 实现一个Base
type Base struct {

}

func (b *Base) Print()  {
	fmt.Println("这是base的print")
}

type Example struct {
	Base
}

func (e Example) Print()  {
	fmt.Println("这是example的print")
}


func main() {
	e := new(Example)
	//e := Example{}
	e.Print()
}