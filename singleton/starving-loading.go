package main

import "fmt"

/*
 * 饿汉模式
 * 在类加载时就初始化对象
 */

// 构建一个结构体，用来实例化单例
type example2 struct {
	name string
}

// 声明一个私有变量，作为单例
var instance2 *example2

// init函数将在包初始化时执行，实例化单例
func init() {
	instance2 = new(example2)
	instance2.name = "初始化单例模式"
}

func GetInstance2() *example2 {
	return instance2
}

func main() {
	s := GetInstance2()
	fmt.Println(s.name)
}
