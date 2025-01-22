package main
import "fmt"


func main(){
	a := 10
	switch {
	   default : 
		  fmt.Println("default")
	   
	   case a > 0 : 
		  fmt.Println("a > 0")
	   
	   case a >5 : 
		  fmt.Println("a > 5")
	}
}