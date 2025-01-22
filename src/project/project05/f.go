package main

import(
"fmt"
"time"
)


func main() {
	startTime := time.Now() // 获取开始时间
	for i := 1; i < 50; i++ {
		f(i) 
	}
	// 计算并输出程序执行时间
    duration := time.Since(startTime)
    fmt.Println("Execution time in milliseconds:", duration.Milliseconds())
}

func f(n int) int {
	switch {
	case n == 1:
		return 1
	case n == 2:
		return 2
	default:
		return f(n-1) + f(n-2)
	}
}
