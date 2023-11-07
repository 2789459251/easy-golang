package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("main.tmpl").Delims("{[", "]}").ParseFiles("./main.tmpl")
	if err != nil {
		fmt.Println("解析错误 err:%v", err)
		return
	}
	name := "小王子"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("连接错误 err:%v", err)
		return
	}
}
