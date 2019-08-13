package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
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

	type People struct {
		Name string
		Age  int
	}

	var people1 People
	people1.Name = "wang"
	people1.Age = 18

	fmt.Println(people1)
	people2 := People{
		"li", 19,
	}
	var name = "Name"
	pp := reflect.ValueOf(&people2)
	field := pp.Elem().FieldByName(name)
	field.SetString("lil")

	fmt.Println(people2)

	str1 := "hello世界你好"
	i := strings.IndexRune(str1, '你')
	c := utf8.RuneCountInString(str1)
	fmt.Println(i, c, len([]rune(str1)))

	i1 := UnicodeIndex(str1, "你")
	fmt.Println(i1)
}

func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}
	return result
}
