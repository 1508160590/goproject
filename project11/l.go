package main
import "fmt"

func main() {
	/*
	var num int = 0
	for i := 1;i<=100;i++{
		if i % 9 == 0{
			num += i
		}
	}
	fmt.Println(num)
	
	fum := 0.0
	n := 0
	for i := 1;i<=5;i++{
		var sum float64
		fmt.Scanln(&sum)
		if sum >= 60{
			n += 1
		}
		fum += sum
	}
	fmt.Println(fum / 5)
	fmt.Println(n)
	*/
	for i:=1;i<6;i++{
		for k :=0;k<=6-i;k++{
			fmt.Print(" ")
		}
		for n:=0;n<2*i-1;n++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	
	for i:=1;i<=6;i++{
		for k:=1;k<=6-i;k++{
			fmt.Print(" ")
	}
		for n:=1;n<=2*i-1;n++{
			if n == 1 || n == 2*i-1 || i==6{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	
	for i:=1;i<=9;i++{
		for n:=1;n<=i;n++{
			fmt.Printf("%d * %d = %d\t",i,n,i*n)
		}
		fmt.Println()
	}
}