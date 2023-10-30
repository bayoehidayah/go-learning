package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian Dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, error := Pembagian(12, 2)

	if error != nil {
		fmt.Println(hasil)
	} else {  
		fmt.Println(error.Error())
	}
}
