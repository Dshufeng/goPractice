package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// str
	str := "hello world"
	fmt.Println(str)

	// arr
	arr := [...]int{1, 2, 3, 4}
	fmt.Println(arr)

	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)

	// slice
	slice := []int{1, 2, 3, 4}
	fmt.Println(slice[1:])

	// map
	map1 := map[int]string{1: "a"}
	map1[2] = "b"
	fmt.Println(map1)

	for k, v := range map1 {
		fmt.Printf("k=%d v=%v\n", k, v)
	}

	map2 := map[string]string{}

	map2["aaa"] = "bb"
	map2["bbb"] = "cc"

	for _, v := range map2 {
		if v == "bb" {
			fmt.Printf("v=%v\n", v)
		}
	}

	// map interface
	map3 := make(map[string]interface{})
	map4 := make(map[string]interface{})
	map3["name"] = "dsf"
	map3["age"] = 19

	map4["map3"] = map3
	map4["sex"] = "male"

	fmt.Println(map4)

	fmt.Println("start")
	for _, v := range map4 {
		if vv, ok := v.(map[string]interface{}); ok {
			fmt.Println(vv)
			for _, v1 := range vv {
				fmt.Println(v1)
			}
		} else {
			fmt.Println(v)
		}
	}
	map4Json, _ := json.Marshal(map4)
	fmt.Println(string(map4Json))
}
