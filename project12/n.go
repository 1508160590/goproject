package main
import "fmt"
import "strconv"

func main() {
	var str string = " 0."
	var num int64
	num, _ = strconv.ParseInt(str,10,64)
	fmt.Printf("num type = %T num=%d",num,num)
}