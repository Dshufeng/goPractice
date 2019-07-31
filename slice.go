package main

import "fmt"

func main() {
	c := []int{1, 2, 3}
	c = Append(c)
	Handle(c)
	fmt.Println("done: ", c)
}

func Handle(c []int) {
	fmt.Println("before: ", c)
	c[0] = 11
	fmt.Println("after: ", c)
}

func Append(c []int) []int {
	c = append(c, 4, 5)
	return c
}
