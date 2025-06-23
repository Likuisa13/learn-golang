package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{}
	fmt.Printf("myslice1=%v\n", myslice1)
	fmt.Printf("length=%d\n", len(myslice1))

	fmt.Println("--------------------")

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Printf("myslice2=%v\n", myslice2)
	fmt.Printf("length=%d\n", len(myslice2))

	fmt.Println("--------------------")

	mysliceA := []int{1, 2, 3}
	mysliceB := []int{4, 5, 6}
	myslice3 := append(mysliceA, mysliceB...)

	fmt.Printf("myslice3=%v\n", myslice3)
	fmt.Printf("length=%d\n", len(myslice3))
}
