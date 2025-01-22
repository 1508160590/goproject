package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	d, _ := os.ReadFile("./hello.txt")
	fmt.Fprintf(w, string(d)) //返回信息
}

func main() {
	http.HandleFunc("/hello", sayHello)      //路由
	err := http.ListenAndServe(":9090", nil) //启动服务
	if err != nil {
		fmt.Printf("http serve failed err:%v", err)
		return
	}
}
