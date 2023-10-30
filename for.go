package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan Ke ", counter)
		counter++
	}

	slice := []string{"Bayu", "Hidayah", "Melandy"}
	for i, value := range slice {
		fmt.Println("Index ", i, "=", value)
	}

	person := map[string]string{
		"name":   "Bayu",
		"middle": "Hidayah",
	}
	for i, value := range person {
		fmt.Println("Index ", i, "=", value)
	}
}
