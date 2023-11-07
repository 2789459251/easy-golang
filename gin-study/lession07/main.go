package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./base.tmpl", "./index.tmpl")
	if err != nil {
		fmt.Println("解析 err:%v", err)
		return
	}
	name := "玫瑰"
	t.Execute(w, name)

}
func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./base.tmpl", "./home.tmpl")
	if err != nil {
		fmt.Println("解析 err:%v", err)
		return
	}
	name := "小王子"
	t.Execute(w, name)
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("连接失败,err:%v\n", err)
	}
}
