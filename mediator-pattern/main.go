package main

import "fmt"

// 根据二叉树的前序遍历和中序遍历的结果，重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

/*
preorder = [3,9,20,15,7]
inorder =  [9,3,15,20,7]
*/

var indexForInorders map[int]int

func init() {
	indexForInorders = make(map[int]int)
}

type Node struct {
	value int   // 值
	left  *Node // 左节点
	right *Node // 右节点
}

// 之前的操作,传入前序遍历和中序遍历
func preBinaryTree(pre []int, in []int) *Node {
	// 将中序遍历的值和索引倒过来
	for k, v := range in {
		indexForInorders[v] = k
	}
	return inBinaryTree(pre, 0, len(pre)-1, 0)
}

// 前序遍历序列  序列左，序列右，中序遍历左定界package main
//
//import "fmt"
//
//// 根据二叉树的前序遍历和中序遍历的结果，重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
///*
//preorder = [3,9,20,15,7]
//inorder =  [9,3,15,20,7]
//*/
//
//var indexForInorders map[int]int
//
//func init() {
//	indexForInorders = make(map[int]int)
//}
//
//type Node struct {
//	value int   // 值
//	left  *Node // 左节点
//	right *Node // 右节点
//}
//
//// 之前的操作,传入前序遍历和中序遍历
//func preBinaryTree(pre []int, in []int) *Node {
//	// 将中序遍历的值和索引倒过来
//	for k, v := range in {
//		indexForInorders[v] = k
//	}
//	return inBinaryTree(pre, 0, len(pre)-1, 0)
//}
//
//// 前序遍历序列  序列左，序列右，中序遍历左定界
//func inBinaryTree(pre []int, l int, r int, inL int) *Node {
//	if l > r {
//		return nil
//	}
//
//	// 根节点(第一个元素是根节点)
//	root := &Node{value: pre[l]}
//
//	// 左限制就是前置序列的下一个元素
//	// 右限制就是左限制加上到中间元素
//	i := indexForInorders[pre[l]] // 中序遍历的根节点的索引
//	ii := i - inL
//	root.left = inBinaryTree(pre, l+1, l+ii, inL)
//	root.right = inBinaryTree(pre, l+ii+1, r, inL+ii+1)
//	return root
//}
//
//// 打印node
//func (n *Node) print() {
//	if n.left != nil {
//		n.left.print()
//	}
//	fmt.Print(n.value, " ")
//	if n.right != nil {
//		n.right.print()
//	}
//}
//
//func main() {
//	pre := []int{3, 9, 20, 15, 7}
//	in := []int{9, 3, 15, 20, 7}
//	r := preBinaryTree(pre, in)
//	r.print()
//}
func inBinaryTree(pre []int, l int, r int, inL int) *Node {
	if l > r {
		return nil
	}

	// 根节点(第一个元素是根节点)
	root := &Node{value: pre[l]}

	// 左限制就是前置序列的下一个元素
	// 右限制就是左限制加上到中间元素
	i := indexForInorders[pre[l]] // 中序遍历的根节点的索引
	ii := i - inL
	root.left = inBinaryTree(pre, l+1, l+ii, inL)
	root.right = inBinaryTree(pre, l+ii+1, r, inL+ii+1)
	return root
}

// 打印node
func (n *Node) print() {
	if n.left != nil {
		n.left.print()
	}
	fmt.Print(n.value, " ")
	if n.right != nil {
		n.right.print()
	}
}

func main() {
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}
	r := preBinaryTree(pre, in)
	r.print()
}
