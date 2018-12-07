package main

import "fmt"

// 容器接口
type IAggregate interface {
	Iterator() IIterator
}

// 迭代器接口
type IIterator interface {
	HasNext() bool
	Current() int
	Next() bool
}

// 具体容器
type Aggregate struct {
	container []int // 容器中装载 int 型容器
}

type Iterator struct {
	cursor    int // 当前游标
	aggregate Aggregate
}

func (i *Iterator) HasNext() bool {
	if i.cursor+1 < len(i.aggregate.container) {
		return true
	}
	return false
}

func (i *Iterator) Current() int {
	return i.aggregate.container[i.cursor]
}
func (i *Iterator) Next() bool {
	if i.cursor < len(i.aggregate.container) {
		i.cursor++
		return true
	}
	return false
}

func (a *Aggregate) Iterator() IIterator {
	i := new(Iterator)
	i.aggregate = a
	return i
}

func main() {
	// 创建容器，并放入初始化数据
	c := &Aggregate{container: []int{1, 2, 3, 4}}
	iterator := c.Iterator()
	for {
		fmt.Println(iterator.Current())
		if iterator.HasNext() {
			iterator.Next()
		} else {
			break
		}
	}
}
