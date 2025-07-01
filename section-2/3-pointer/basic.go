package main

import "fmt"

func main() {
	tanpaPointer := 2000
	var denganPointer *int
	denganPointer = &tanpaPointer

	fmt.Println("Nilai tanpa pointer: ", tanpaPointer)
	fmt.Println("Nilai dengan pointer: ", denganPointer)
	fmt.Println("Akses nilai dengan pointer: ", *denganPointer)

	luas, err := hitungLuasPersegiPanjang(10, 20)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Luas persegi panjang: ", *luas)
}

func hitungLuasPersegiPanjang(panjang, lebar int) (*int, error) {
	if panjang <= 0 || lebar <= 0 {
		return nil, fmt.Errorf("PANJANG ATAU LEBAR TIDAK BOLEH KURANG DARI 0")
	}

	luas := panjang * lebar
	return &luas, nil
}
