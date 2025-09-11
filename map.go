package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// format dari pembuatan field atau atribut di go adalah pascalcase
	Name, Address string
	age           int
}

func main() {

	kodeProgramSederhana()
	kodeArrayWithMap()
	kodeArrayWithMapTwo()
}

func kodeProgramSederhana() {
	person := map[string]string{
		"name":    "John",
		"age":     "25",
		"address": "San Francisco",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["address"])
}

func kodeArrayWithMap() {
	arrayMap := [5]map[int]string{}

	for i := 0; i < 5; i++ {
		person2 := map[int]string{
			i: "Rendi : " + strconv.Itoa(i),
		}
		arrayMap[i] = person2
	}

	fmt.Println(arrayMap)
	fmt.Println(len(arrayMap))
	fmt.Println(arrayMap[3])
	fmt.Println(arrayMap)
}

func kodeArrayWithMapTwo() {
	mapNew := make(map[string]Person)

	person := Person{
		Name:    "John",
		Address: "San Francisco",
		age:     25,
	}

	mapNew[person.Name] = person

	personBudy := Person{
		Name:    "Budy",
		Address: "San Francisco",
		age:     20,
	}

	mapNew[personBudy.Name] = personBudy
	fmt.Println(mapNew)
	// menampilkan data dengan key John
	fmt.Println(mapNew[person.Name])

	// Menghapus data dengan key John
	delete(mapNew, person.Name)

	fmt.Println(mapNew)

}
