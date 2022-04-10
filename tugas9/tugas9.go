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

	nilai, err := strconv.Atoi(input)

	if err == nil {
		fmt.Println(nilai, " Ini Adalah Angka")
	} else {
		panic(err.Error())
	}
}

func getRecover() {
	if re := recover(); re != nil {
		fmt.Printf("Salah Anda Memasukkan %T Bukan Int", re)
	}
}
