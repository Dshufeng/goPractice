package main

import (
	"fmt"
	"time"
)

//var wg sync.WaitGroup

func main() {
	timeStart := time.Now()

	//wg.Add(2)
	//// ch
	//ch := make(chan string)
	//go say("hello", ch)
	//go sing("world", ch)
	//
	//i := 1
	//for {
	//	fmt.Println(<-ch)
	//	if i >= 15 {
	//		close(ch)
	//		break
	//	}
	//	i++
	//}
	//wg.Wait()

	_say("hello")
	_sing("world")
	// time end
	timeEnd := time.Now()
	fmt.Println("time: ", timeEnd.Sub(timeStart))
}

func say(s string, c chan string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second * 1)
		c <- s
	}
	wg.Done()
}

func sing(s string, c chan string) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second * 1)
		c <- s
	}
	wg.Done()
}

func _say(s string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(s)
	}
}

func _sing(s string) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(s)
	}
}
