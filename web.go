package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello World!</h1>")
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h1>Halaman Utama</h1>")
	})
	http.HandleFunc("/index", index)
	fmt.Println("Memulai Web Server Pada Port : 8080")
	http.ListenAndServe(":8080", nil)
}
