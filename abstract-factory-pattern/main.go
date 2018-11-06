package main

import "fmt"

type FactoryInterface interface {
	CreatePigMeatBuns() ProductInterface // 创建猪肉馅产品
	Create3SBuns() ProductInterface      // 创建三鲜馅产品
}

type ProductInterface interface {
	Intro()
}

type GDPigMeatBuns struct {
}

func (p GDPigMeatBuns) Intro() {
	fmt.Println("广东猪肉馅包子")
}

type GD3SBuns struct {
}

func (p3 GD3SBuns) Intro() {
	fmt.Println("广东三鲜馅包子")
}

type QSPigMeatBuns struct {
}

func (q QSPigMeatBuns) Intro() {
	fmt.Println("齐市猪肉馅包子")
}

type QS3SBuns struct {
}

func (q3 QS3SBuns) Intro() {
	fmt.Println("齐市三鲜馅包子")
}

//type name

type QSFactory struct {
}

func (qs QSFactory) CreatePigMeatBuns() ProductInterface {
	return QSPigMeatBuns{}
}

func (qs QSFactory) Create3SBuns() ProductInterface {
	return QS3SBuns{}
}

type GDFactory struct {
}

func (gd GDFactory) CreatePigMeatBuns() ProductInterface {
	return GDPigMeatBuns{}
}

func (gd GDFactory) Create3SBuns() ProductInterface {
	return GD3SBuns{}
}

func main() {
	// 创建工厂
	var f FactoryInterface
	f = new(QSFactory)
	b := f.CreatePigMeatBuns()
	b.Intro()
}
