package main

func main() {

	var name string

	name = "Ahmad Rendi"

	println(viewName(name))

	// kita bisa membuat variable tanpa perlu sebutkan jenis type datanya asal kita langsung inisialisan datanya
	var age = 23

	println("My Age : ", viewAge(age))

	/*
		digolang ketika kita mendeklarasikan variable harus menggunakan keyword var
		namun kita juga bisa mendeklarasikan variable tanpa perlu menggunakan keyword var
	*/

	address := "Ibu Kota Nusantara Kalimantan Timur"

	println(viewName(address))

	// kita juga bisa mendeklarasikan multiple variable
	var (
		firstName string
		lastName  string
	)

	lastName2, firsName2 := "Ahmad", "Rendi"

	firstName = "Ahmad"
	lastName = "Rendi"

	println(viewName(firstName))
	println(viewName(lastName))

	println("Nama Lengkap : ", firstName, lastName)
	println("Nama Lengkap Kedua : ", firsName2, lastName2)
}

func viewName(name string) string {
	return "My name is " + name
}

func viewAge(age int) int {
	return age
}
func viewAddress(address string) string {
	return "My address is " + address
}
