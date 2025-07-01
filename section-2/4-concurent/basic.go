package main

import (
	"fmt"
	"time"
)

func getOngkirJNE(ch chan string) {
	time.Sleep(2 * time.Second) // simulasi delay API JNE
	ch <- "JNE: Rp12.000"
}

func getOngkirSiCepat(ch chan string) {
	time.Sleep(1 * time.Second) // simulasi delay API SiCepat
	ch <- "SiCepat: Rp10.500"
}

func getOngkirJNT(ch chan string) {
	time.Sleep(1 * time.Second) // simulasi delay API J&T

	// Simulasi error
	err := true
	if err {
		ch <- "J&T: ERROR saat mengambil ongkir"
		return
	}

	ch <- "J&T: Rp11.000"
}

func main() {
	chJNE := make(chan string)
	chSiCepat := make(chan string)
	chJNT := make(chan string)

	go getOngkirJNE(chJNE)
	go getOngkirSiCepat(chSiCepat)
	go getOngkirJNT(chJNT)

	ongkirJNE := <-chJNE
	ongkirSiCepat := <-chSiCepat
	ongkirJNT := <-chJNT // Menambahkan variabel untuk ongkir J&T

	fmt.Println("Ongkir JNE    :", ongkirJNE)
	fmt.Println("Ongkir SiCepat :", ongkirSiCepat)
	fmt.Println("Ongkir J&T     :", ongkirJNT) // Mencetak ongkir J&T
}
