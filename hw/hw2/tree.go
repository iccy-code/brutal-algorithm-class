/**
 * @File: homework_2.go
 *
 * @Author: iccy
 *
 * @Data:  22:48
 * @Software: GoLand
 *
 * : 第二次作业
 *
 */

 package main

 import (
	 "errors"
	 "fmt"
 )

 type binaryTree struct {
	 data  int
	 left  *binaryTree
	 right *binaryTree
 }
 
 func Init2tree(data int) *binaryTree {
	 return &binaryTree{data: data, left: nil, right: nil}
 }
 
 func (a *binaryTree) Append(data int) {
	 if data < a.data {
		 if nil == a.left {
			 a.left = &binaryTree{data: data, left: nil, right: nil}
			 return
		 }
		 a.left.Append(data)
	 } else {
		 if nil == a.right {
			 a.right = &binaryTree{data: data, left: nil, right: nil}
			 return
		 }
		 a.right.Append(data)
	 }
 }
 
 func (a *binaryTree) DFS() {
	 if nil == a {
		 return
	 }
	 a.left.DFS()
	 a.right.DFS()
	 fmt.Print(a.data, " ")
 }
 
 // 递归n次, 每次操作均为O(1), 所以时间复杂度为O(n), 空间复杂度为O(1)
 func (a *binaryTree) Traversal(str string, f func(tree *binaryTree) bool) {
	 ok := true
 
	 var traversal func(*binaryTree)
	 traversal = func(b *binaryTree) {
		 // pre
		 if nil == b || !ok {
			 return
		 }
		 if str == "pre" {
			 fmt.Print(b.data, " ")
		 }
 
		 // /in
		 traversal(b.left)
		 if !ok {
			 return
		 }
		 if str == "in" {
			 fmt.Print(b.data, " ")
		 }
 
		 // post
		 traversal(b.right)
		 if !ok {
			 return
		 }
		 if str == "post" {
			 fmt.Print(b.data, " ")
		 }
 
		 ok = f(b)
	 }
 
	 traversal(a)
 }
 
 func newTree(data int, nums []int) *binaryTree {
	 //elements := []int{4, 2, 5, 1, 3, 8, 7, 9}
 
	 b := Init2tree(data)
	 for i := 0; i < len(nums); i++ {
		 b.Append(nums[i])
	 }
 
	 return b
 }
 
 type iterator struct {
	 str      string
	 index    int
	 stack    []*binaryTree
	 tmpIndex int
	 tmp      []*binaryTree
	 ok       bool
 }
 
 func (a *binaryTree) InitIter(str string) *iterator {
	 ret := iterator{index: -1, str: str, stack: make([]*binaryTree, 10),
		 tmpIndex: 0, tmp: make([]*binaryTree, 20), ok: false}
	 ret.Push(a)
	 return &ret
 }
 
 func (a *iterator) Push(tree *binaryTree) {
	 a.index++
	 a.stack[a.index] = tree
 }
 
 func (a *iterator) Pop() (tree *binaryTree) {
	 tree = a.stack[a.index]
	 a.stack[a.index] = nil
	 a.index--
 
	 return tree
 }
 
 func (a *iterator) detection(tree *binaryTree) bool {
	 for i := 0; i < a.tmpIndex; i++ {
		 if tree == a.tmp[i] {
			 return false
		 }
	 }
	 return true
 }
 
 // 前序迭代是时间复杂度为O(1), 空间复杂度为O(1)
 // 中序迭代是时间复杂度为O(M)(M为最远的左子树长度), 空间复杂度为O(n), 需要临时空间
 // 后序迭代是时间复杂度为O(M)(M为最远的左子树长度), 空间复杂度为O(n), 需要临时空间
 func (a *iterator) Next() (ret int, err error) {
	 // 当前操作的节点可以看成是在一可最小树中, 该树最多三个节点
	 switch a.str {
	 case "pre":
		 {
			 // 没有节点可迭代
			 if -1 == a.index {
				 return 0, errors.New("traverse complete")
			 }
 
			 tmp := a.Pop()
			 if nil != tmp.right {
				 a.Push(tmp.right)
			 }
 
			 if nil != tmp.left {
				 a.Push(tmp.left)
			 }
 
			 return tmp.data, nil
		 }
	 case "in":
		 {
			 // 没有节点可迭代
			 if -1 == a.index && nil == a.tmp[a.tmpIndex-1].right.left &&
				 nil == a.tmp[a.tmpIndex-1].right.right {
				 return 0, errors.New("traverse complete")
			 }
 
			 // 中序遍历 --> 闭路循环
			 // 1. 若一个节点没有左右子节点, pop()
			 // 2. 循环(直到左子节点为nil): 若左子节点存在且a.tmp != 左子节点, push(左子节点), a.tmp = 左子节点
			 // 3. a.tmp = 根节点, pop() --> 此时是将根节点出栈(不是树的根节点)
			 // 4. 若a.tmp.right不存在, 不做任何操作; 若存在, push(右节点)
 
			 // 1
			 if nil == a.stack[a.index].left && nil == a.stack[a.index].right {
				 return a.Pop().data, nil
			 }
 
			 // 2
			 if nil != a.stack[a.index].left && a.detection(a.stack[a.index].left) {
				 // 顺着最左子树, 一路push()
				 tmp := a.stack[a.index].left
				 for tmp != nil {
					 a.Push(tmp)
					 tmp = tmp.left
				 }
			 }
 
			 // 3
			 a.tmp[a.tmpIndex] = a.stack[a.index]
			 a.tmpIndex++
			 ret = a.tmp[a.tmpIndex-1].data
			 a.Pop()
 
			 // 4
			 if nil != a.tmp[a.tmpIndex-1].right {
				 a.Push(a.tmp[a.tmpIndex-1].right)
			 }
 
			 return ret, nil
		 }
	 case "post":
		 {
			 if a.index == -1 {
				 return 0, errors.New("traverse complete")
			 }
 
			 tmp := a.stack[a.index]
		 LOOP:
			 for tmp.left != nil && a.detection(tmp.left) {
				 a.Push(tmp.left)
				 tmp = tmp.left
			 }
 
			 for tmp.right != nil && a.detection(tmp.right) {
				 a.Push(tmp.right)
				 tmp = tmp.right
				 if tmp.left != nil {
					 goto LOOP
				 }
			 }
 
			 tmp = a.Pop()
			 a.tmp[a.tmpIndex] = tmp
			 a.tmpIndex++
 
			 return tmp.data, nil
		 }
	 default:
		 fmt.Println("parameter error")
	 }
 
	 return 0, nil
 }
 
 func main() {
	orders := []string{"pre", "in", "post"}

	Trees := []*binaryTree{
		newTree(6, []int{4, 2, 5, 1, 3, 8, 7, 9}),
		newTree(7, []int{2, 1, 4, 3, 5, 6, 9, 8, 10}),
		newTree(1, []int{2, 3, 4, 5}),
		newTree(9, []int{10, 1, 2, 3, 4, 5}),
	}
	// 创建四棵树:
	// 0:								1:
	//           6						           7
	//         /   \					         /   \
	//        4     8					        2     9
	//      /  \   /  \					      /  \   / \
	//     2    5 7    9				     1    4 8    10
	//   /  \							         / \
	//  1    3							        3   5
	//									             \
	//									              6
	//
	// 前序遍历: 6 4 2 1 3 5 8 7 9		前序遍历: 7 2 1 4 3 5 6 9 8 10
	// 中序遍历: 1 2 3 4 5 6 7 8 9		中序遍历: 1 2 3 4 5 6 7 8 9 10
	// 后序遍历: 1 3 2 5 4 7 9 8 6		后序遍历: 1 3 6 5 4 2 8 10 9 7

	// 2:								3:
	//        1							           9
	//         \						         /  \
	//          2						        1    10
	//           \						         \
	//            3						          2
	//             \					           \
	//              4					            3
	//               \					             \
	//                5					              4
	//									               \
	//									                5
	// 退化成链表的树
	// 前序遍历: 1 2 3 4 5				前序遍历: 9 1 2 3 4 5 10
	// 中序遍历: 1 2 3 4 5				中序遍历: 1 2 3 4 5 9 10
	// 后序遍历: 5 4 3 2 1				后序遍历: 5 4 3 2 1 10 9

	// 第二题第一问
	exampleFunc := func(tree *binaryTree) bool { return true }

	for _, i := range Trees {
		for _, j := range orders {
			fmt.Printf("\n%-4s --> ", j)
			i.Traversal(j, exampleFunc)
		}
		fmt.Println()
	}

	fmt.Println()

	// 第二题第二问
	var iterators = make([]*iterator, len(Trees))

	for _, q := range orders {
		for w := range iterators {
			iterators[w] = Trees[w].InitIter(q)
		}
		fmt.Printf("%-4s\n", q)
		for _, e := range iterators {
			for r := 0; r < 12; r++ {
				m, err := e.Next()
				if err != nil {
					fmt.Print(err)
					break
				}
				fmt.Print(m, " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
 }