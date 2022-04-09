package main

import "fmt"

func main() {
	nomor := 23
	var alamat_nomor *int = &nomor
	fmt.Println("Alamat Dari Nomor pointer", alamat_nomor)
	fmt.Println("Alamat Dari Nomor ", &nomor)
	if alamat_nomor == &nomor {
		println("Benar")
	} else {
		println("salah")
	}
}
