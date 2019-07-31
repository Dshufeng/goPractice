package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filePath := "./test.txt"
	//data, _ := readLine(filePath)
	//fmt.Println(data)

	c, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(c))
}

// 逐行读取
func readLine(filePath string) (data []string, err error) {
	var c []string

	src, _ := os.Open(filePath)
	bfRd := bufio.NewReader(src)
	for {
		line, err := bfRd.ReadBytes('\n')
		c = append(c, string(strings.Replace(string(line), "\n", "", -1)))
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return c, nil
			}
			return nil, err
		}
	}
}
