package main

func main() {

	type NoKtp string

	var noKtp NoKtp = "39032"
	println(noKtp)
	println(NoKtp("432"))
	println(noKtp)

}
