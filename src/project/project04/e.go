package main

import (
   "fmt"
   "strings"
)


func main() {
  name := "jSDD DoeDSCSDC"
  age := 27

  // 以小写形式打印字符串
  fmt.Printf("hello, %s\n", strings.ToLower(name))

  // 以大写形式打印字符串
  fmt.Printf("hello, %s\n", strings.ToUpper(name))

  // 以十进制数打印年龄
  fmt.Printf("%s is %d years old\n", name, age)

  // 年龄打印为带两位小数的浮点数
  fmt.Printf("%s is %.2f years old\n", name, age)
}