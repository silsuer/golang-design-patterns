# 模板方法模式

## 定义

> wiki: 模板方法模式定义了一个演算法的步骤，并允许子类别为一个或多个步骤提供其实践方式。让子类别在不改变演算法架构的情况下，重新定义演算法中的某些步骤。在软件工程中，它是一种软件设计模式，和C++模板没有关连。

简单一点说模板方法模式就是将一个类种能够公共使用的方法放置在抽象类种实现，不能公共使用的方法作为抽象方法，强制子类去实现，这样就做到了将一个类作为一个模板，让开发者去填充需要去填充的地方。



## 角色

在`go`语言中，不存在抽象类的概念，无法强制子类实现某些方法，但是我们可以使用结构体组合来实现模板方法模式。下面的角色分工只适用于`go`语言。
  
  - 基类结构体: 相当于其他 `OO` 语言的抽象类，实现顶级行为(模板方法所代表的行为称为顶级行为，其逻辑称为顶级逻辑)
  
  - 实现结构体： 相当于其他 `OO` 语言的子类，重写顶级行为或补充顶级行为
  
## 举个栗子

  1. 创建一个基类结构体
   
   ```    
    type Base struct {
    
    }
    
    func (b Base) Print()  {
        fmt.Println("这是print方法")
    }
    
    func (b Base) Echo()  {
    
    }
   ```
   
   我们实现了这个基类结构体得`Print()` 方法，而定义了一个空的 `Echo()` 方法留待子类实现，下面我们开始组合结构体
   
   2. 组合结构体
   
   ```go
     type Son struct {
     	Base
     }
   ```
   将 `Base` 结构体内置到 `Son`结构体中，实现结构体组合的功能
   
   3. 实现 `Echo` 方法
   
   ```go
      func (s Son) Echo()  {
      	fmt.Println("这是Echo方法")
      }
   ```
   
   4. 试验一下
   
   ```go
     func main() {
     	s := new(Son)
     	s.Print()  // output: 这是Print方法
     	s.Echo()   // output: 这是Echo方法
     }
   ```
   
   `Print()` 方法调用了父类的`Print()` 方法，`Echo()` 方法由于被子类重写，则调用子类的`Echo()`方法。
   
## 实际应用

  `golang`的结构体组合实在是太常见了... 
  
  屁如说... `Beego`的控制器:
  
  ```go
     type ExampleController struct {
         beego.Controller
     }
  ```
  
  `Beego` 控制器中都要内嵌一个 `beego.Controller`来继承一些方法.
   
   还有很多很多栗子...
   
**结构体组合有一些要注意的地方:**

`golang` 提倡的是把所有逻辑都显式操作，所以没有继承、抽象等等东西，虽然可以使用结构体组合达到一些重写的效果，但是还是有一点区别的:

1. 子类调用父类的方法时，父类方法获得的属性也是父类属性，无法获取子类属性

   ```go
   
      type Base struct {
   	    column string
      }
   
      type Son struct{
   	    Base
	    column string
      }
   
      func (b Base) Print()  {
             fmt.Println(b.column)
      }
   ```
  这时在 `Base` 类中是拿不到子类的字段的.
  
2. 如果结构体中藏一个指针的话，是无法做到组合的，因为指针不会自动初始化:

  ```go
    type Example struct {
    	*Base
    }
  ```
  此时执行:
  
  ```go
   e := new(Example)
   e.Print()
  ```
  会报空指针错误.
  
3. 组合时不论父类的模板方法是值调用还是指针调用，只要子类的方法中有和父类同名的方法，就会以子类方法为准（不论子类是值调用还是指针调用）

## 总结一下

其实`golang`的结构体组合就是一种模板方法模式的实现，`golang`不像`JAVA`等其他语言，缺失了一些特性，无法做到强制子类实现某些方法，但是上面的那种形式，确实符合模板方法模式的定义，规定了逻辑流程，那么我们就可以说是实现了模板方法模式，设计模式更多的是思想，在不同语言中都有不同的体现。

> 上述代码均放在 [golang-design-patterns](https://github.com/silsuer/golang-design-patterns) 这个仓库中

> 打个广告，推荐一下自己写的 go web框架 [bingo](https://github.com/silsuer/bingo),求star，求PR ~

