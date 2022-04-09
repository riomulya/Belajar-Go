package main

import "fmt"

func main() {
	rata2 := hitung(12, 312, 23, 45, 67, 678, 9, 34, 456, 32, 1211)
	pesan := fmt.Sprintf("Rata rata nya : %.2f", rata2)
	fmt.Println(pesan)
}

func hitung(nilai ...int) float64 {
	total := 0

	for _, value := range nilai {
		total += value
		fmt.Println(value)
	}
	avg := float64(total) / float64(len(nilai))
	return avg
}
