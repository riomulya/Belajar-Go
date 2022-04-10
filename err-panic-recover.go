package main

import (
	"fmt"
	"strconv"
)

func main() {
	defer getRecover()
	var input string
	fmt.Print("Masukkan Nilai : ")
	fmt.Scanf("%s", &input)
	// fmt.Println("Nilainya Adalah : ", input)
	var number int
	var err error

	number, err = strconv.Atoi(input)
	// print(err)
	// print(number)

	if err == nil {
		fmt.Println(number, " Ini Adalah Angka")
	} else {
		panic(err.Error())
	}
}

func getRecover() {
	if re := recover(); re != nil {
		fmt.Println("Eror Salah ")
	}
}
