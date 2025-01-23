package main

import (
	"fmt"
)

type mysum func(int, int) int

func sum(a, b int) int {
	return a + b
}
func sum2(a, b, c int) int {
	return a + b
}
func myFunc(Var mysum, a, b int) int {
	return Var(a, b)
}

func main() {
	a := sum
	b := sum2
	c := myFunc(a, 10, 20)
	fmt.Println(c)
	fmt.Println(b)
}
