package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println("--- 1 ----")
	//
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Printf("panic: %s\n", r)
	//	}
	//	fmt.Println("--- 2 ----")
	//}()
	////panic("i am panic")
	//
	//var slice = []int{1, 2, 3, 4, 5}
	//slice[6] = 11
	var a string
	fmt.Scanln(&a)
	i, _ := strconv.Atoi(a)
	fmt.Printf("%T", i)
	//var i interface{} = a
	assert(i)
}
func assert(i interface{}) {
	s, ok := i.(int)
	if ok {
		fmt.Println("is int: ", s, ok)
	} else {
		fmt.Println("not int: ", s, ok)
	}
}
