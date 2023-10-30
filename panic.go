package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
		//Stop apps
	}

	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
}
