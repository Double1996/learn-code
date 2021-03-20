package main

/**
  排序思想；
  时间复杂读:  最糟糕的情况是 O(n2) 平均是 O(log n)
  空间复杂度:  O(n) 平均是 O(log n)
*/

func quickSort(array []int, left, right int) {
	if len(array) == 0 {
		return
	}

	piovt := array[len(array)-1] // 基准点

}

func partition(array []int, left, right int) {

}

func main() {
	var nums = []int{120, 13, 23, 516, 99, 87, 23, 7, 103}
	quickSort(nums)

}
