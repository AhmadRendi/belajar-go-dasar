package helper

func SayHello(name string) string {
	return "Hello : " + name
}

// ini adalah access modifier agar tidak dapat dipanggil diluar package
func callMayName(name string) string {
	return name
}

func CallMayName(name string) string {
	return name
}
