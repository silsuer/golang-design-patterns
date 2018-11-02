package main

import "fmt"

// 简单工厂模式

// 产品接口，所有可以通过工厂实例创建的对象都会实现这个接口
type Product interface {
	create()
}

// 工厂类
type Factory struct {
}

// 工厂创建产品的方法，传入一个产品名，返回这个产品的实例
func (f Factory) Generate(name string) Product {
	switch name {
	case "product1":
		return Product1{}
	case "product2":
		return Product2{}
	default:
		return nil
	}
}

// 产品1，实现产品接口
type Product1 struct {
}

func (p1 Product1) create() {
	fmt.Println("this is product 1")
}

// 产品2，实现产品接口
type Product2 struct {
}

func (p1 Product2) create() {
	fmt.Println("this is product 2")
}

func main() {

	// 创建一个工厂类，在应用中可以将这个工厂类实例作为一个全局变量
	factory := new(Factory)

	// 在工厂类中传入不同的参数，获取不同的实例
	p1 := factory.Generate("product1")
	p1.create() // output:   this is product 1

	p2 := factory.Generate("product2")
	p2.create() // output:   this is product 2
}
