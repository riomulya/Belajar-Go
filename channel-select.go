package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	go spamData(100, "Beliau Ini Kocak Geming")
	number := []int{12, 321, 112, 321, 23, 123, 667, 1}

	fmt.Println("Isi Slice : ", number)

	channel1 := make(chan int)
	go getMax(number, channel1)

	channel2 := make(chan float64)
	go getRata2(number, channel2)

	for i := 0; i < 2; i++ {
		select {
		case rata2 := <-channel2:
			fmt.Printf("Rata Ratanya : %.2f\n", rata2)
		case max := <-channel1:
			fmt.Printf("Nilai Maksimalnya : %d\n", max)
		}
	}
}

func getRata2(nilai []int, ch chan float64) {
	sum := 0
	for _, e := range nilai {
		sum += e
	}
	ch <- float64(sum) / float64(len(nilai))
}

func getMax(nilai []int, ch chan int) {
	max := 0
	for _, e := range nilai {
		if max < e {
			max = e
		}
	}
	ch <- max
}

func spamData(banyak int, pesan string) {
	for i := 0; i < banyak; i++ {
		fmt.Println(pesan)
	}
}
