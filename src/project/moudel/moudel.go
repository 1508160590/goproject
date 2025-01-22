package moudel

type person struct {
	Name string
	Age  int
}

func (p *person) Say(name string, age int) *person {
	return &person{
		Name: name,
		Age:  age,
	}
}
