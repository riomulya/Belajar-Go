package main

import "fmt"

func tampilPesan(nama string) {
	fmt.Println("Hallo ", nama)
}
func tambah(x int, y int, z int) (string, int) {
	jumlah := x + y + z
	tampil := "nilainya adalah"
	return tampil, jumlah
}
func faktorial(number int) int {
	if number == 0 {
		return 1
	} else {
		return number * faktorial(number-1)
	}
}

//map
func main() {
	harga_makanan := map[string]int{"Ayam": 10000, "bakso": 70000, "kue": 25000}
	fmt.Println("Ayam Rp \t", harga_makanan["Ayam"])
	fmt.Println("Bakso Rp\t", harga_makanan["bakso"])
	fmt.Println("Kue Rp  \t", harga_makanan["kue"])
	// fungsi
	tampilPesan("Rio Ganteng")
	fmt.Println(faktorial(5))
	fmt.Println(tambah(12000, 1231212, 234324))
}
