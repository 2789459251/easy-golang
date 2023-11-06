package main

import (
	"fmt"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "<h1>hello goland!</h1>")
}
func main() {
	http.HandleFunc("/hello", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http serve fail err:%v\n", err)
		return
	}
}
