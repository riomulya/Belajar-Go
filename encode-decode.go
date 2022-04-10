package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	namaSaya := "Rio Mulya Syawal"

	encodeString := base64.StdEncoding.EncodeToString([]byte(namaSaya))
	fmt.Println(encodeString)
	decodeString, _ := base64.StdEncoding.DecodeString(encodeString)
	fmt.Println(string(decodeString))
}
