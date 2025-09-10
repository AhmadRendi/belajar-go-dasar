package main

import "fmt"

func main() {
	var a [10]int
	a[0] = 30
	a[1] = 40

	println(a[0], a[1])

	// Kita Juga Bisa Langsung mendeklarasikan value dari erray
	var b = [10]int{
		30, 40, 4320,
	}

	println(b[0], b[1])

	// mengetahui panjang dari array
	println(len(b))

	// mendapatkan posisi data didalam index array
	println(a[1])

	// mengubah data array berdasarkan index
	a[1] = 50
	println(a[1])

	var c = [10]string{
		"Ahmad Rendi",
		"Ahmad",
		"Rendi",
	}

	println("Nama Saya adalah : ", c[0])
	fmt.Println(c)
}
