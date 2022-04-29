package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "mysql-master"
	"net/http"
)

type pelajar struct {
	id    int
	nama  string
	kelas int
	umur  int
}

var data []pelajar

func koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_sql")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	sql_tampil()
	http.HandleFunc("/data", ambilData)
	http.HandleFunc("/cariData", cariData)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h1>Halaman Utama</h1>")
	})
	http.HandleFunc("/index", index)
	fmt.Println("Memulai Web Server Pada Port : 8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello World!</h1>")
}

func ambilData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func cariData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		namaData := r.FormValue("nama")
		for _, d := range data {
			if d.nama == namaData {
				result, err := json.Marshal(d)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				w.Write(result)
				return
			}
		}
		http.Error(w, "Data Pelajar Tidak Tersedia", http.StatusBadRequest)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func sql_tampil() {
	db, err := koneksi()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from tbl_data")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := pelajar{}
		err := rows.Scan(&p.id, &p.nama, &p.kelas, &p.umur)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, p)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
