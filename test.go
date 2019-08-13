package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var myMap map[int]string
	//myMap = make(map[int]string, 2)
	//myMap[0] = "apple"
	//myMap[1] = "apple1"
	//myMap[2] = "apple2"
	//myMap[3] = "apple3"
	//fmt.Println(myMap)

	fz()
}

func fz() {
	type MyStruct struct {
		N []string
	}

	n := MyStruct{
		[]string{
			"111", "2222",
		},
	}
	fmt.Println(n)

	test := []string{
		"aa", "bb",
	}
	mutable := reflect.ValueOf(&n).Elem()
	mutable.FieldByName("N").Set(reflect.ValueOf())
	fmt.Println(n)
}
