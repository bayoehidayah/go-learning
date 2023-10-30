package main

import "fmt"

func main() {
	//Option 1
	var name string

	name = "Bayu Hidayah"
	fmt.Println(name)

	name = "Bayu Melandy"
	fmt.Println(name)

	//Option 2
	var name2 = "Bayoe"
	fmt.Println(name2)

	var age = 38
	fmt.Println(age)

	//Option 3
	country := "Indonesia" //Initialition
	fmt.Println(country)

	country = "Medan"
	fmt.Println(country)

	//Initialitation Multiple Variable
	var (
		firstName = "Bayu Hidayah"
		lastName  = "Melandy"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
