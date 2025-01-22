package main

import (
	"fmt"
	"net/http"
	"html/template"
)
func sayHello(w http.ResponseWriter,r *http.Request){
	//1.定义模板hello。tmpl
	//2.解析模板
	//func(t*Template)Parse(src string)(*Template, error)
	//func ParseFiles(filenames ...string)(*Template, error)
	//func ParseGlob(pattern string)(*Template, error)
	
	t,err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed err:%v",err)
		return
	}
	//3.渲染模板
	//func(t*Template)Execute(wr io.Writer, data interface{})error
	//func (t *Template)ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	
	err = t.Execute(w,"红蓝毛兔奇侠传")
	if err != nil{
		fmt.Printf("render template failed err:%v",err)
		return
	}
}

func main() {
	http.HandleFunc("/hello",sayHello)//路由
	err := http.ListenAndServe(":9090",nil)//启动服务
	if err != nil{
		fmt.Printf("http serve failed err:%v",err)
		return
	}
}