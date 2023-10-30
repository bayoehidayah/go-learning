package main

import "fmt"

func endApps() {
	message := recover() //Get panic message
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}

	fmt.Println("Aplikasi selesai")
}

func runApps(error bool) {
	defer endApps()
	if error {
		panic("Aplikasi Error")
		//Stop apps
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApps(true)
	fmt.Println("Bayu")
}
