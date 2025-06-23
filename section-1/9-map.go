package main

import "fmt"

func main() {
	// Deklarasi map dengan properti key string dan value int
	ages := map[string]int{
		"Bambang": 25,
		"Siti":    30,
	}

	fmt.Println("--------------------")
	fmt.Println("Initial map:", ages)
	fmt.Println("--------------------")

	// Mengakses elemen map
	fmt.Println("Umur Bambang:", ages["Bambang"])

	// Menambahkan dan mengubah elemen map
	ages["Romi"] = 28
	fmt.Println("Map after adding Romi:", ages)

	// Mengubah nilai
	ages["Bambang"] = 26
	fmt.Println("Age of Bambang:", ages["Bambang"])
	fmt.Println("--------------------")

	// Menghapus elemen map
	delete(ages, "Siti")
	fmt.Println("Map after deleting Siti:", ages)
	fmt.Println("--------------------")

	// Menggunakan make
	mhsages := make(map[string]int)
	mhsages["Putra"] = 22
	mhsages["Dewi"] = 23
	fmt.Println("Map after creating Putra:", mhsages)
	fmt.Println("--------------------")

	// Mengecek apakah key ada
	age, exists := ages["Putra"]
	if exists {
		fmt.Printf("Putra's age exists and is: %d\n", age)
	} else {
		fmt.Println("Putra not found")
	}
	fmt.Println("--------------------")

	// Iterasi Map
	fmt.Println("Iterating over map:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
	fmt.Println("--------------------")

	// Map bersarang
	users := map[string]map[string]interface{}{
		"user1": {
			"name": "Budi",
			"age":  20,
		},
		"user2": {
			"name": "Putra",
			"age":  30,
		},
	}

	fmt.Println(users)

	for id, user := range users {
		fmt.Println("ID:", id, "Name:", user["name"])
	}
}
