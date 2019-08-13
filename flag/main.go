package main

import (
	"flag"
)

var f1 int
var h *bool

func init() {
	flag.IntVar(&f1, "f1", 1234, "f1")
	h = flag.Bool("h", false, "help")
}

// go build main.go
// main -h
// main f1 1212
func main() {
	flag.Parse()
	if *h {
		usage()
	}
}

func usage() {
	flag.PrintDefaults()
}
