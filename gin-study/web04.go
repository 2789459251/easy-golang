package main

// 原生渲染
import (
	"fmt"
	"html/template"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	//解析
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("解析错误：err：%v", err)
		return
	}
	//渲染
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("渲染错误：err：%v", err)
		return
	}

}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("连接失败,err:%v\n", err)
	}
}
