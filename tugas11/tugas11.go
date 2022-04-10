package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	text1 := "1232"
	text2 := "132"

	bil1, err1 := strconv.Atoi(text1)
	bil2, err2 := strconv.Atoi(text2)

	if err1 == nil && err2 == nil {
		jumlah := bil1 + bil2
		fmt.Println("Jumlah : ", jumlah)
		time.Sleep(time.Second * 2)

		kurang := bil1 - bil2
		fmt.Println("kurang : ", kurang)
		time.Sleep(time.Second * 2)

		bagi := bil1 / bil2
		fmt.Println("bagi : ", bagi)
		time.Sleep(time.Second * 2)

		kali := bil1 * bil2
		fmt.Println("kali : ", kali)

	} else {
		panic(err1.Error())
	}
}
