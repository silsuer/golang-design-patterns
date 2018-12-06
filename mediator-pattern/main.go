package main

import "fmt"

// 中介者接口
type IMediator interface {
	Notify(c IClient) // 接受通知
}

// 客户端接口
type IClient interface {
	Column()
}

// 具体中介者，定义关联关系
type Mediator struct {
	From IClient
	To   IClient
}

func (m *Mediator) Notify(c IClient) {
	// 会通知所有
}

func main() {
	fmt.Println(1)
}
