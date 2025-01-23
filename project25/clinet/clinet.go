package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server")
	for {
		var message string
		fmt.Print("Enter message: ")
		fmt.Scanln(&message)
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing to server:", err)
			return
		}
		response := make([]byte, 1024)
		_, err = conn.Read(response)
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}
		fmt.Println("Server response:", string(response))
	}
}
