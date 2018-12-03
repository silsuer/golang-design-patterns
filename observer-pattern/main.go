package main

import "fmt"

// 抽象观察者
type IObserver interface {
	Notify() // 当被观察对象有更改的时候，出发观察者的Notify() 方法
}

// 抽象被观察者
type ISubject interface {
	AddObservers(observers ...IObserver) // 添加观察者
	NotifyObservers()                    // 通知观察者
}

type Observer struct {
}

func (o *Observer) Notify() {
	fmt.Println("已经触发了观察者")
}

type Subject struct {
	observers []IObserver
}

func (s *Subject) AddObservers(observers ...IObserver) {
	//fmt.Println(1)
	s.observers = append(s.observers, observers...)
	//fmt.Println(s.observers)
}

func (s *Subject) NotifyObservers() {
	for k := range s.observers {
		s.observers[k].Notify() // 触发观察者
	}
}

func main() {
    // 创建被观察者
    s := new(Subject)
    // 创建观察者
    o := new(Observer)
    // 为主题添加观察者
    s.AddObservers(o)

    // 这里的被观察者要做各种更改...

    // 更改完毕，触发观察者
    s.NotifyObservers()
}
