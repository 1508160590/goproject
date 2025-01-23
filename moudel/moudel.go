package moudel

type person struct {
	Name string
	Age  int
}

func Say(name string, age int) *person {
	return &person{
		Name: name,
		Age:  age,
	}
}
