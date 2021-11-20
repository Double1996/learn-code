package main

import "fmt"

/**
排序思想：
时间复杂度：
空间复杂度:
*/

func heapSort(arr []int) []int {
	arrlen := len(arr)
	buildMaxHeap(arr, arrlen)
}

func buildMaxHeap(arr []int, arrlen int) {

}

func heapify(arr []int, i, arrlen int) {
	left := 2*i - 1
	right := 2*i + 2
}

func main() {
	fmt.Println(heapSort([]int{120, 13, 23, 516, 99, 87, 23, 7, 103}))