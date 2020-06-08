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
 
 // 递归链表
 type node struct {
	 data int
	 next *node
 }
 
 func InitList(data int) *node {
	 return &node{data: data, next: nil}
 }
 
 func (a *node) GetLength() int {
	 length := 1
 
	 var getLength func(*node)
	 getLength = func(b *node) {
		 if nil == b.next {
			 return
		 }
		 length++
		 getLength(b.next)
	 }
 
	 getLength(a)
	 return length
 }
 
 func (a *node) ListAppend(data int) {
	 if nil == a.next {
		 a.next = &node{data: data, next: nil}
		 return
	 }
 
	 a.next.ListAppend(data)
 }
 
 func (a *node) ListAppendHead(data int) {
	 a.next = &node{data: a.data, next: a.next}
	 a.data = data
 }
 
 func (a *node) ListDisplay() {
	 fmt.Printf("%d -> ", a.data)
 
	 if nil != a.next {
		 a.next.ListDisplay()
	 }
 }
 
 func (a *node) ListInsert(index, data int) error {
	 if 0 > index || index >= a.GetLength() {
		 return errors.New("the subscript error")
	 }
 
	 var listInsert func(*node)
	 listInsert = func(b *node) {
		 if 0 == index {
			 b.ListAppendHead(data)
			 return
		 }
		 index--
		 listInsert(b.next)
	 }
	 listInsert(a)
 
	 return nil
 }
 
 func (a *node) GetData(index int) (int, error) {
	 if 0 > index || index >= a.GetLength() {
		 return 0, errors.New("the subscript error")
	 }
 
	 var listInsert func(*node) int
	 listInsert = func(b *node) int {
		 if 0 == index {
			 return b.data
		 }
		 index--
		 return listInsert(b.next)
	 }
 
	 return listInsert(a), nil
 }
 
 func newList(num int) *node {
	 a := InitList(1)
	 for i := 0; i < num; i++ {
		 a.ListAppend(i + 2)
	 }
 
	 return a
 }

 // 后序遍历链表的时间复杂度为O(n), 空间复杂度为O(n*sizeof(data))
 func (a *node) ListLoopDisplay() {
	tmpPtr := a
	elements := make([]int, a.GetLength()) // O(n*sizeof(data))

	for i := 0; i < len(elements); i++ { // O(n)
		elements[i] = tmpPtr.data
		tmpPtr = tmpPtr.next
	}

	for i := len(elements) - 1; i >= 0; i-- { // O(n)
		fmt.Printf("%d -> ", elements[i])
	}
 }
 func main() {
	 // 第一题
	a := newList(10)
	a.ListLoopDisplay()
 }