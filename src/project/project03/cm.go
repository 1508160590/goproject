package main

import (
    "fmt"
    "time"
)

func main() {
    startTime := time.Now() // 获取开始时间

    for i := 0; i < 100000; i++ {
        fibonacci(i)
    }

    // 计算并输出程序执行时间
    duration := time.Since(startTime)
    fmt.Println("Execution time in milliseconds:", duration.Milliseconds())
}

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }

    fib := make([]int, n+1)
    fib[0] = 0
    fib[1] = 1

    for i := 2; i <= n; i++ {
        fib[i] = fib[i-1] + fib[i-2]
    }

    return fib[n]
}
