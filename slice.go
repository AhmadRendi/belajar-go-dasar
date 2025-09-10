package main

import "fmt"

func main() {

	names := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
		"Rendi",
	}

	days := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	programSederhanaSlice(names)
	menghitungPanjangSlice(names)
	mendapatkanKapasitasSlice(names)

	appendSlice(days)
	makeSlice()
	copySlice(days)

}

func programSederhanaSlice(names []string) {

	// dengan cara seperti ini kita mengambil data dari index ke 0 sampai dengan index ke 2
	slice := names[0:2]

	fmt.Println(slice[0])
	fmt.Println(slice[1])

	// dengan cara ini, kita akan mengambil data dari index ke 3 sampai dengan akhir
	slice = names[3:]

	fmt.Println(slice[0])

	// dengan cara ini, kita akan mengambil data dari index ke 3 sampai paling awal
	slice = names[:3]

	fmt.Println(slice[2])
}

func menghitungPanjangSlice(names []string) {
	length := len(names)
	fmt.Println("Panjang dari slice : ", length)
}

func mendapatkanKapasitasSlice(names []string) {
	cop := cap(names)
	fmt.Println("Kapasitas Dari slice : ", cop)
}

// digunakan untuk melakukan penambahan data didalam slice
func appendSlice(days []string) {
	daySlice := days[5:]
	daySlice[0] = "Sabtu Baru"
	daySlice[1] = "Minggu Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
}

// digunakan untuk membuat slice baru
func makeSlice() {
	// membuat slice
	newSlice := make([]string, 3)
	newSlice[0] = "John"
	newSlice[1] = "Paul"
	newSlice[2] = "George"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
}

// melakukan copy pada slice
func copySlice(days []string) {
	toSlice := make([]string, len(days), cap(days))
	copy(toSlice, days)
	fmt.Println(toSlice)
}

func arrayVSSlice() {
	/*
		kita perlu berhati-hari saat membuat slice, karena jika kita salah maka yang kita buat adalah array dan
		bukanlah slice
	*/
	// ini adalah contoh deklarasi dan inisialisasi dari array
	intArray := [...]int{1, 3, 4, 5, 6, 7, 8, 9}

	// ini adalah contoh dari deklarasi dan inisialisasi dari slice
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(intArray)
	fmt.Println(intSlice)
}
