package main

import "fmt"

type Dog struct {
	Name   string
	Age    int
	weight float32
}

func (d *Dog) Say() string {
	return fmt.Sprintf("Name: %s, Age: %d, Weight: %.2f", d.Name, d.Age, d.weight)
}
func main() {
	d := Dog{Name: "Rufus", Age: 3, weight: 50.5}
	d.Say()
	fmt.Println(d.Say())
}
