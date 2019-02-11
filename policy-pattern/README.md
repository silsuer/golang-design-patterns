# 策略模式

## 定义

> wiki: 定义一组算法，将每个算法都封装起来，并且使他们之间可以互换

在实际应用中， 我们对不同的场景要采取不同的应对措施，也就是不同的策略，比如一个对数据排序的方法，根据数据量和数据特征的不同，我们需要调用不同的排序方法，我们可以把所有的排序算法都封装在同一个函数中，然后通过`if...else`的形式来调用不同的排序算法，这种方式称之为硬编码，可是在实际应用中，功能和体量的不断增长就会使得我们要经常修改源代码，让这个函数越来越难以维护。所以还是为了解耦，策略模式定义一些独立的类来封装不同的算法，每一个类封装一个具体的算法（即策略），策略模式和模板模式有些相似，需要定义一个抽象类来作为策略的基本模板，每一种策略就是这个抽象类延伸出来的具体类来。

## 角色

 - Context: 上下文环境
 
 - Strategy: 抽象策略类
 
 - ConcreteStrategy: 具体策略类

## 类图

  ![](https://design-patterns.readthedocs.io/zh_CN/latest/_images/Strategy.jpg)
  
  从类图可以看出策略模式和模板模式的相似，只是多了一个 上下文(`Context`)来控制使用不同的策略

## 举个栗子

  还用上面说的选择排序算法的栗子:
  
  1. 定义抽象策略接口
  
   ```
     type IStrategy interface {
   	    SortList() // 对列表进行排序
     }
   ```
  2. 定义具体策略
  
   ```
    // 这里定义了冒泡排序和归并排序两种策略
    type BubbleSortStrategy struct {}
    
    func ( b BubbleSortStrategy) SortList()  {
    	fmt.Println("这是冒泡排序")
    }
    
    type MergeSortStrategy struct {}
    
    func (m MergeSortStrategy) SortList()  {
    	fmt.Println("这是归并排序")
    }
   ```
 
  3. 定义上下文
   
   ```
     type Context struct {
     	Strategy IStrategy  // 上下文中指定的策略
     }
     
     func (c Context) Exec() {
     	c.Strategy.SortList()
     }
   ```
   
  4. 开始使用
  
   ```
     // 载入不同的策略，就可以使用不同的算法
     func main() {
       var ctx Context
       fmt.Println("====使用冒泡排序算法=====")
       ctx = Context{Strategy:BubbleSortStrategy{}}
       ctx.Exec()
     
       fmt.Println("====使用归并排序算法=====")
       ctx = Context{Strategy:MergeSortStrategy{}}
       ctx.Exec()
     }
   ```