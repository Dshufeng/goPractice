package main

import "fmt"

type IPeople interface {
	hello()
}
type People struct {
}

func (p *People) hello() {
	fmt.Println("github.com/meetbetter")
}

//错误：将nil的people给空接口后接口就不为nil,因为interface中的value为nil但type不为nil
func errFunc() *People {
	var p *People

	return p
}

//正确处理返回nil给接口的方式：直接将nil赋给interface
func rightFunc() IPeople {
	var p *People
	return p
}
func main() {
	if errFunc() == nil {
		fmt.Println("对了哦，外部接收到也是nil")
	} else {
		fmt.Println("错了咦，外部接收到不是nil哦")
	}

	if rightFunc() == nil {
		fmt.Println("对了哦，外部接收到也是nil")
	} else {
		fmt.Println("错了咦，外部接收到不是nil哦")
	}
}
