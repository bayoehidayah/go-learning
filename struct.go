package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Married       bool //Added
}

// Struct Function
func (customer Customer) sayHello() {
	fmt.Println("Hello My Name is", customer.Name)
}

func main() {
	var bayu Customer
	bayu.Name = "Bayu Hidayah"
	bayu.Address = "Indonesia"
	bayu.Age = 23

	fmt.Println(bayu)
	bayu.sayHello()
	fmt.Println(bayu.Name)
	fmt.Println(bayu.Address)
	fmt.Println(bayu.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30, true} //Harus terisi semua dan berurut berdasarkan posisi struct
	budi.sayHello()
	fmt.Println(budi)

	rully := Customer{Name: "Rully"}
	rully.sayHello()
}
