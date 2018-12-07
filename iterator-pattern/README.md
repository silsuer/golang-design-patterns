# 迭代器模式

## 定义

> wiki: 在 物件导向程式设计里，迭代器模式是一种设计模式，是一种最简单也最常见的设计模式。它可以让使用者透过特定的介面巡访容器中的每一个元素而不用了解底层的实作。

简单点说，为一个容器设置一个迭代函数，可以使用这个迭代函数来顺序访问其中的每一个元素，而外部无需知道底层实现。

如果再结合 [访问者模式](https://github.com/silsuer/golang-design-patterns/tree/master/visitor-pattern),向其中传入自定义的访问者，那么就可以让访问者访问容器中的每个元素了。

## 类图

![](https://images2015.cnblogs.com/blog/527668/201601/527668-20160109145125950-1322949510.jpg)

## 角色

 - 抽象聚合类: 定义一个抽象的容器
 
 - 具体聚合类: 实现上面的抽象类，作为一个容器，用来存放元素，等待迭代
 
 - 抽象迭代器: 迭代器接口，每个容器下都有一个该迭代器接口的具体实现
 
 - 具体迭代器: 根据不同的容器，需要定义不同的具体迭代器，定义了游标移动的具体实现
 
## 举个栗子

 1. 创建抽象容器结构体
   ```go
    // 容器接口
    type IAggregate interface {
    	Iterator() IIterator
    }
   ```
   
 2. 创建抽象迭代器
 
   ```go
    // 迭代器接口
    type IIterator interface {
	    HasNext() bool
	    Current() int
	    Next() bool
    }
   ```
   迭代器的基本需求，需要有判定是否迭代到最后的方法`HasNext()`,需要有获得当前元素的方法`Current()`,需要有将游标移动到下一个元素的方法 `Next()`
   
 3. 实现容器
   
   ```go 
     // 具体容器
     type Aggregate struct {
       container []int // 容器中装载 int 型容器
     }
     // 创建一个迭代器，并让迭代器中的容器指针指向当前对象
     func (a *Aggregate) Iterator() IIterator {
          i := new(Iterator)
          i.aggregate = a
          return i
      }
   ```
   为了简便，这里我们仅仅让容器中存放 `int`类型的数据
   
 4. 实现迭代器
 
   ```go
    type Iterator struct {
        cursor    int // 当前游标
        aggregate *Aggregate // 对应的容器指针
    }
    
    // 判断是否迭代到最后，如果没有，则返回true
    func (i *Iterator) HasNext() bool {
        if i.cursor+1 < len(i.aggregate.container) {
            return true
        }
        return false
    }
    
    // 获取当前迭代元素（从容器中取出当前游标对应的元素）
    func (i *Iterator) Current() int {
        return i.aggregate.container[i.cursor]
    }
    
    // 将游标指向下一个元素
    func (i *Iterator) Next() bool {
        if i.cursor < len(i.aggregate.container) {
            i.cursor++
            return true
        }
        return false
    }
   ```
 5. 使用迭代器
 
   ```go
    func main() {
	    // 创建容器，并放入初始化数据
	    c := &Aggregate{container: []int{1, 2, 3, 4}}
	    // 获取迭代器
	    iterator := c.Iterator() 
	    for {
	    	// 打印当前数据
		    fmt.Println(iterator.Current())
		    // 如果有下一个元素，则将游标移动到下一个元素
		    // 否则跳出循环，迭代结束
		    if iterator.HasNext() {
			    iterator.Next()
		    } else {
			    break
		    }
	    }
    }
   ```
   
   上面的例子比较简单，大家可以想象一下，如果容器类种的 `container` 不是一个切片或数组，而是一个结构体数组，我们却只需要遍历每个结构体中的某个字段呢？结构体的某个字段是数组，又或者，这个结构体的需要迭代的字段是私有的，在包外无法访问，所以这时是用迭代器，就可以自定义一个迭代函数，对外仅仅暴露几个方法，来获取数据，而外部不会在意容器内究竟是结构体还是数组。
   
   ## 实际栗子
   
   我可能又双叕要写自己造的轮子了...
   
   不过这里我只说说我对于迭代器模式的一个实际应用， **这个轮子没有开发完，只是个半半半成品**
   
   [github.com/silsuer/bingo-tpl](https://github.com/silsuer/bingo-tpl)
   
   这是我计划写的一个不借助 标准库中的`template` 实现的模板引擎，后来工作太忙，就搁浅了，涉及到了编译原理的一些知识，目前只写到了词法分析...
   
   可能要等 [Bingo](https://github.com/silsuer/bingo-tpl) 的所有模块都写完之后再去实现了吧...
   
   废话有点多... 入正题...
   
   模板引擎需要将模板中的字符，按照模板标记的左右定界符分割成词法链，例如下面的模板:
   
   ```
    {{ for item in navigation }}
        <li>tag</li>
    {{ endfor }}
   ```
   这个模板的意思是遍历 `navigation`，打印出对应数量的 `li` 标签
   
   将会生成如下的词法链
   
   ```
     [for]-> [item]-> [in]-> [navigation]-> [<li>tag</li>]-> [endfor] 
   ```
   每个方括号代表一个词法链上的节点，每个节点都会区分出是文本节点还是语法节点。
   
   具体的实现方法这里不说了，涉及到了词法分析器的状态转换，有兴趣的自己搜一搜就好，下面要实现的就是调试时打印词法链的过程，用到了迭代器模式。
   
   词法链的结构体如下:
   
   ```go
    // 词法分析链包含大量节点
    type LexicalChain struct {
    	Nodes       []*LexicalNode     
    	current     int                    // 当前指针
    	Params      map[string]interface{} // 变量名->变量值
    	TokenStream *TokenStream           // token流，这是通过节点解析出来的
    }
   ```
   对应的词法节点的结构体如下:
   
   ```go
    type LexicalNode struct {
    	T       int      // 类型（词法节点还是文本节点）
    	Content []byte   // 内容，生成模版的时候要使用内容进行输出
    	tokens  []*Token // token流
    	root    *Token   // 抽象语法树跟节点
    	lineNum int      // 行数
    	stack   []int    // 符栈，用来记录成对的操作符
    }
   ```
   
   每个节点的打印方法 `Print()`:
   
   ```go
    // 打印节点值
    func (n *LexicalNode) Print() {
    	switch n.T {
    	case textNode:
    		fmt.Println("[node type]: TEXT") // 文本节点
    	case lexicalNode:
    		fmt.Println("[node type]: LEXICAL") // 词法节点
    	default:
    		fmt.Println("[node type]: UNKNOWN TYPE") // 未知类型
    		break
    	}
    	fmt.Println("[line number]: " + strconv.Itoa(n.lineNum))
    	fmt.Println("[content]: " + string(n.Content))
    }
   ```
   上面是打印一个节点的，当要打印整个词法链时，需要迭代整个词法链，对每个节点调用 `Print()` 方法:
   
   ```go
     func (l *LexicalChain) Print() {
     	// 打印当前节点
     	l.Iterator(func(node *LexicalNode) {
     		fmt.Println("====================")
     		fmt.Println("[index]: " + strconv.Itoa(l.current))
     		node.Print()
     	})
     }    
   ```
   
   这里的实现方法与第一个栗子的实现方式不同，可以看做这里是迭代器模式与[访问者模式](https://github.com/silsuer/golang-design-patterns/tree/master/visitor-pattern)的结合使用，对迭代器方法`Iterator()`,传入一个回调函数作为访问者，
   
   对每个节点调用 `Print()` 方法来打印节点。
   
   下面看看 `Iterator()` 方法的实现:
   
   ```go
     func (l *LexicalChain) Iterator(call func(node *LexicalNode)) {
     	// 对于链中的每个节点，执行传入的方法
     	call(l.Current()) // 调用传入的方法，将当前节点作为参数传入
     	for {
     		if l.Next() != -1 {  // 这里和第一个栗子一样，将游标指向下一个元素，并继续调用 传入的回调
     			call(l.Current())
     		} else {
     			break  // 如果迭代到了最后，则直接跳出循环，结束迭代
     		}
     	}
     }
   ```
   
   `Next()` 函数的实现和第一个栗子差不多，就不贴了，代码在[这里](https://github.com/silsuer/bingo-tpl/blob/master/lexical_analysis.go)
   
   这样就可以对生成的词法链进行打印了，方便了后续调试开发...
   
## 总结

迭代器模式算是最简单也最常用的行为型模式了，在`JAVA`中尤其常见，对于`phper`如果习惯了 `laravel` 的[集合](https://laravel-china.org/docs/laravel/5.5/collections/1317),可以去看看其中`filter`/`map` 等方法的实现，都是通过传入一个回调，然后方法内部迭代元素的方式实现过滤的。

优点:
  1. 这个不用说了，设计模式的最大优点——解耦，增加容器或者迭代器都无需修改源代码，并且简化了容器类，将迭代逻辑抽出来，放在了迭代器中
  2. 可以用不同的方式来遍历同一个对象（这就是上面说的，通过传入不同的回调来进行不同的迭代）
缺点:
  1. 迭代器模式算是将一个对象（容器）的存储职责和遍历职责分离了，就像第一个栗子中说的，新增一个容器类就要新增一个迭代器类，增加了系统的代码量和复杂性。
      
> 上述代码均放在 [golang-design-patterns](https://github.com/silsuer/golang-design-patterns) 这个仓库中
      
> 打个广告，推荐一下自己写的 go web框架 [bingo](https://github.com/silsuer/bingo),求star，求PR ~
