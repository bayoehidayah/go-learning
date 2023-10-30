package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	nama := newMap("")
	if nama == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(nama)
	}
	// var person map[string]string = nil
	// fmt.Println(person)
}
