package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi is zero")
	}
	return nilai / pembagi, nil
}

func main() {
	//hasil, err := pembagian(2, 2)
	//
	//if err == nil {
	//	fmt.Println(hasil)
	//} else {
	//	fmt.Println(err.Error())
	//}

	checkError("Rendio")
}

// custom error handling

type validationData struct {
	Message string
}

type notFoundData struct {
	Message string
}

func (v *validationData) Error() string {
	return v.Message
}

func (v *notFoundData) Error() string {
	return v.Message
}
func saveData(data string, value any) error {
	if data == "" {
		return &validationData{"value is empty"}
	}
	if data != "Rendi" {
		return &notFoundData{"Data Not Found"}
	}
	return nil
}

func checkError(data string) {
	err := saveData(data, nil)

	if validationData, ok := err.(*validationData); ok {
		fmt.Println(validationData.Message)
	} else if notFoundData, ok := err.(*notFoundData); ok {
		fmt.Println(notFoundData.Message)
	} else {
		fmt.Println("Unknown Error", err.Error())
	}
}
