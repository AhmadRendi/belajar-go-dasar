package main

func main() {
	const name = "Ahmad Rendi"
	const age = 10

	println("My Name Is ", name)
	println("My Age Is ", age)

	const (
		names string = "Ahmad Rendi"
		ages         = 10
	)

	println(names, ages)
}
