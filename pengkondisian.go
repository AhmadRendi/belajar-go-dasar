package main

import "fmt"

func main() {

	percabanganIf()
	percabanganIfElse()
	shortStatement()
	switchCase(1)
	switchShort("Ahma")
}

func percabanganIf() {

	name := "Rendii"

	if name == "Rendi" {
		fmt.Println("Nama " + name)
	} else {
		fmt.Println("Tidak Ada nama yang ditemukan")
	}

}

func percabanganIfElse() {
	name := "Rendi"

	if name == "Rendi" {
		fmt.Println("Nama " + name)
	} else if name == "Rendii" {
		fmt.Println("Nama " + name)
	} else {
		fmt.Println("Tidak Ada nama yang ditemukan")
	}
}

func shortStatement() {
	name := "Rendi"

	fmt.Println(len(name))
	if lengthName := len(name); lengthName > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Cocok")
	}
}

func switchCase(pilihan int) {
	switch pilihan {
	case 1:
		fmt.Println("Pilihan anda adalah 1")
		break
	case 2:
		fmt.Println("Pilihan anda adalah 2")
		break
	case 3:
		fmt.Println("Pilihan anda adalah 3")
		break
	case 4:
		fmt.Println("Pilihan anda adalah 4")
		break
	default:
		fmt.Println("Pilihan anda tidak tersedia")
	}
}

func switchShort(name string) {
	switch lengthName := len(name); lengthName >= 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Cocok")
	}
}
