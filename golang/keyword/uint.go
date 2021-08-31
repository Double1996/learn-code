package main

import "fmt"

func main() {
	var a uint32 = 2
	var b uint32 = 3
	v := int(a) - int(b)
	fmt.Println(v)
	fmt.Println(uint(-1))
}
