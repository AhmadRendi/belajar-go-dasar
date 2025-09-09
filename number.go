package main

import "fmt"

func main() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lime = ", 3.5)

	fmt.Println("Hasil Dari Menambahkan : ", add(10, 20))

}

func add(a int, b int) int {
	fmt.Println("Add a = ", a, "b = ", b)
	return a + b
}
