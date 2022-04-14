package main

import (
	"fmt"
	"net/url"
)

func main() {

	urlText := "https://awangbali.wordpress.com/author/awangbali/page/34/?key=ekonomik%20&author=awang%20&tahun=2009"
	url, err := url.Parse(urlText)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Url        \t: ", urlText)
		fmt.Println("Host       \t: ", url.Host)
		fmt.Println("Path       \t: ", url.Path)

		key := url.Query()["key"][0]
		author := url.Query()["author"][0]
		tahun := url.Query()["tahun"][0]
		fmt.Println("Key-Nya    \t: ", key)
		fmt.Println("author-Nya \t: ", author)
		fmt.Println("tahun-Nya  \t: ", tahun)
	}

}
