package main

import (
	"fmt"
)

func main() {

	// ini adalah array name
	names := [...]string{
		"Ahmad Rendi",
		"Ahmad",
		"Rendi",
		"Supri",
		"Joko",
	}

	namesSlice := []string{
		"Ahmad Rendi",
		"Ahmad",
		"Rendi",
		"Supri",
		"Joko",
	}
	forCounter(0)
	forWithArray(names)
	forWithSlice(namesSlice)
	forWithRange(names)
	forWithRangeIndex(names)
	forWithBreak(names)

	forWithContinue(names)
}

func forCounter(counter int) {
	for counter < 10 {
		fmt.Println("Perulangan ke- ", counter+1)
		counter++
	}
	fmt.Println("Selesai")
}

func forWithArray(names [5]string) {
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}

func forWithSlice(names []string) {
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}

func forWithRange(names [5]string) {
	for name := range names {
		fmt.Println(names[name])
	}

	for _, name := range names {
		fmt.Println(name)
	}
}

func forWithRangeIndex(names [5]string) {
	for index, name := range names {
		fmt.Println("Nama Ke- ", index, name)
	}
}

func forWithBreak(names [5]string) {
	for index, name := range names {
		if index == 2 {
			break
		}
		fmt.Println("Nama Ke- ", index, name)
	}
}

func forWithContinue(names [5]string) {
	for index, name := range names {
		if index == 2 {
			continue
		}
		fmt.Println("Nama Ke- ", index, name)
	}
}
