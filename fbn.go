package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	//n := fbn(45）
	n := fibonacci(45)
	fmt.Println(n)
	t2 := time.Now()
	fmt.Println("time=", t2.Sub(t1))
}

func fbn(i int) int {
	if i == 1 || i == 2 {
		return i
	} else {
		return fbn(i-1) + fbn(i-2)
	}
}

var fibarry = [3]int{0, 1, 0}

func fibonacci(n int) int {
	for i := 2; i <= n; i++ {
		fibarry[2] = fibarry[0] + fibarry[1]
		fibarry[0] = fibarry[1]
		fibarry[1] = fibarry[2]
	}
	return fibarry[2]
}
