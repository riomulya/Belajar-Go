package main

import (
	"fmt"
	"os"
)

func main() {
	err := fmt.Errorf(" Program Berakhir ")
	// Printing the error message
	defer fmt.Println(err.Error())

	fmt.Println("Jalankan")
	os.Exit(10)

}
