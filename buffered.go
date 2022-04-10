package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	pesan := make(chan int, 3)

	go func() {
		for {
			i := <-pesan
			fmt.Println("Terima Data : ", i)
		}
	}()

	for i := 1; i <= 10; i++ {
		fmt.Println("Kirim Data :", i)
		pesan <- i
	}
}
