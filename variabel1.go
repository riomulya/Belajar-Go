package main

import (
	"fmt"
)

func main() {
	// var namaGw string = "Rio Mulya Syawal"
	// namaOrang := 12.9

	//tipe data go
	// var Bilbul uint = 5000000000000000000
	// var bilnegatif = -23
	// // const x = "string X"
	// fmt.Printf("Hallo %s \nHallo Pak ,%.2f\n%d\n\n", namaGw, namaOrang, Bilbul)
	// fmt.Println(bilnegatif)
	// // fmt.Println(x)

	// var value = (((2+6)%3)*4 - 2) / 3
	// var isEqual = (value == 2)

	// fmt.Printf("nilai %d (%t) \n", value, isEqual)
	var x = 10
	var z = 50
	if x < z {
		fmt.Println("Z lebih besar dari x")
	} else if x == z {
		fmt.Println("Z Sama Dengan x")
	} else {
		fmt.Println("Z lebih kecil dari x")
	}
	for x < 40 {
		println(x - 9)
		x++
	}

	fruits := [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("\nJumlah element \t\t", len(fruits))
	for i := 0; i < len(fruits); i++ {
		fmt.Println("Isi element ke -", i+1, "adalah ", fruits[i])
	}
}
