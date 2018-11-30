# 访问者模式

## 定义

> wiki: 表示一个作用于某对象结构中的各元素的操作。它使你可以在不改变各元素类的前提下定义作用于这些元素的新操作。

平时我们定义完一个类之后，这个类所能执行的逻辑就是确定的了，但是我们经常会遇到一种场景: 根据外部环境更改这个类所能执行的行为。

而 **访问者模式** 就是在不更改这个类的前提下，更改这个类中方法所能执行的逻辑。

## 角色组成

  - 抽象访问者
  
  - 访问者
  
  - 抽象元素类
  
  - 元素类
  
  - 结构容器: (非必须) 保存元素列表，可以放置访问者
## 类图

![](https://upload-images.jianshu.io/upload_images/2484780-ffe629e0fda0bab3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/840/format/webp)
(图源网络)

大概的流程就是:
 
  1. 从结构容器中取出元素
  
  2. 创建一个访问者
  
  3. 将访问者载入传入的元素（即让访问者访问元素）
  
  4. 获取输出

## 举个栗子

例如，一个对象的方法，在测试环境要打印出: **这是测试环境**,在生产环境中要打印出: **这是生产环境**

一般最开始会这样写:

  ```go
   type EnvExample struct {
   }
   
   func (e EnvExample) Print() {
   	if GetEnv() == "testing" {
   		fmt.Println("这是测试环境")
   	}
   
   	if GetEnv() == "production" {
   		fmt.Println("这是生产环境")
   	}
   }
  ```
  
这样这个`Print()` 方法的逻辑就耦合在当前结构体中了，扩展性差，现在假如我们要添加一个打印 **这是本地环境** 的逻辑呢？

我们就需要更改`Print()` 方法了，注意! 这是一个非常简单的例子，可以随意更改`Print()` 方法，没什么关系，但是在实际开发过程中

一个函数的实现是十分复杂的，有可能更改了这个方法，会导致整个系统崩溃，所以解耦是一个十分迫切的需要.

对象容器也只是一个容器而已，我就不实现了，用`Map`,`List`都能实现

1. 定义访问者接口

  ```go
   // 定义访问者接口
   type IVisitor interface {
   	Visit() // 访问者的访问方法
   }
  ```

2. 实现该接口

   ```go
     
    type ProductionVisitor struct {
    }
    
    func (v ProductionVisitor) Visit() {
        fmt.Println("这是生产环境")
    }
    
    type TestingVisitor struct {
    }
    
    func (t TestingVisitor) Visit() {
        fmt.Println("这是测试环境")
    }
   ```

3. 创建元素接口
  
   ```go
    // 定义元素接口
    type IElement interface {
    	Accept(visitor IVisitor)
    }
   ```
4. 实现元素接口

   ```go
    type Element struct {
    }
    
    func (el Element) Accept(visitor IVisitor) {
        visitor.Visit()
    }
   ```

5. 修改 `Print()` 方法
   
   ```go
    type EnvExample struct {
    	Element
    }
    
    func (e EnvExample) Print(visitor IVisitor) {
    	e.Element.Accept(visitor)
    }
   ```
6. 开始调用

  ```go
   // 创建一个元素
   e := new(Element)
   e.Accept(new(ProductionVisitor)) // output: 这是生产环境
   e.Accept(new(TestingVisitor))    // output: 这是测试环境
    
   m := new(EnvExample)
   m.Print(new(ProductionVisitor))
   m.Print(new(TestingVisitor))
  ```

> 自从各个语言开始支持匿名函数之后，访问者模式就变得极其简单了，每一种传入匿名方法的操作都可以看做是变相的访问者模式， golang 中的方法也是一种类型的对象，所以可以用它便利的实现访问者模式，下面看一个实际栗子

## 看一个实际应用的栗子

上面的例子不是很合适，也看不出来它的好处，这是因为变性太少，只有有限的几种 `生产环境`/`测试环境`/`本地环境`/`预发布环境`等等，现在我们看一个有很大变数的情况: `ORM`

有 `Laravel` 使用经验的同学，对**迁移**一定不陌生，快速的通过 `Blueprint`建表，因为不同的业务，使用的数据表结构一定不同，这样就有无数种变数了，看个栗子:

```php
  public function up()
  {
       Schema::create('authors', function(Blueprint $table)
      {
              $table->increments('id');
              $table->string('name');
              $table->timestamps();
        });
  }
```

所以这个 `Blueprint` 就相当于一个访问者，里面的各个方法就是可被元素类访问的方法，`Schema::create()`就是元素类了，不会`php`也没关系，我们看一个`go`实现

[github.com/silsuer/bingo-orm](https://github.com/silsuer/bingo-orm)

这是我自己写的框架下的一个子模块，专注数据库操作，具体安装过程就不说了，可以直接看`README`

下面是创建一个`test`表的代码:

```go
   err := c.Schema().CreateTable("test", func(table db.IBlueprint) {
   	  table.Increments("id").Comment("自增id")  // 设置备注与主键
      table.String("name").Comment("姓名")  
      table.Integer("age").Nullable().Comment("年龄") // 允许为空
     })
```

下面我们来看这个`CreateTable`源码:

```go
  func (ms *MysqlSchemaBuilder) CreateTable(tableName string, call func(table IBlueprint)) error {
  	// 创建一个 schema 对象，用来作为访问函数的参数
  	schema := new(MysqlBlueprint)
  	// 设置一些默认值
  	schema.engine = SchemaDefaultEngine
  	schema.name = tableName
  	schema.operator = AlterTable
  	
  	// 此处就是调用访问者了...
  	call(schema)
  
  	s := Assembly(CreateDefaultType, schema) // 拼装成语句
  	
  	// 执行sql语句
  	stmt, err := ms.GetConn().Prepare(s)
  	if err != nil {
  		return err
  	}
  	_, err = stmt.Exec()
  	return err
  }
```

这样就将创建表语句的方法解耦了，因为不同数据库的`sql`语句或特性等会有所不同，如果直接写死了的话，只能支持某一种数据库，那么我们就可以将组装`SQL`代码的逻辑抽象成 `IBlueprint`接口，实现各种方法，我们只需要为每一种数据库实现各自的 `IBlueprint`即可，

这对我们的开发包扩展提供了极大的方便。

> 如果想练手，可以 fork 一下这个项目，然后实现一下各种数据库的访问(我现在正在写mysql，并且还没有实现完全)，并提 PR 给我喔~

具体内容可以去看源码，这个包里的注释我写了不少...

大概就说这么多吧，针对这个 `ORM` 包，以后我还会出一个开发过程...先把这24种设计模式搞完...感觉给自己挖了个大坑...慢慢填吧...

> 上述代码均放在 [golang-design-patterns](https://github.com/silsuer/golang-design-patterns) 这个仓库中

> 打个广告，推荐一下自己写的 go web框架 [bingo](https://github.com/silsuer/bingo),求star，求PR ~




 