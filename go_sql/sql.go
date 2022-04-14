package main

import (
	"database/sql"
	"fmt"
	_ "mysql-master"
)

type pelajar struct {
	id    int
	nama  string
	kelas int
	umur  int
}

func koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_sql")
	if err != nil {
		return nil, err
	}
	return db, nil
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

	var result []pelajar

	for rows.Next() {
		var each = pelajar{}
		err := rows.Scan(&each.id, &each.nama, &each.kelas, &each.umur)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, each := range result {
		fmt.Println("Id    \t:", each.id)
		fmt.Println("nama  \t:", each.nama)
		fmt.Println("kelas \t:", each.kelas)
		fmt.Println("umur  \t:", each.umur)
		fmt.Println("\n==========================")
	}
}

func sql_tambah() {
	db, err := koneksi()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into tbl_data values (null,?,?,?)", "Hasyim Kipuy", 22, 3)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Tambah Data Berhasil")
}

func sql_delete(nama string) {
	db, err := koneksi()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("delete from tbl_data where nama = ", &nama)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("\nDelete Data Berhasil")
}

func sql_ubahKelas(nama string, kelas int) {
	db, err := koneksi()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("update from tbl_data set kelas = ? where nama = ? ", &kelas, &nama)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("\nUbah Kelas Berhasil")
}

func sql_ubahUmur(nama string, umur int) {
	db, err := koneksi()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("update from tbl_data set umur = ? where nama = ? ", &umur, &nama)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("\nUbah umur Berhasil")
}
func sql_ubahNama(nama string, ubahNama string) {
	db, err := koneksi()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("update from tbl_data  where nama = ? set nama = ?", &ubahNama, &nama)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("\nUbah Nama Berhasil")
}

func main() {
	var pilih int
loop:
	for i := 0; true; i++ {
		fmt.Printf("Klik 1 untuk menampilkan data\nklik 2 untuk delete data\nklik 3 untuk merubah data nama\nklik 4 untuk merubah data kelas\nklik 5 untuk merubah data umur : ")
		fmt.Scanln(&pilih)
		if pilih == 1 {
			sql_tampil()
		} else if pilih == 2 {
			var deleteNama string
			fmt.Print("Nama Yang Akan Anda Delete Datanya :")
			fmt.Scanln(&deleteNama)
			sql_delete(deleteNama)
		} else if pilih == 3 {
			var nama, ubahNama string
			fmt.Print("Nama Yang Akan Anda Ubah Datanya :")
			fmt.Scanln(&nama)
			fmt.Print("Menjadi : ")
			fmt.Scanln(&ubahNama)
			sql_ubahNama(nama, ubahNama)
		} else if pilih == 4 {
			var nama string
			var ubahKelas int
			fmt.Print("Nama Yang Akan Anda Ubah Datanya :")
			fmt.Scanln(&nama)
			fmt.Print("Menjadi Kelas : ")
			fmt.Scanln(&ubahKelas)
			sql_ubahKelas(nama, ubahKelas)
		} else if pilih == 5 {
			var nama string
			var ubahUmur int
			fmt.Print("Nama Yang Akan Anda Ubah Datanya :")
			fmt.Scanln(&nama)
			fmt.Print("Menjadi Kelas : ")
			fmt.Scanln(&ubahUmur)
			sql_ubahUmur(nama, ubahUmur)
		} else {
			break loop
		}
	}
}
