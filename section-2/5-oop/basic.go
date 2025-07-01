package main

import "fmt"

// Class/Object Account
type Account struct {
	nama     string // private (unexported)
	Username string // public (exported)
}

// Method untuk set nama (encapsulated access)
func (u *Account) SetNama(n string) {
	u.nama = n
}

// Method untuk ambil nama (encapsulated read)
func (u Account) GetNama() string {
	return u.nama
}

func main() {
	u := Account{Username: "dwikil"}
	u.SetNama("Dwiki Likuisa")

	fmt.Println("Username:", u.Username)
	fmt.Println("Nama Lengkap:", u.GetNama()) // akses via method
	// fmt.Println(u.nama) // error kalau beda package (karena 'nama' private)
}
