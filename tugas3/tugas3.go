package main

import "fmt"

func main() {
	buah := []string{"apel", "jeruk", "anggur", "melon"}
	buah = append(buah, "pepaya", "naga")
	buah = append(buah, "naga")
	buah = append(buah, "jambu")
	fmt.Println("Jumlah Element : ", len(buah))
	fmt.Println("Isi Element : ", buah)

	for i := 0; i < len(buah); i++ {
		fmt.Println("Element Ke - ", i, buah[i])
	}
}
