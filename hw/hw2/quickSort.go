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
	 "fmt"
	 "math/rand"
	 "time"
 )
 
 // 第一个版本
 func quickSort(a []int) {
	 if len(a) <= 1 {
		 return
	 }
	 left, right := 0, len(a) - 2
	 pivot := right+1
 
	 for left < right {
		 for a[left] <= a[pivot] && left < right {
			 left++
		 }
 
		 for a[right] >= a[pivot] && left < right {
			 right--
		 }
 
		 a[left], a[right] = a[right], a[left]
	 }
	 if a[left] >= a[pivot] {
		 a[left], a[pivot] = a[pivot], a[left]
	 }
 
	 quickSort(a[:left])
	 quickSort(a[left+1:])
 }
 
 // 所以O = O(logn)*O(n)*O(4)*O(m^2)
 //       = O(n*logn)*O(m^2)
 //       = 因为m最大为10, 所以O(m^2)可以近似的看成O(1)的操作
 //       = O(n*logn)
 func quickSort1(arr []int, start, end int) {
	 if end - start + 1 < 10 {
		 InsertSort(arr[start:end+1])				// O(m^2), m最大为10
		 return
	 }
	 if start < end {
		 i, j := start, end
		 key := arr[(start+end)/2]
 
		 if end - start >= 3 {						// O(4), 可以看成O(1)
			 if arr[start] > arr[end] {
				 if key > arr[start] {
					 key = arr[start]
				 } else if key < arr[end] {
					 key = arr[end]
				 }
			 } else {
				 if key < arr[start] {
					 key = arr[start]
				 } else if key > arr[end] {
					 key = arr[end]
				 }
			 }
		 }
 
		 //key := arr[start]
		 for i <= j {								// O(n)
			 for arr[i] < key {
				 i++
			 }
			 for arr[j] > key {
				 j--
			 }
			 if i <= j {
				 arr[i], arr[j] = arr[j], arr[i]
				 i++
				 j--
			 }
		 }
 
		 if start < j {								// O(logn)
			 quickSort1(arr, start, j)
		 }
		 if end > i {
			 quickSort1(arr, i, end)
		 }
	 }
 }
 
 func InsertSort(values []int) {
	 length := len(values)
	 if length <= 1 {
		 return
	 }
 
	 for i := 1; i < length; i++ {
		 tmp := values[i] // 从第二个数开始，从左向右依次取数
		 key := i - 1     // 下标从0开始，依次从左向右
 
		 // 每次取到的数都跟左侧的数做比较，如果左侧的数比取的数大，就将左侧的数右移一位，直至左侧没有数字比取的数大为止
		 for key >= 0 && tmp < values[key] {
			 values[key+1] = values[key]
			 key--
			 //fmt.Println(values)
		 }
 
		 // 将取到的数插入到不小于左侧数的位置
		 if key+1 != i {
			 values[key+1] = tmp
		 }
		 //fmt.Println(values)
	 }
 }
 
 func main() {
	 //第三题
	 // 创建四个数组, 分别是随机, 升序, 降序, 重复
	 MAX := 1000000
	 var e = make([][]int, 4)
	 for i, _ := range e {
		 e[i] = make([]int, MAX)
	 }
 
	 for i, _ := range e[0] {
		 e[0][i] = rand.Intn(MAX)
	 }
 
	 copy(e[1], e[0])
	 quickSort(e[1])
 
	 j := len(e[2])-1
	 for i := 0; i < MAX; i++ {
		 e[2][j] = e[1][i]
		 j--
	 }
 
	 for i, _ := range e[3] {
		 e[3][i] = rand.Intn(100)
	 }
 
	 /*
		 | 数组/排序方案 | 标准排序 | 三数取中值版本 | 三数+插排版本 |
		 | :----------: | :-----| | :-----------: | :----------: |
		 | 随机数组 | 92ms, 86ms  | 88ms, 92ms    | 24ms, 23ms   |
		 | 升序数组 | 2m, 2m8s    | 23ms, 25ms    | 10ms, 21ms   |
		 | 降序数组 | 3m20s, 3m13s| 25ms, 26ms    | 43ms, 43ms   |
		 | 重复数组 | 1.6s, 1.56s | 43ms, 44ms    | 25ms, 27ms   |
 
		 三数取中值有利于已经趋近有序的数据排序
		 在数据量很小的时候, 快排是慢于插排的, 所以在要排序的数组大小 < 10时使用插排
		 还有问题就是重复数据过多时, 不好排序, 没有优化
	 */
 
	 fmt.Println("标准版本")
	 for _, v := range e {
		 thin := time.Now()
		 quickSort(v)
		 fmt.Println(time.Since(thin))
	 }
 
	 fmt.Println("三数取中版本")
	 for _, v := range e {
		 thin := time.Now()
		 quickSort1(v, 0, len(v)-1)
		 fmt.Println(time.Since(thin))
	 }
 }