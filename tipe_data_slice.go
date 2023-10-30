package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	// fmt.Println(len(slice1))
	// fmt.Println(cap(slice1))

	//Arry diubah, slice pun berubah
	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "Mei lagi"
	// fmt.Println(months)

	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Bayu")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)

	fmt.Println(months)

	newSclice := make([]string, 2, 5)
	newSclice[0] = "Bayu"
	newSclice[1] = "Bayu"

	fmt.Println(newSclice)

	copySlice := make([]string, len(newSclice), cap(newSclice))
	copy(copySlice, newSclice)
	fmt.Println(copySlice)

	//Notes!!
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
