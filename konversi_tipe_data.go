package main

import "fmt"

func main() {
	//Int
	var nilai32 int32 = 98755
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	//String
	var name = "Bayu Hidayah"
	var e = name[0] //To byte
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
