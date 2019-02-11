package strategy_pattern

import "fmt"

// 1. 定义抽象策略接口
type IStrategy interface {
	SortList() // 对列表进行排序
}


// 2. 定义具体策略
type BubbleSortStrategy struct {}

func ( b BubbleSortStrategy) SortList()  {
	fmt.Println("这是冒泡排序")
}

type MergeSortStrategy struct {}

func (m MergeSortStrategy) SortList()  {
	fmt.Println("这是归并排序")
}

// 3. 定义上下文
type Context struct {
	Strategy IStrategy  // 上下文中指定的策略
}

// 定义上下文中执行策略的方法
func (c Context) Exec() {
	c.Strategy.SortList()
}

// 策略模式
func main() {
  var ctx Context
  fmt.Println("====使用冒泡排序算法=====")
  ctx = Context{Strategy:BubbleSortStrategy{}}
  ctx.Exec()

  fmt.Println("====使用归并排序算法=====")
  ctx = Context{Strategy:MergeSortStrategy{}}
  ctx.Exec()
}
