package main

import (
	"fmt"
)

type EnvExample struct {
	Element
}

func (e EnvExample) Print(visitor IVisitor) {
	e.Element.Accept(visitor)
}

func GetEnv() string {
	return "testing"
}

// 定义访问者接口
type IVisitor interface {
	Visit() // 访问者的访问方法
}

type ProductionVisitor struct {
}

func (v ProductionVisitor) Visit() {
	fmt.Println("这是生产环境")
}

type TestingVisitor struct {
}

func (t TestingVisitor) Visit() {
	fmt.Println("这是测试环境")
}

// 定义元素接口
type IElement interface {
	Accept(visitor IVisitor)
}

type Element struct {
}

func (el Element) Accept(visitor IVisitor) {
	visitor.Visit()
}

func main() {
	// 创建一个元素
	e := new(Element)
	e.Accept(new(ProductionVisitor)) // output: 这是生产环境
	e.Accept(new(TestingVisitor))    // output: 这是测试环境

	m := new(EnvExample)
	m.Print(new(ProductionVisitor))
	m.Print(new(TestingVisitor))
}
