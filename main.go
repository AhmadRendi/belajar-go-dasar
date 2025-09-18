package main

import (
	"fmt"
	"golang-dasar/helper"
	_ "golang-dasar/learn" // ini digunakan agar ketika kita import package tanpa menggunakannya tidak error
)

func main() {
	result := helper.SayHello("Ahmad Rendi")
	fmt.Println(result)
	yeah := helper.Example("yeaahh")
	fmt.Println(yeah)

	helper.CallMayName(yeah)
}
