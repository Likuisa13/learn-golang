package main

import "fmt"

const statusOK = 200
const statusFailed = 500

const (
	statusNotFound   = 404
	sttusServerError = 500
)

func main() {
	fmt.Println("OK: ", statusOK)
	fmt.Println("Not Found: ", statusNotFound)
	fmt.Println("Server Error: ", sttusServerError)
}
