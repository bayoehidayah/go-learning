package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var ktpBayu NoKTP = "127102312311123"
	var marriedStatus Married = false
	fmt.Println(ktpBayu)
	fmt.Println(marriedStatus)
}
