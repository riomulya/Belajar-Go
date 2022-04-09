package main

import "fmt"

func main() {
	dataMahasiswa := map[string]uint{"Aldo": 182, "Yosep": 178}
	tampil("Aldo", dataMahasiswa["Aldo"])
	tampil("Yosep", dataMahasiswa["Yosep"])
}

func tampil(nama string, nilai uint) {
	fmt.Println(nama, ": ", nilai, " cm")
}
