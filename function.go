package main

import (
	"fmt"
	"strconv"
	"strings"
)

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("Your black list", name)
	} else {
		fmt.Println("Your unblack list", name)
	}
}

func anonymouseFunc(name string) {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser(name, blacklist)

	// atau bisa juga seperti ini
	registerUser(name, func(s string) bool {
		return s == "Anjing"
	})
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return factorialRecursive(value-1) + value
	}
}

func closures() {
	counter := 0

	increment := func() {
		fmt.Println("counter", counter)
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}

func main() {

	// cara memanggil multiple return
	name, age := multipleDataReturn("Ahmad Rendi", 23)

	fmt.Println(name, " Umur Saya : ", strconv.Itoa(age))

	// Kita Juga Bisa menghiraukan kembalian salah satu data
	names, _ := multipleDataReturn("Ahmad Rendi", 23)
	_, ages := multipleDataReturn("Ahmad Rendi", 23)

	fmt.Println(names)
	fmt.Println(ages)

	// Kita juga dapat mengembalikan name data seperti ini
	firstName, lastName := nameReturnData()

	fmt.Println(firstName, lastName)

	// Kita juga bisa membuat sebuah function dengan parameter seperti ini
	total := variadicFunction(2, 5, 3, 53, 323, 532, 123)
	fmt.Println("total : ", total)

	value := []int{423, 523, 523, 623, 234, 523}
	totalSlice := variadicFunction(value...)
	fmt.Println("Total Slice : ", totalSlice)

	// kita dapat menyimpan sebuah function didalam sebuah type data
	variableFuncName := functionValue

	// menggunakan function sebagai sebuah parameter
	fmt.Println(variableFuncName("Ahmad Rendi"))

	//
	sayMyNameWithFilter("Anjing", spamFilter)

	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	printEvenNumber(1, number, evenNumber)

	howMuchCharacter("Ahmad Rendi Abdul Malik")

	myName := "ahmad Rendi"

	// menghitung berapa kali muncul sebuah karakter didalam string
	count := strings.Count(myName, "a")

	fmt.Println(count)

	blackList := func(name string) bool {
		return name == "Ahmad Rendi"
	}

	registerUser("Ahmad Rendi", blackList)
	anonymouseFunc("Anjing")

	fmt.Println(factorialLoop(8))

	fmt.Println(factorialRecursive(2))

	closures()
}

func multipleDataReturn(name string, age int) (string, int) {
	return "Nama Saya adalah : " + name, age
}

func nameReturnData() (firstName, lastName string) {
	firstName = "Ahmad"
	lastName = "Rendi"

	return firstName, lastName
}

func variadicFunction(value ...int) int {
	number := 0

	for _, v := range value {
		number += v
	}
	return number
}

func functionValue(name string) string {
	return name
}

func spamFilter(name string) string {
	if name == "Anjing" ||
		name == "anjing" {
		return "..."
	}
	return name
}

func sayMyNameWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello ", filter(name))
}

func printEvenNumber(number int, value []int, filter func(int, []int)) {
	filter(number, value)
}

func evenNumber(number int, value []int) {
	if number == 1 {
		for _, v := range value {
			if v%2 != 0 {
				fmt.Println("Odd Number : ", v)
			}
		}
	} else if number == 2 {
		for _, v := range value {
			if v%2 == 0 {
				fmt.Println("Even Number : ", v)
			}
		}
	}
}

func howMuchCharacter(value string) {
	runes := []rune(value)
	for i, r := range runes {
		ch := string(r)

		if strings.Index(value, ch) != i {
			continue
		}

		if ch == " " {
			continue
		}

		number := strings.Count(value, ch)

		fmt.Println(ch, " : ", number)
	}
}
