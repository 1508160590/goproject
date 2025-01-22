package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func demo1(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./t.tmpl", "./u.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed err:%v", err)
		return
	}
	name := "光头强"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/template", demo1)      //路由
	err := http.ListenAndServe(":9000", nil) //启动服务
	if err != nil {
		fmt.Printf("http serve failed err:%v", err)
		return
	}
}
