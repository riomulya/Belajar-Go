package main

import (
	"fmt"
	"runtime"
)

var pesanChannel = make(chan string)

func main() {
	runtime.GOMAXPROCS(3)

	go tampilPesan("aku")
	go tampilPesan("siapa")
	go tampilPesan("ya")

	var message1 = <-pesanChannel
	fmt.Println(message1)
	var message2 = <-pesanChannel
	fmt.Println(message2)
	var message3 = <-pesanChannel
	fmt.Println(message3)

}

func tampilPesan(nama string) {
	data := fmt.Sprintf("HAllO %s", nama)
	pesanChannel <- data
}
