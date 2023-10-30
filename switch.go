package main

import "fmt"

func main() {
	name := "Bayu"

	switch name {
	case "Bayu":
		fmt.Println("Hello Bayu")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hi, Boleh kenalan?")
	}

	//Short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	//Without Expression
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama sedikit panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
