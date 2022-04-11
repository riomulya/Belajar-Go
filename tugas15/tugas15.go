package main

import (
	"fmt"
	"net/http"
)

func ulang(w http.ResponseWriter, r *http.Request) {
	for i := 1; i <= 100; i++ {
		for j := 1; j <= i; j++ {
			fmt.Fprint(w, "# ")
		}
		fmt.Fprint(w, "\n")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello World!</h1>")
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h1>Halaman Utama</h1>")
	})
	http.HandleFunc("/ulang", ulang)
	http.HandleFunc("/index", index)
	fmt.Println("Memulai Web Server Pada Port : 8080")
	http.ListenAndServe(":8080", nil)
}
