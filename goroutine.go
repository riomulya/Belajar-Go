package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	// tampilPesan(10, "KOSONG")
	go tampilPesan(100, "spam")
	tampilPesan(100, "Bukan spam")

	var input string
	fmt.Scanln(&input)
}

func tampilPesan(ulang int, pesan string) {
	for i := 0; i < ulang; i++ {
		fmt.Println(i+1, pesan)
	}
}
