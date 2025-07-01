package main

import (
	"errors"
	"fmt"
)

func validateAge(age int) error {
	if age < 0 {
		return errors.New("umur tidak boleh negatif")
	}
	if age > 70 {
		return errors.New("umur tidak boleh lebih dari 70")
	}
	return nil
}

func main() {
	umur := 25
	err := validateAge(umur)
	if err != nil {
		fmt.Printf("Error untuk umur %d: %s\n", umur, err.Error())
	} else {
		fmt.Printf("Umur %d valid.\n", umur)
	}
}
