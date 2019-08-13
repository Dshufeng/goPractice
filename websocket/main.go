package main

import (
	"fmt"
	"html/template"
	"net/http"

	"strings"

	"io"

	"golang.org/x/net/websocket"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("template err: ", err)
		return
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/upper", websocket.Handler(upper))
	if err := http.ListenAndServe("0.0.0.0:3000", nil); err != nil {
		fmt.Sprintln("http err=", err)
	}
	fmt.Println("listen on ")
}

func upper(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("ws recive err ", err)
			continue
		}
		send := strings.ToUpper(reply)
		if err = websocket.Message.Send(ws, send); err != nil {
			fmt.Println("ws send err ", err)
			continue
		}
	}
}
