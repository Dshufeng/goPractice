package main

import "fmt"

func main() {
	var myMap map[int]string
	myMap = make(map[int]string, 2)
	myMap[0] = "apple"
	myMap[1] = "apple1"
	myMap[2] = "apple2"
	myMap[3] = "apple3"
	fmt.Println(myMap)
}
