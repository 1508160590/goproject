package main

import "fmt"

func main() {
    const (
		a = iota
		b
		c
		d = 100
		e = iota
		f
    )
    fmt.Println(a,b,c,d,e,f)
}