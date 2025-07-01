package main

import "fmt"

type Student struct {
	Name  string
	Age   *int // Field Age adalah pointer ke int
	Major string
}

func main() {
	studentA := Student{
		Name:  "Dwiki Likuisa",
		Major: "Teknik Informatika",
	}

	fmt.Println(studentA)

	studentAge := 30
	studentA.Age = &studentAge // Mengisi field Age dengan alamat memori dari studentAge
	fmt.Println(*studentA.Age) // Mengakses nilai yang ditunjuk oleh pointer
}
