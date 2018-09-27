# 单例模式

> wiki百科: 单例模式，也叫单子模式，是一种常用的软件设计模式。在应用这个模式时，单例对象的类必须保证只有一个实例存在。许多时候整个系统只需要拥有一个的全局对象，这样有利于我们协调系统整体的行为。比如在某个服务器程序中，该服务器的配置信息存放在一个文件中，这些配置数据由一个单例对象统一读取，然后服务进程中的其他对象再通过这个单例对象获取这些配置信息。这种方式简化了在复杂环境下的配置管理。


单例模式要实现的效果就是，对于应用单例模式的类，整个程序中只存在一个实例化对象

go并不是一种面向对象的语言，所以我们使用结构体来替代

有几种方式:

  - 懒汉模式
  
  - 饿汉模式
  
  - 双重检查锁机制

下面拆分讲解：

### 懒汉模式

1. 构建一个示例结构体

  ```go
     type example struct {
     	name string
     }
  ```
2. 设置一个私有变量作为每次要返回的单例

  ```go
    var instance *example
  ```
  
3. 写一个可以获取单例的方法

  ```go
      func GetExample() *example {
      
      	// 存在线程安全问题，高并发时有可能创建多个对象
      	if instance == nil {
      		instance = new(example)
      	}
      	return instance
      }
  ```
  
4. 测试一下

   ```go
     func main() {
     	s := GetExample()
     	s.name = "第一次赋值单例模式"
     	fmt.Println(s.name)
     
     	s2 := GetExample()
     	fmt.Println(s2.name)
     }
   ```
   
懒汉模式存在线程安全问题，在第3步的时候，如果有多个线程同时调用了这个方法，
那么都会检测到`instance`为`nil`,就会创建多个对象，所以出现了饿汉模式...


### 饿汉模式

与懒汉模式类似，不再多说，直接上代码

```go

  // 构建一个结构体，用来实例化单例
  type example2 struct {
  	name string
  }
  
  // 声明一个私有变量，作为单例
  var instance2 *example2
  
  // init函数将在包初始化时执行，实例化单例
  func init() {
  	instance2 = new(example2)
  	instance2.name = "初始化单例模式"
  }
  
  func GetInstance2() *example2 {
  	return instance2
  }
  
  func main() {
  	s := GetInstance2()
  	fmt.Println(s.name)
  }

``` 

饿汉模式将在包加载的时候就创建单例对象，当程序中用不到该对象时，浪费了一部分空间

和懒汉模式相比，更安全，但是会减慢程序启动速度


### 双重检查机制

懒汉模式存在线程安全问题，一般我们使用互斥锁来解决有可能出现的数据不一致问题 

所以修改上面的`GetInstance()` 方法如下:

```go
   var mux Sync.Mutex
   func GetInstance() *example {
       mux.Lock()                    
       defer mux.Unlock()
       if instance == nil {
           instance = &example{}
       }
      return instance
   }
```

如果这样去做，每一次请求单例的时候，都会加锁和减锁，而锁的用处只在于解决对象初始化的时候可能出现的并发问题
当对象被创建之后，加锁就失去了意义，会拖慢速度，所以我们就引入了双重检查机制（`Check-lock-Check`）,
也叫`DCL`(`Double Check Lock`), 代码如下:

```go
  func GetInstance() *example {
      if instance == nil {  // 单例没被实例化，才会加锁 
          mux.Lock()
          defer mux.Unlock()
          if instance == nil {  // 单例没被实例化才会创建
  	            instance = &example{}
          }
      }
      return instance
  }
```

这样只有当对象未初始化的时候，才会又加锁和减锁的操作

但是又出现了另一个问题：每一次访问都要检查两次，为了解决这个问题，我们可以使用golang标准包中的方法进行原子性操作:

```go
   import "sync"  
   import "sync/atomic"
   
   var initialized uint32
   
   func GetInstance() *example {
   	
   	  // 一次判断即可返回
      if atomic.LoadUInt32(&initialized) == 1 {
   		return instance
   	   }
       mux.Lock()
       defer mux.Unlock()
       if initialized == 0 {
            instance = &example{}
            atomic.StoreUint32(&initialized, 1) // 原子装载
   	}
   	return instance
   }
```
以上代码只需要经过一次判断即可返回单例，但是golang标准包中其实给我们提供了相关的方法:

`sync.Once`的`Do`方法可以实现在程序运行过程中只运行一次其中的回调，所以最终简化的代码如下:

```go

 type example3 struct {
 	name string
 }
 
 var instance3 *example3
 var once sync.Once
 
 func GetInstance3() *example3 {
 
 	once.Do(func() {
 		instance3 = new(example3)
 		instance3.name = "第一次赋值单例"
 	})
 	return instance3
 }
 
 func main() {
 	e1 := GetInstance3()
 	fmt.Println(e1.name)
 
 	e2 := GetInstance3()
 	fmt.Println(e2.name)
 }
```

单例模式是开发中经常用到的设计模式，我在制作自己的web框架 [silsuer/bingo](https://github.com/silsuer/bingo) 的时候
在环境变量控制、配置项控制等位置都用到了这种模式。

想把所有设计模式使用golang实现一遍，开了个新坑[silsuer/golang-design-patterns](https://github.com/silsuer/golang-design-patterns),
这是第一篇，以后会陆续更新，需要请自取～