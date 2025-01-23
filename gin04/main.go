package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 自定义一个夸人的模板函数
	kua := func(name string) string {
		return name + "真帅"
	}
	//1.定义模板hello.tmpl
	//2.解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed err:%v", err)
		return
	}
	// 告诉模板引擎，我现在多了一个自定义的函数
	t.Funcs(template.FuncMap{
		"kua": kua,
	})
	//3.渲染模板
	ui := User{
		Name:   "红猫",
		Gender: "男",
		Age:    21,
	}
	m1 := map[string]interface{}{
		"name":   "蓝兔",
		"gender": "女",
		"age":    18,
	}
	hobbyList := []string{
		"猪猪侠",
		"菲菲",
		"呆呆",
		"波比",
		"超人强",
	}
	t.Execute(w, map[string]interface{}{
		"u1":    ui,
		"m1":    m1,
		"hobby": hobbyList,
	})
	if err != nil {
		fmt.Printf("Execute template failed err:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/hello", sayHello)      //路由
	err := http.ListenAndServe(":9090", nil) //启动服务
	if err != nil {
		fmt.Printf("http serve failed err:%v", err)
		return
	}
}
