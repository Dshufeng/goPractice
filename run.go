package main

import (
	"fmt"
	"github.com/pkg/errors"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	//Caller()
	name, err := showErr("")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(name)
}

func Caller() {
	pc, file, line, _ := runtime.Caller(2)
	fmt.Printf("pc=%v file=%v line=%v\n", pc, file, line)

	functionName := runtime.FuncForPC(pc).Name()
	functionName = filepath.Ext(functionName)
	functionName = strings.TrimPrefix(functionName, ".")
	fmt.Println(functionName)
}

func showErr(name string) (str string, err error) {
	if name == "" {
		err = errors.New("name not null")
		return
	}
	str = name
	return
}
