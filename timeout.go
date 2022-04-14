package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)

	message := make(chan int)
	go getData(message)
	sendData(message)

}

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func getData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Println("Terima Data - ", data)
		case <-time.After(time.Second * 5):
			fmt.Println("Error TimeOut, Tidak Ada Aktivitas Dalam Waktu Tertentu")
			break loop
		}

	}

}
