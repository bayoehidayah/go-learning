package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Bayu"
	names[1] = "Hidayah"
	names[2] = "Melandy"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)

	fmt.Println("Panjang =", len(names))

	var values = [3]int{
		90,
		80,
		95,
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values)

	fmt.Println("Panjang =", len(values))
}
