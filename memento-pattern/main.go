package main

import "fmt"

// 发起人
type Originator struct {
	state string // 这里就简单一点，要保存的状态就是一个字符串	
}

func (o *Originator) SetState(s string) {
	o.state = s
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{state: o.state}
}

// 备忘录
type Memento struct {
	state string // 这里就是保存的状态
}

func (m *Memento) SetState(s string) {
	m.state = s
}

func (m *Memento) GetState() string {
	return m.state
}

// 负责人
type Caretaker struct {
	memento *Memento
}

func (c *Caretaker) GetMemento() *Memento {
	return c.memento
}

func (c *Caretaker) SetMemento(m *Memento) {
	c.memento = m
}

func main() {
	// 创建一个发起人并设置初始状态
	// 此时与备忘录模式无关，只是模拟正常程序运行
	o := &Originator{state: "hello"}
    fmt.Println("当前状态:",o.GetState())
	// 现在需要保存当前状态
	// 就创建一个负责人来设置（一般来说，对于一个对象的同一个备忘范围，应当只有一个负责人，这样方便做多状态多备忘管理）
	c := new(Caretaker)
	c.SetMemento(o.CreateMemento())

	o.SetState("world")
	fmt.Println("更改当前状态:",o.GetState())

	// 恢复备忘
	o.SetState(c.GetMemento().GetState())
	fmt.Println("恢复后状态",o.GetState())
}
