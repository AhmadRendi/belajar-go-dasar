package main

import "fmt"

type Customer struct {
	FirstName, LastName string
	Age                 int
}

func sayHello(customer Customer) {
	fmt.Println("Hello " + customer.FirstName + " " + customer.LastName)
}

func main() {

	var rendi Customer

	rendi.FirstName = "Ahmad"
	rendi.LastName = "Rendi"
	rendi.Age = 23

	fmt.Println(rendi)

	ahmad := Customer{
		FirstName: "Ahmad",
		LastName:  "Rendi",
		Age:       23,
	}

	fmt.Println(ahmad)

	budi := Customer{"Budi", "Gunawan", 40}
	fmt.Println(budi)

	sayHello(rendi)
	sayHello(ahmad)
	sayHello(budi)

	personStruct := PersonStruct{
		"Ahmad",
		40,
		"Binuang",
	}

	sayHelloForHasName(personStruct)

	dataMap := newMap("")

	fmt.Println(dataMap)

	personStruct2 := PersonStruct{
		"Budi Gunawan",
		50,
		"Jakarta Selatan",
	}

	personStruct3 := PersonStruct{
		"Jon Snow",
		28,
		"Jakarta Selatan",
	}

	mapData := make(map[string]PersonStruct)

	newMapPerson(mapData, personStruct2)
	newMapPerson(mapData, personStruct3)

	fmt.Println(mapData)

}

// Data Nil

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}

func newMapPerson(m map[string]PersonStruct, person PersonStruct) {
	m[person.Name] = person
}

//type HasName interface {
//	GetName() string
//}

func sayHelloForHasName(name PersonService) {
	fmt.Println(name.GetName())
}

type PersonStruct struct {
	Name    string
	Age     int
	Address string
}

func (personStruct PersonStruct) GetAge() int {
	//TODO implement me
	panic("implement me")
}

func (personStruct PersonStruct) GetAddress() string {
	//TODO implement me
	panic("implement me")
}

func (personStruct PersonStruct) GetName() string {
	return personStruct.Name
}

type PersonService interface {
	GetName() string

	GetAge() int

	GetAddress() string
}
