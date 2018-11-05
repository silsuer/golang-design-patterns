package main

type FactoryInterface interface {
	CreatePigMeatBuns()  // 创建猪肉馅产品
	Create3SBuns()   // 创建三鲜馅产品
}

type ProductInterface interface {
	Intro()
} 

func main() {
}
