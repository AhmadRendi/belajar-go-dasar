package main

import "fmt"

func main() {

	fmt.Println("Ahmad")
	fmt.Println("Ahmad Rendi")
	fmt.Println("Rendi")

	// Fungsi Untuk Mengetahui Panjang dari string
	fmt.Println(len("Ahmad Rendi"))

	// cara mengabil karakter dari pasisi tertentu
	fmt.Println(string("Ahmad Rendi Abdul Malik"[10]))
	fmt.Println(string("Rendi"[1]))

	//nama string;
	//nama = "Ahmad Rendi";

	println(tampikanNama("Ahmad Rendi"))

}

func tampikanNama(nama string) string {
	return "Nama Saya adalah : " + nama
}
