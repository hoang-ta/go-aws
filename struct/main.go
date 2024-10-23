package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age: age,
	}
} 

func (person *Person) changeName(newName string) {
	fmt.Println(&person.Name)
	person.Name = newName
}

func main() {
	person := Person{
		Name: "Hoang",
	}


	fmt.Println(person)
	person.Name = "Michael"
	fmt.Printf("%+v\n", person)
	fmt.Println(&person.Name)
	person.changeName("Hoang")
	fmt.Printf("%+v\n", person)

	person2 := NewPerson(
		"Hoang 2",
		12,
	)
	fmt.Printf("Person 2 %+v\n", person2)
	person2.changeName("Hoang 1")
	fmt.Printf("Person 2 %+v\n", person2)

	a := 1
	b := &a
	println(a, b)
	*b = 2
	println(a, *b)
}