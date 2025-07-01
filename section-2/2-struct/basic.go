package main

import "fmt"

type Mahasiswa struct {
	Nama    string
	Umur    uint8
	Jurusan string
}

func main() {
	mahasiswa := Mahasiswa{
		Nama:    "Dwiki",
		Umur:    25,
		Jurusan: "Informatika",
	}

	mahasiswa.Jurusan = "Ilmu Komputer"
	fmt.Println(mahasiswa.Nama)
}
