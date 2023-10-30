package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello ", firstName, lastName)
}

// Returning Value
func getHello(name string) string {
	return "Hello " + name
}

// Returning Multiple Value
func getFullName() (string, string, string) {
	return "Bayu", "Hidayah", "Melandy"
}

// Named Return Value
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Bayu"
	middleName = "Hidayah"
	lastName = "Melandy"

	// return firstName, middleName, lastName (optional)
	return //(must)
}

// Variadic Function
// Parameter in the last params(must!)
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

// Function Value
func getGoodBye(name string) string {
	return "Good Bye " + name
}

// Function As Parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

type Filter func(string) string //With Type Declarations
func sayHelloWithFilter2(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// Anonymous Function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// Recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	sayHello()

	sayHelloTo("Bayu", "Hidayah")

	result := getHello("Bayu")
	fmt.Println(result)
	fmt.Println(getHello(""))

	//Multiple Value
	firstName, lastName, _ := getFullName() // _ = ignoring value
	fmt.Println(firstName, lastName)

	//Named Return Values
	first, middle, last := getCompleteName()
	fmt.Println(first, middle, last)

	//Variadic Function
	total := sumAll(10, 10, 10, 20, 30, 40)
	fmt.Println(total)
	//From slice
	numbers := []int{10, 10, 10, 20, 30}
	total = sumAll(numbers...)
	fmt.Println(total)

	//Function Value
	goodbye := getGoodBye
	fmt.Println(goodbye("Bayu"))

	//Function As Parameter
	sayHelloWithFilter("Bayu", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	//Anonymous function
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("bayu", blacklist)
	registerUser("anjing", func(s string) bool {
		return s == "anjing"
	})

	fmt.Println(factorialRecursive(5))
}
