package main

import "fmt"

/*
 * 单例模式之懒汉模式
 * 非线程安全
 */

// 首先构造一个结构体
type example struct {
	name string
}

// 设置一个变量作为单例变量，这是一个私有变量，包外不可访问
var instance *example

// 写一个方法用来返回单例
func GetExample() *example {

	// 存在线程安全问题，高并发时有可能创建多个对象
	if instance == nil {
		instance = new(example)
	}
	return instance
}

func main() {
	s := GetExample()
	s.name = "第一次赋值单例模式"
	fmt.Println(s.name)

	s2 := GetExample()
	fmt.Println(s2.name)
}
