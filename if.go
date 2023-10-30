package main

import "fmt"

func main() {
	name := "Bayu"

	if name == "Bayu" {
		fmt.Println("Hello Bayu")
	} else if name == "Bayu Hidayah" {
		fmt.Println("Hello Bayu Hidayah")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	//Short Statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
