package main

import "fmt"

func main() {
	printMyName("John")
	fmt.Println("Hello, World!")
}

func printMyName(name string) {
	fmt.Println("My name is", name)
}
