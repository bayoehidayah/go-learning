package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

// Interface kosong
func Ups(i int) interface{} {
	if i == 1 {
		return false
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var eko Person
	eko.Name = "Bayu"

	sayHello(eko)

	cat := Animal{
		Name: "Push",
	}

	sayHello(cat)

	woo := Ups(2)
	fmt.Println(woo)
}
