package main

// 原生渲染
import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gebder string
	Age    int
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	//解析
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("解析错误：err：%v", err)
		return
	}
	//渲染
	u1 := User{
		Name:   "小王子",
		Gebder: "男",
		Age:    18,
	}
	m1 := map[string]interface{}{
		"name":   "小王子",
		"gebder": "男",
		"age":    18,
	}
	err = t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
	})
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
