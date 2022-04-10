package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	text := "data rahasia"
	s := sha1.New()
	s.Write([]byte(text))
	enkripsi := s.Sum(nil)
	fmt.Printf("Enkripsi : %x", enkripsi)
}
