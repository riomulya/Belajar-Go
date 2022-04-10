package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	namaSaya := "Rio Mulya Syawal"

	encodeString := base64.StdEncoding.EncodeToString([]byte(namaSaya))
	fmt.Println(encodeString)

	// text := "data rahasia"
	s := sha1.New()
	s.Write([]byte(encodeString))
	enkripsi := s.Sum(nil)
	fmt.Printf("Enkripsi : %X", enkripsi)
}
