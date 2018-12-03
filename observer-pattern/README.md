# 观察者模式

## 写在前面

#### 定义

> wiki: 观察者模式是软体设计模式的一种。在此种模式中，一个目标物件管理所有相依于它的观察者物件，并且在它本身的状态改变时主动发出通知。这通常透过呼叫各观察者所提供的方法来实现。此种模式通常被用来实时事件处理系统。

简单点说，可以想象成多个对象同时观察一个对象，当这个被观察的对象发生变化的时候，这些对象都会得到通知，可以做一些操作...

#### 多啰嗦几句

在`PHP`世界中，最出名的观察者模式应该就是 `Laravel` 的事件了，`Laravel`是一个事件驱动的框架，所有的操作都通过事件进行解耦，实现了一个简单的观察者模式，比较典型的一个使用就是[数据库模型](https://laravel-china.org/articles/6657/model-events-and-observer-in-laravel)，当观察到模型更改的时候，就会触发事件（`created`/`updated`/`deleted`...）

最开始用模型观察者的时候，只要在 `Observers` 目录中创建一个观察者对象，并且添加观察者关联，当修改模型的时候，就可以自动触发了,感觉好神奇喔... 

观察者模式在实际开发中经常用到，主要存在于底层框架中，与业务逻辑解耦,业务逻辑只需要实现各种观察者被观察者即可。

## 类图

![](https://gss1.bdstatic.com/-vo3dSag_xI4khGkpoWK1HF6hhy/baike/s%3D500/sign=7d0c44b1d52a283447a6360b6bb4c92e/3801213fb80e7bec49b2574d2e2eb9389b506b35.jpg)

(图源网络)

## 角色

 - 抽象观察者
 
 - 具体观察者
 
 - 抽象被观察者
 
 - 具体被观察者
 
 ## 举个栗子
 
  1. 创建抽象观察者
  
   ```
     // 抽象观察者
     type IObserver interface {
     	Notify() // 当被观察对象有更改的时候，出发观察者的Notify() 方法
     }
   ```
   
  2. 创建抽象被观察者
  
   ```go 
   // 抽象被观察者
   type ISubject interface {
	   AddObservers(observers ...IObserver) // 添加观察者
	   NotifyObservers()                    // 通知观察者
   }
   ```
 
 3. 实现观察者
   ```go
    
    type Observer struct {
    }

    func (o *Observer) Notify() {
	    fmt.Println("已经触发了观察者")
    }
   ```
  
 4. 实现被观察者
 
   ```go
    
    type Subject struct {
        observers []IObserver
    }
    
    func (s *Subject) AddObservers(observers ...IObserver) {
        s.observers = append(s.observers, observers...)
    }
    
    func (s *Subject) NotifyObservers() {
        for k := range s.observers {
            s.observers[k].Notify() // 触发观察者
        }
    }
   ```
  
 5. 使用实例
   ```go
        // 创建被观察者
         s := new(Subject)
         // 创建观察者
         o := new(Observer)
         // 为主题添加观察者
         s.AddObservers(o)
     
         // 这里的被观察者要做各种更改...
     
         // 更改完毕，触发观察者
         s.NotifyObservers()  // output: 已经触发了观察者
   ``` 

## 举个实际应用的例子

  对`PHP` 熟悉的同学可以看看这个,**Laravel中的[事件系统](https://github.com/kevinyan815/Learning_Laravel_Kernel/blob/master/articles/Event.md)和[观察者模式](https://github.com/kevinyan815/Learning_Laravel_Kernel/blob/master/articles/Observer.md)**
  
  下面写一个我自己项目中的栗子:  [github.com/silsuer/bingo-events](https://github.com/silsuer/bingo-events)
  
  这是一个`golang`语言实现的事件系统，我正在试着把它应用到自己的[框架](https://github.com/silsuer/bingo)中,它实现了两种观察者模式，一种是实现了观察者接口的观察者模式，一种是使用了反射进行类型映射的观察者模式，下面分别来说一下...
  
  1. 实现了观察者接口的方式:
  
   ```go
      // 创建一个结构体，实现事件接口
      type listen struct {
         bingo_events.Event
         Name string
      }
      
      func func main() {
      	// 事件对象 
      	app := bingo_events.NewApp() 
        l := new(listen)  // 创建被观察者
        l.Name = "silsuer"
        l.Attach(func (event interface{}, next func(event interface{})) { 
        	 // 由于监听器可监听的对象不一定非要实现 IEvent 接口，所以这里需要使用类型断言，将对象转换回原本的类型
             a := event.(*listen)
             fmt.Println(a.Name) // output: silsuer
             a.Name = "god"      // 更改结构体的属性
             next(a)             // 调用next函数继续走下一个监听器，如果此处不调用，程序将会终止在此处，不会继续往下执行
          })
        
        // 触发事件
        app.Dispatch(l)
      }
   ```
   这里我们使用结构体组合的形式实现了事件接口的实现，只要把`bingo-events.Event` 放入要监听的结构体中，就实现了`IEvent`接口，可以使用 `Attach()` 方法来添加观察者了，
   
   这里的观察者是一个 `func (event interface{}, next func(event interface{})) ` 类型的方法，
   
   第一个参数是触发的对象，毕竟观察者有时需要用到被观察者的属性，比如上面提到的`Laravel`的模型...
   
   第二个参数是一个 `func(event interface{})` 类型的方法，实际上这里实现了一个 `pipeline`，来做拦截功能，可以实现观察者的截断，只有在观察者的最后调用了 `next()` 方法，事件才会继续向下一个观察者传递，
   
   `pipeline` 的原理我在 [参考Laravel制作基于golang的路由包](https://juejin.im/post/5bd498cbf265da0ac669986f) 中写过，用来做中间件拦截的操作。
   
   
   2. 使用反射做观察者映射的方式
   
       ```go
         // 创建一个对象，不必实现事件接口
         type listen struct {
              	Name string      
         }

         func main() {
              	// 事件对象
              	app := bingo_events.NewApp()
              	// 添加观察者
              	app.Listen("*main.listen", ListenStruct)  // 直接使用 Listen 方法，为监听的结构体添加一个回调 
                l := new(listen)              
              	l.Name = "silsuer"  // 为对象属性赋值
          
              	// 复制完毕，开始分发事件，从上可知，共添加了两个观察者： ListenStruct 和 L2
                // 会按照监听的顺序执行，由于最开始已经添加了 ListenStruct 监听器，所以第二次再次添加的时候不会重复添加
                // 此处的分发，就是将参数顺序传入每一个监听器，进行后续操作
                app.Dispatch(l)
              }      
        func ListenStruct(event interface{}, next func(event interface{})) {
              	// 由于监听器可监听的对象不一定非要实现 IEvent 接口，所以这里需要使用类型断言，将对象转换回原本的类型       
        	    a := event.(*listen)
              	fmt.Println(a.Name) // output: silsuer
              	a.Name = "god"   // 更改结构体的属性
              	next(a)   // 调用next函数继续走下一个监听器，如果此处不调用，程序将会终止在此处，不会继续往下执行
              }
       ```
       
       这里我们没有使用实现了事件对象的 `Attach` 方法来添加观察者，而是使用一个字符串代表被观察者，这样就做到了无需实现观察者接口，就可以做到观察者模式。完整代码可以直接去看git仓库
       
       这里我们需要关注两个方法: `Listen()` 和 `Dispatch()`
       
       `bingo_events.App` 是一个服务容器，装载了所有事件和事件之间的映射关系
       
       ```go
        // 服务容器
        type App struct {
        	sync.RWMutex // 读写锁
        	events map[string][]Listener // 事件映射
        }
       ```
       
       下面看源代码:
       
       `Listen()`:
         ```go
           // 监听[事件][监听器]，单独绑定一个监听器
           func (app *App) Listen(str string, listener Listener) {
           	app.Lock()
           	app.events[str] = append(app.events[str], listener) 
           	app.Unlock()
           }
         ```      
         
         `app.events` 中保存的是通过字符串监听的对象，字符串就是通过`reflect.TypeOf(v).String()` 得到的字符串，例如上面的`listen`对象就是`*main.listen`，这是键，对应的值就是绑定的监听器方法
        
        `Dispatch()`
         
         ```go
          // 分发事件，传入各种事件，如果是
          func (app *App) Dispatch(events ...interface{}) {
          	// 容器分发数据
          	var event string
          	for k := range events {
          		if _, ok := events[k].(string); ok { // 如果传入的是字符串类型的
          			event = events[k].(string)
          		} else {
          			// 不是字符串类型的，那么得到其类型
          			event = reflect.TypeOf(events[k]).String()
          		}
          
          		// 如果实现了 事件接口 IEvent，则调用事件的观察者模式，得到所有的
          		var observers []Listener
          		if _, ok := events[k].(IEvent); ok {
          			obs := events[k].(IEvent).Observers()
          			observers = append(observers, obs...) // 将事件中自行添加的观察者，放在所有观察者之后
          		}
          
    		     // 如果存在map映射，则也放入观察者数组中
          		if obs, exist := app.events[event]; exist {
          			observers = append(observers, obs...)
          		}
          
          		if len(observers) > 0 {
          			// 得到了所有的观察者，这里通过pipeline来执行，通过next来控制什么时候调用这个观察者
          			new(Pipeline).Send(events[k]).Through(observers).Then(func(context interface{}) {
          
          			})
          		}
          
          	}
          
          }
        ``` 
        `Dispatch()` 方法传入一个对象（或对象类型的字符串），是否实现事件接口都可以，遍历所有对象（如果传入的是字符串，就不做处理，如果是对象，通过反射提取对象的字符串名称），然后从当前 `App` 中的`map`提取对应的观察者，如果对象还实现了事件接口，则再获取挂载在这个事件上的所有观察者，将他们装在同一个 `slice` 中，构建成一个 `pipeline`，顺序执行观察者。
    
> 上述代码均放在 [golang-design-patterns](https://github.com/silsuer/golang-design-patterns) 这个仓库中

> 打个广告，推荐一下自己写的 go web框架 [bingo](https://github.com/silsuer/bingo),求star，求PR ~
   