package main
import "fmt"


func y(n int) int {
	if n == 10{
		return 1
	}else{
		return (y(n+1)+1) *2
	}
}
func getSum(n1 int,n2 int) int {
	return n1 + n2
}
func main() {
	/*
	num := 9
	ff := &num
	fff := &ff
	**fff = 10
	fmt.Println("num:",*fff,&fff,fff)
	fmt.Println(*ff,&ff,ff)
	fmt.Println(&num,num)
	
	var n int = 0
	i := 1000000.00
	for{
		if i == 0{
			break
		}
		if i > 50000{
			i = i * 0.05
		}else{
			i -= 1000
		}
	n++
	}
	fmt.Println(n)
	*/
	fmt.Println(y(10))
	a := getSum
	fmt.Println(a(10,40))
}