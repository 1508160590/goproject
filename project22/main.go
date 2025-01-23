package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Salary   float64
}

func main() {
	// monster := Monster{
	// 	Name:     "张三",
	// 	Age:      18,
	// 	Birthday: "2022-01-01",
	// 	Salary:   8000.00,
	// }
	// data, err := json.Marshal(&monster) //返回值是 字节切片和一个异常
	// if err != nil {
	// 	fmt.Println("序列化失败: ", err)
	// }
	// fmt.Println(string(data)) //{"Name":"张三","Age":18,"Birthday":"2022-01-01","Salary":8000}
	monster := new(Monster) //可以直接使用Monster{}，不用new开辟内存，Unmarshal中会new
	jsonStr := "{\"Name\":\"张三\",\"Age\":18,\"Birthday\":\"2022-01-01\",\"Salary\":8000}"
	json.Unmarshal([]byte(jsonStr), &monster) //转换，参数为字节切片和要转换的类型指针
	fmt.Println(*monster)                     //{张三 18 2022-01-01 8000}
}
