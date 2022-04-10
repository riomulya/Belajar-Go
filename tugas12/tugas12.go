package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "jeruk Pisang durian Angur1"
	regex, err := regexp.Compile(`[A-Z]?[a-z]+[0-9]?`)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		hasil := regex.FindAllString(text, -1)
		fmt.Println(hasil)
		index := regex.FindStringIndex(text)
		fmt.Println(index)
	}
}
