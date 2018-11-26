# 原型模式

> wiki:原型模式是创建型模式的一种，其特点在于通过“复制”一个已经存在的实例来返回新的实例,而不是新建实例。被复制的实例就是我们所称的“原型”，这个原型是可定制的。原型模式多用于创建复杂的或者耗时的实例，因为这种情况下，复制一个已经存在的实例使程序运行更高效；或者创建值相等，只是命名不一样的同类数据。

这是一个十分简单的设计模式,可以看做是其他语言中的克隆方法，例如 `JAVA`/`PHP` 中都有相关方法，从一个内存中已经存在的对象中，拷贝出一个一模一样的对象来，针对复杂对象或比较大的对象，要比使用各种设计模式`new`出来的对象要快的多,

而且原型模式很少单独使用，一般与其他对象结合使用。

### 栗子

1. 创建一个结构体

   ```go
     // 示例结构体
     type Example struct {
     	Content string
     }
   ```
   
2. 添加克隆方法

   ```go
     func (e *Example) Clone() *Example {
     	res := *e
     	return &res
     }
   ```
   我们仅仅一行代码就完成了值的拷贝，使用 `*指针`，直接获取了一个拷贝的值，然后将这个拷贝的值得指针返回，原理请阅读下面的扩展阅读。
     
3. 编写主代码

  ```go
     func main() {
     	r1 := new(Example)
     	r1.Content = "this is example 1"
     	r2 := r1.Clone()
     	r2.Content = "this is example 2"
     
     	fmt.Println(r1)
     	fmt.Println(r2)
     
     }
  ```



### 扩展阅读: 深拷贝与浅拷贝

  `go` 语言中的传递都是值传递，传递一个对象，就会把对象拷贝一份传入函数中，传递一个指针，就会把指针拷贝一份传入进去。
  
  赋值的时候也是这样，`res:=*e` 就会把传递的 `Example` 对象拷贝一份，如果是 `res:=e` 的话，那么拷贝的就是对象的指针了.
  
  而深拷贝和浅拷贝也可以这样理解，深拷贝就是拷贝整个对象，浅拷贝就是拷贝对象指针。
  
  对于深度拷贝，`go`和其他语言还经常使用序列化后反序列化的形式进行拷贝:
  
  ```go
   func deepCopy(dst, src interface{}) error {
       var buf bytes.Buffer
       if err := gob.NewEncoder(&buf).Encode(src); err != nil {
           return err
       }
       return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
    }
  ```
  
  实际上`gob`包序列化的时候也是用到了 `reflect`包来实现拷贝的
  
  **注意:** golang完全是按值传递，所以如果深度拷贝的对象中包含有指针的话，那么深度拷贝后，这些指针也会相同，会导致部分数据共享，要注意这一点.
  
至此，所有创建型的设计模式就已经全部写完了，可以去下面的仓库中找到其他的设计模式喔....

> 上述代码均放在 [golang-design-patterns](https://github.com/silsuer/golang-design-patterns) 这个仓库中

> 打个广告，推荐一下自己写的 go web框架 [bingo](https://github.com/silsuer/bingo),求star，求PR ~