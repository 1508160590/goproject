package main

import "fmt"

// 判断年份是否为闰年
func isLeapYear(year int) bool {
    if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
        return true
    }
    return false
}

// 判断日期是否有效
func isValidDate(year, month, day int) bool {
    // 判断月份范围
    if month <= 0 || month > 12 {
        return false
    }

    // 每个月的天数
    daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    
    // 如果是闰年，二月天数为29
    if isLeapYear(year) {
        daysInMonth[1] = 29
    }
    
    // 判断日期是否有效
    if day <= 0 || day > daysInMonth[month-1] {
        return false
    }

    return true
}

func main() {
    for {
        var y, m, r int
        
        // 获取年份
        fmt.Println("请输入年份：")
        fmt.Scanf("%d", &y)

        // 获取月份
        fmt.Println("请输入月份：")
        fmt.Scanf("%d", &m)

        // 获取日期
        fmt.Println("请输入日期：")
        fmt.Scanf("%d", &r)

        if !isValidDate(y, m, r) {
            fmt.Println("无效的日期，请重新输入！")
            continue
        }

        fmt.Printf("有效日期：%d年%d月%d日\n", y, m, r)
        break
    }
}
