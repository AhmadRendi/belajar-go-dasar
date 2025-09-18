package main

import "fmt"

type Address struct {
	Street, City, State string
}

func beforePointer() {
	address1 := Address{
		Street: "123",
		City:   "San Francisco",
		State:  "CA",
	}
	address2 := address1

	address2.Street = "12345"

	fmt.Println(address1)
	fmt.Println(address2)
}

func afterUsingPointer() {
	address1 := Address{
		Street: "123",
		City:   "San Francisco",
		State:  "CA",
	}
	// kita menggunakan pass by reference
	address2 := &address1

	address2.Street = "12345"

	fmt.Println(address1)
	fmt.Println(address2)
}

func main() {
	//beforePointer()
	//afterUsingPointer()
	//beforeUsingAsterisk()
	//afterUsingAsterisk()

	address := Address{
		Street: "123",
		City:   "San Francisco",
		State:  "CA",
	}

	withoutPointerParameter(address)

	fmt.Println(address)

	withPointerParameter(&address)

	fmt.Println(address)

	address.printAddress()
}

func beforeUsingAsterisk() {
	address := Address{
		Street: "123",
		City:   "San Francisco",
		State:  "CA",
	}

	address2 := &address

	address2.City = "Bandung"

	address2 = &Address{
		Street: "12345",
		City:   "California",
		State:  "US",
	}

	fmt.Println(address)
	fmt.Println(address2)
}

func afterUsingAsterisk() {
	address := Address{
		Street: "123",
		City:   "San Francisco",
		State:  "CA",
	}

	address2 := &address

	address2.City = "Bandung"

	*address2 = Address{
		Street: "12345",
		City:   "California",
		State:  "US",
	}

	fmt.Println(address)
	fmt.Println(address2)
}

func withoutPointerParameter(address Address) {
	address.State = "Indonesia"
}

func withPointerParameter(address *Address) {
	address.Street = "Jl. Kalibuana Singgah"
}

func (address *Address) printAddress() {
	fmt.Println(address.Street, address.City, address.State)
}
