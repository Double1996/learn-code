package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("a", 1)
	fmt.Println(m.LoadOrStore("a", 1))
	m.Delete("a")
	fmt.Println(m.LoadOrStore("a", 1))
	fmt.Println(m.LoadOrStore("b", 2))

	// range
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return false
	})

}
