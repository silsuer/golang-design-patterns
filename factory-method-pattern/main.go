package main

import "fmt"

type FactoryInterface interface {
	CreateProduct(t string) ProductInterface
}

// 创建工厂结构体并实现工厂接口
type Factory1 struct {
}

func (f Factory1) CreateProduct(t string) ProductInterface {
	switch t {
	case "product1":
		return Product1{}
	default:
		return nil
	}

}

// 产品接口
type ProductInterface interface {
	Intro()
}

// 创建产品1并实现产品接口
type Product1 struct {
}

func (p Product1) Intro() {
	fmt.Println("this is product 1")
}

func main() {
	// 创建工厂
	f := new(Factory1)

	p := f.CreateProduct("product1")
	p.Intro()
}
