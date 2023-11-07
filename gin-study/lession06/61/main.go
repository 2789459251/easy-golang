package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//定义
	//解析
	t, err := template.ParseFiles("./base.tmpl", "./oil.tmpl")
	if err != nil {
		fmt.Println("解析错误 err：%v", err)
		return
	}
	//渲染
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("渲染错误 err：%v", err)
		return
	}
}
func main() {
	http.HandleFunc("/tmp", f1)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("连接失败,err:%v\n", err)
	}
}
