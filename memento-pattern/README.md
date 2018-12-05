# 备忘录模式

## 定义

>wiki: 备忘录模式是一种软件设计模式：在不破坏封闭的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。这样以后就可将该对象恢复到原先保存的状态。

这也是一种比较常见的设计模式，可以用来创建程序某个时刻运行状态的快照，当程序异常崩溃或者因为其他原因导致退出后，可以使用备忘后的数据，恢复到原始状态，最常见的操作应该就是编辑器的撤销了，编辑器应用了备忘录模式，将编辑过程中的代码状态放在一个状态栈中，当使用`ctrl+z` 的时候，就从栈中弹出上一次保存的状态，来恢复到上一次的情况（即撤销）。

## 角色

  - 发起人: 发起人的内部要规定要备忘的范围，负责提供备案数据
  
  - 备忘录: 存储发起人对象的内部状态，在需要的时候，可以向其他人提供这个内部状态，以方便负责人恢复发起人状态
  
  - 负责人: 负责对备忘录进行管理（保存或提供备忘录）


## 类图
![](http://my.csdn.net/uploads/201206/27/1340804147_2145.jpg)
(图源网络)

看起来很简单，当需要保存状态的时候，负责人从发起人处拿到状态，然后保存到备忘录中，

当需要恢复的时候，负责人从备忘录中拿出上一个状态，然后恢复到发起人中

## 举个简单的栗子

1. 备忘录结构体

   ```go
    // 备忘录
    type Memento struct {
    	state string // 这里就是保存的状态
    }
    
    func (m *Memento) SetState(s string) {
    	m.state = s
    }
    
    func (m *Memento) GetState() string {
    	return m.state
    }
   ```
   
   备忘录结构体中的`state`只是一个单一的状态，我们用一个最简单的字符串形式来表示，在实际开发过程中，根据需要变为对应的数据结构，
   
   如果是多状态的，比如编辑器的撤销操作，就保存一个栈。
   
2. 发起人结构体
   
   ```go
     // 发起人
     type Originator struct {
     	state string // 这里就简单一点，要保存的状态就是一个字符串	
     }
     
     func (o *Originator) SetState(s string) {
     	o.state = s
     }
     
     func (o *Originator) GetState() string {
     	return o.state
     }
      
     // 这里就是规定了要保存的状态范围
     func (o *Originator) CreateMemento() *Memento {
     	return &Memento{state: o.state}
     }
   ```
  
3. 负责人结构体

   ```go
    // 负责人
    type Caretaker struct {
    	memento *Memento
    }
    
    func (c *Caretaker) GetMemento() *Memento {
    	return c.memento
    }
    
    func (c *Caretaker) SetMemento(m *Memento) {
    	c.memento = m
    }
   ```

4. 搞一下...

   ```go
     func main() {
     	// 创建一个发起人并设置初始状态
     	// 此时与备忘录模式无关，只是模拟正常程序运行
     	o := &Originator{state: "hello"}
         fmt.Println("当前状态:",o.GetState())
     	// 现在需要保存当前状态
     	// 就创建一个负责人来设置（一般来说，对于一个对象的同一个备忘范围，应当只有一个负责人，这样方便做多状态多备忘管理）
     	c := new(Caretaker)
     	c.SetMemento(o.CreateMemento())
     
     	o.SetState("world")
     	fmt.Println("更改当前状态:",o.GetState())
     
     	// 恢复备忘
     	o.SetState(c.GetMemento().GetState())
     	fmt.Println("恢复后状态",o.GetState())
     }
   ```

## 小结
   
上面就是一个单一状态的备忘录模式的完整流程了，还可以做很多操作，比如把保存的备忘录序列化成字符串保存在磁盘中，下次启动的时候从磁盘中获取状态，这样就可以做一个简单的快照了。

当然还有多状态的，一般来说，一个对象有一个对应的备忘对象，记录对象中要备忘的字段，而多个对象的备忘，同一由同一个负责人进行管理，可以用 `map` 来做到这一点，

负责人中的备忘容器是一个`map`类型的数据，值是一个实现备忘接口的数据结构即可。

另外要注意一下备忘模式的优缺点:

  - 优点: 
    
     - 备忘录模式仅做数据备忘，不论该数据是否正确。
     - 设计模式最大的优点就是解耦，各司其职，发起人只需要提供备忘数据，不需要对其进行管理
  
  
  - 缺点

     - 实际应用中，备忘录模式大多是多状态的，如果进行大量备忘的话，会占用大量内存，当然，如果持久化在磁盘中的话，会减少内存占用，但会增加IO操作，这就需要开发者根据实际业务情况进行取舍了。


> 上述代码均放在 [golang-design-patterns](https://github.com/silsuer/golang-design-patterns) 这个仓库中

> 打个广告，推荐一下自己写的 go web框架 [bingo](https://github.com/silsuer/bingo),求star，求PR ~