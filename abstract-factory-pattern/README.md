# 抽象工厂模式

## 简介

> wiki: 抽象工厂模式（英语：Abstract factory pattern）是一种软件开发设计模式。抽象工厂模式提供了一种方式，可以将一组具有同一主题的单独的工厂封装起来。在正常使用中，客户端程序需要创建抽象工厂的具体实现，然后使用抽象工厂作为接口来创建这一主题的具体对象。客户端程序不需要知道（或关心）它从这些内部的工厂方法中获得对象的具体类型，因为客户端程序仅使用这些对象的通用接口。抽象工厂模式将一组对象的实现细节与他们的一般使用分离开来。

在前面的文章中我们介绍过工厂模式的前两种:

  - [简单工厂模式](https://juejin.im/post/5bdbcc08f265da61561eb493)
  
  - [工厂方法模式](https://juejin.im/post/5bdede60518825171a180c44)
  
而今天要介绍的 `抽象工厂模式` 就是工厂模式中的最后一种了，是前两种模式的补充，

工厂模式用来创建一组相关或者相互依赖的对象，与工厂方法模式的区别就在于，工厂方法模式针对的是一个产品等级结构；而抽象工厂模式则是针对的多个产品等级结构，
我们可以将一种产品等级想象为一个产品族，所谓的产品族，是指位于不同产品等级结构中功能相关联的产品组成的家族。

## 代码实现

还是以 `工厂方法模式` 中我们举的卖包子的例子:

在之前的工厂方法中，需要在齐市和广东开两家包子店，那么就需要从一个工厂接口中实现两个工厂结构体，齐市店和广东店属于两个产品族，猪肉包和三鲜馅包子属于同一个等级结构，

所以在抽象工厂模式中，我们要添加两个工厂，每个工厂实现两个产品的创建方法:

1. 工厂接口和产品接口

   ```go
      type FactoryInterface interface {
      	CreatePigMeatBuns() ProductInterface // 创建猪肉馅产品
      	Create3SBuns() ProductInterface      // 创建三鲜馅产品
      }
      
      type ProductInterface interface {
      	Intro()
      }
   ```
   
2. 实现4种产品
  
  ```go
     
    type GDPigMeatBuns struct {
    }

    func (p GDPigMeatBuns) Intro() {
	    fmt.Println("广东猪肉馅包子")
    }
    // TODO ... 其他产品实现方法没区别... 就省略掉了，需要的话请去仓库里看源码
  ```
  
3. 实现工厂

  ```go
    // 齐市包子铺 
    type QSFactory struct {
    }

    func (qs QSFactory) CreatePigMeatBuns() ProductInterface {
	    return QSPigMeatBuns{}
    }

    func (qs QSFactory) Create3SBuns() ProductInterface {
	    return QS3SBuns{}
    }
    // 广东包子铺
    type GDFactory struct {
    }

    func (gd GDFactory) CreatePigMeatBuns() ProductInterface {
	    return GDPigMeatBuns{}
    }

    func (gd GDFactory) Create3SBuns() ProductInterface {
	    return GD3SBuns{}
    }
  ```

4. 这样就可以通过抽象工厂创建了

   ```go
        var f FactoryInterface  // 特意以这种方式声明，更好的体会抽象工厂模式的好处
      	f = new(QSFactory)  
      	b := f.CreatePigMeatBuns()  
      	b.Intro()
   ```

## 优缺点

  - 优点: 抽象工厂模式除了具有工厂方法模式的优点外，最主要的优点就是可以在类的内部对产品族进行约束。所谓的产品族，一般或多或少的都存在一定的关联，抽象工厂模式就可以在类内部对产品族的关联关系进行定义和描述，而不必专门引入一个新的类来进行管理。
  
  - 缺点: 产品族的扩展将是一件十分费力的事情，假如产品族中需要增加一个新的产品，则几乎所有的工厂类都需要进行修改。所以使用抽象工厂模式时，对产品等级结构的划分是非常重要的
  
## 总结

  现在我们就讲完了三种工厂模式的构建了，他们之间有区别又有联系，具体使用什么模式，就见仁见智了，经常性的，当业务发展过程中，会从简单工厂模式一步一步变成工厂方法，或者抽象工厂模式
  
  创建型设计模式的结果都是得到指定对象，模式之间没有好坏，按需使用，只要能够达到解耦的目的，就是好模式。
  

> 上述代码均放在 [golang-design-patterns](https://github.com/silsuer/golang-design-patterns) 这个仓库中

> 打个广告，推荐一下自己写的 go web框架 [bingo](https://github.com/silsuer/bingo),求star，求PR ~
