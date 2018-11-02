# 简单工厂模式
       
>   wiki: 简单工厂模式并不属于 GoF 23 个经典设计模式，但通常将它作为学习其他工厂模式的基础，它的设计思想很简单，其基本流程如下：
        
>   首先将需要创建的各种不同对象（例如各种不同的 Chart 对象）的相关代码封装到不同的类中，这些类称为具体产品类，而将它们公共的代码进行抽象和提取后封装在一个抽象产品类中，每一个具体产品类都是抽象产品类的子类；然后提供一个工厂类用于创建各种产品，在工厂类中提供一个创建产品的工厂方法，该方法可以根据所传入的参数不同创建不同的具体产品对象；客户端只需调用工厂类的工厂方法并传入相应的参数即可得到一个产品对象。
        
>   简单工厂模式定义如下：
        
>   简单工厂模式（Simple Factory Pattern）：定义一个工厂类，它可以根据参数的不同返回不同类的实例，被创建的实例通常都具有共同的父类。因为在简单工厂模式中用于创建实例的方法是静态（static）方法，因此简单工厂模式又被称为静态工厂方法（Static Factory Method）模式，它属于类创建型模式。
        
>   简单工厂模式的要点在于：当你需要什么，只需要传入一个正确的参数，就可以获取你所需要的对象，而无须知道其创建细节。简单工厂模式结构比较简单，其核心是工厂类的设计，其结构如图所示

![](http://wiki.jikexueyuan.com/project/design-pattern-creation/images/20130711143612921.jpg)

上面都是我抄来的...

大概要做的事情就是，当我们想要创建一个对象的时候，调用同一个方法，传入不同的参数就可以返回给我们不同的对象了

当然，前提是这些对象对应的类都实现了相同的接口

**例如：**

我们创建一个工厂结构体，并且创建一个产品接口，工厂可以创建产品，只要在工厂的某个方法中传入不同的参数，就可以返回实现产品接口的不同的对象，

1. 创建工厂结构体:

  ```go
     type Factory struct {
     }
  ```

2. 创建产品接口，这里为了方便，只写了一个方法，请根据自己的需要扩展

  ```go
     type Product interface {
     	create()
     }
  ```
  
3. 创建两个产品：产品1和产品2，它们实现了产品接口:

  ```go
    
  // 产品1，实现产品接口
  type Product1 struct {
  }

  func (p1 Product1) create() {
    	fmt.Println("this is product 1")
  }

  // 产品2，实现产品接口
  type Product2 struct {
  }

  func (p1 Product2) create() {
    	fmt.Println("this is product 2")
  }

  ```

2. 为工厂结构体添加一个方法用于生产产品（实例化对象）:
  
  ```go     
   func (f Factory) Generate(name string) Product {
	   switch name {
	   case "product1":
		   return Product1{}
	   case "product2":
		   return Product2{}
	   default:
		   return nil
	   }
   }   
  ```
  
3. 这样就可以通过传入不同的方法得到不同的产品实例了:

  ```go
        // 创建一个工厂类，在应用中可以将这个工厂类实例作为一个全局变量
      	factory := new(Factory)
      
      	// 在工厂类中传入不同的参数，获取不同的实例
      	p1 := factory.Generate("product1")
      	p1.create() // output:   this is product 1
      
      	p2 := factory.Generate("product2")
      	p2.create() // output:   this is product 2
  ```
  
 上面的例子只是为了解释工厂模式的思想而设置的最简单的例子，下面举一个在实际中应用的例子：
 
   `bingo-log` 是一个`go`语言的日志包，可以自定义日志输出格式，这里就用到了简单工厂模式，所有实现了 `Connector` 接口的结构体都可以作为参数传入日志结构体中，达到自定义输出格式的目的
   
   项目地址: [bingo-log](https://github.com/silsuer/bingo-log)
   
   思路解析: [基于go开发日志处理包](https://juejin.im/post/5bcd796f51882577b82ffaee)
 
   请直接去项目 `README.md` 中查看使用方法，去思路解析中查看整体的设计思路
   
 下面说说工厂模式的优缺点:
 
  - 优点: 工厂类是整个工厂模式的核心，我们只需要传入给定的信息，就可以创建所需实例，在多人协作的时候，无需知道对象之间的内部依赖，可以直接创建，有利于整个软件体系结构的优化
  
  - 缺点: 工厂类中包含了所有实例的创建逻辑，一旦这个工厂类出现问题，所有实例都会受到影响，并且，工厂类中生产的产品都基于一个共同的接口，一旦要添加不同种类的产品，这就会增加工厂类的复杂度，将不同种类的产品混合在一起，违背了单一职责，系统的灵活性和可维护性都会降低，并且当新增产品的时候，必须要修改工厂类，违背了『系统对扩展开放，对修改关闭』的原则
  
 所以我们还有更加复杂的设计模式去适应更加复杂的系统~
 
 且听下回分解~ ~
 
 **此文章的源码都在这个仓库中： [golang设计模式](https://github.com/silsuer/golang-design-patterns)**
 
 > 打个广告，推荐一下自己写的 go web框架 [bingo](https://github.com/silsuer/bingo),求star，求PR ~
 