# 建造者模式

最近一直在写框架，这个系列也好久没更新了，抽时间更新一篇~

### 概述

> wiki: **建造者模式(Builder Pattern)**：将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。

直白一点的说，就是将我们在开发过程中遇到的大型对象，拆分成多个小对象，然后将多个小对象组装成大对象，并且对外部隐藏建造过程.

### 结构

建造者模式由一下4个部分组成

- `Builder`: 抽象建造者

- `ConcreteBuilder`: 具体建造者

- `Director`: 指挥者

- `Production`: 产品(大型产品以及拆分成的小型产品)

### 类图 && 时序图

![](https://design-patterns.readthedocs.io/zh_CN/latest/_images/Builder.jpg)
![](https://design-patterns.readthedocs.io/zh_CN/latest/_images/seq_Builder.jpg)
-----------
(*图源网络*)

从上面两张图可以看出建造者模式的使用流程:

  1. 创建大型产品建造者
  2. 创建指挥者
  3. 将建造者传入指挥者对象中
  4. 由指挥者指挥建造者创建对象，并返回

### 举个栗子

说一个网上说烂了的组装汽车的栗子吧,

比如说我是个老司机，但是除了开车还想造车，但是车的构造实在是太复杂了，那么我们就可以将车拆分...

4个轮子、1个底盘、1个驾驶位...

好了，为了简便，就造这三个吧，先造个爬犁出来...

所以我需要一个大型项目构造者`CarBuilder`:

```go
  
type CarBuilder struct {
	Car *Car
}

func (cb *CarBuilder) GetResult() interface{} {
	return cb.Car
}

func (cb *CarBuilder) NewProduct() {
	cb.Car = new(Car)
}

func (cb *CarBuilder) BuildWheels() {
	cb.Car.Wheels = "build wheels"
}

func (cb *CarBuilder) BuildChassis() {
	cb.Car.Chassis = "build chassis"
}

func (cb *CarBuilder) BuildSeat() {
	cb.Car.Seat = "build seat"
}
```

这个建造者实现了`Builder`接口:

```go
  type Builder interface {
  	  NewProduct()       // 创建一个空产品
	  BuildWheels()      // 建造轮子
	  BuildChassis()     // 建造底盘
	  BuildSeat()        // 建造驾驶位
	  GetResult() interface{}  // 获取建造好的产品
  }
```

下面要把具体建造者传入指挥者:

```go
type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}
```

现在指挥者和建造者都已经准备好了，可以进行建造了，调用指挥者的`Generate()` 方法:

```go
func (d *Director) Generate() *Car {
	d.builder.NewProduct()
	d.builder.BuildChassis()
	d.builder.BuildSeat()
	d.builder.BuildWheels()
	return d.builder.GetResult().(*Car)
}
```

这样，就得到了我们需要的 `Car` 对象：

```go
   func main() {
   	// 创建一个指挥者
   	director := new(Director)
   	// 创建建造者
   	builder := new(CarBuilder)
       director.SetBuilder(builder)
   	car := director.Generate()
   	car.Show()
   }
```

### 总结

  上面的代码，是将一个本来就不是很复杂的对象，强行拆分，只是将其中的字段设为最简单的`string` 类型，实际上，这些字段应该是更小的对象结构体，然后还可以继续把这些小结构体继续拆分，拆分为最小单元，这样才是结构最清晰的思路.
  
  本来想举一个应用在实际项目中的栗子的，但是框架还没有写完，这样，先占个坑，等`bingo`框架完成后我会再来补充这一部分.


> 上述代码均放在 [golang-design-patterns](https://github.com/silsuer/golang-design-patterns) 这个仓库中

> 打个广告，推荐一下自己写的 go web框架 [bingo](https://github.com/silsuer/bingo),求star，求PR ~

