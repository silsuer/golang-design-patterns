package main

import "fmt"

type Base struct {

}

func (b Base) Print()  {
	fmt.Println("这是print方法")
}

func (b Base) Echo()  {

}

type Son struct {
	Base
}

func (s Son) Echo()  {
	fmt.Println("这是Echo方法")
}


func main() {
	// 模板模式
	s := new(Son)
	s.Print()
	s.Echo()
}

