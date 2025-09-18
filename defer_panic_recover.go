package main

import "fmt"

// penggunaan dari defer

func logging() {
	fmt.Println("Inside logging")
}

func runApp() {
	defer logging()
	fmt.Println("Inside runApp")
}

// penggunaan dari panic

func endApp() {
	fmt.Println("Inside endApp")
	message := recover()

	fmt.Println(message)
}

func runApps(err bool) {
	defer endApp()
	if err {
		panic("Telah Terjadi error disini")
	}
}

func main() {

	runApp()

	runApps(true)

}
