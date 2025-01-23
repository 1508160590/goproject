package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//修改默认标识符
	//template.New("test").Delims("{[","]}").ParseFiles("./t.tmpl")

	t, err := template.ParseFiles("./base.tmpl", "./index.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed err:%v", err)
		return
	}
	msg := "这是index页面"
	t.ExecuteTemplate(w, "index.tmpl", msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("base.tmpl", "./home.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed err:%v", err)
		return
	}
	msg := "这是home页面"
	t.ExecuteTemplate(w, "home.tmpl", msg)
}

func main() {
	http.HandleFunc("/index", index)         //创建引擎路由
	http.HandleFunc("/home", home)           //创建引擎路由
	err := http.ListenAndServe(":9090", nil) //启动服务
	if err != nil {
		fmt.Printf("http serve failed err:%v", err)
		return
	}
}
