package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server is running...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listen.Close()
	for {
		fmt.Println("开始")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			return
		}
		defer conn.Close()
		fmt.Println("连接成功")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		fmt.Println("Received:", string(buf[:n]))
		conn.Write([]byte("Hello, client!")) // 发送消息给客户端
	}
}
