package main

func main() {
	var pelajar1 mahasiswa
	pelajar1.nama = "Hasyim Kipuy"
	pelajar1.nim = "211011401091"
	pelajar1.umur = 17
	pelajar1.kelas = "02TPLP015"
	pelajar1.methodTampil()
}

type mahasiswa struct {
	nama  string
	nim   string
	umur  int
	kelas string
}

func (m mahasiswa) methodTampil() {
	println(
		"nama  \t:", m.nama,
		"\nNIM \t:", m.nim,
		"\nUmur \t:", m.umur,
		"\nKelas \t:", m.kelas)
}
