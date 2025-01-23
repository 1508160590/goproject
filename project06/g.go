package main
import (
"fmt"
"unsafe"
)

func main(){
	//var a string
	//var b int
	//var c,d,e float64
	var f,g,h = 5,6.3,false
	s,m,w := 55,"acc",0.25
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	fmt.Println(a,b,c)
	fmt.Println(f,g,h)
	fmt.Print(s,m,w)
}