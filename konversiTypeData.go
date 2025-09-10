package main

func main() {
	var nilai32 int32 = 9302
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	println(nilai32)
	println(nilai64)
	println(nilai16)

	var name string = "Ahmad Rendi"
	var e = name[7]
	var eString = string(e)

	println(name)
	println(eString)
}
