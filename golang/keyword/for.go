package main

import "fmt"

func for1() {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr)
}

func for2() {
	arr := []int{1, 2, 3}
	var newArr []*int
	for _, v := range arr {
		newArr = append(newArr, &v) //
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}

func for3() {
	arr := []int{1, 2, 3}
	var newArr []*int
	for index, _ := range arr {
		newArr = append(newArr, &arr[index]) // 需要使用 &arr[index]
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}

func main() {
	for2()
}
