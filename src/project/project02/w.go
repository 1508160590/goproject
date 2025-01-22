package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		} else {
			return name + suffix
		}
	}
}

func main() {
	//统计字符串的长度，按字节1en(str)
	//golang的编码统一为utf-8(ascii的字符(字母和数字)占一个字节，汉字占用3个字节)
	str := "hello北"
	fmt.Println("str len", len(str))
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
		fmt.Printf("%c\n", str[i])
	}
}
