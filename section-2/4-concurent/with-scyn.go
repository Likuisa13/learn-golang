package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func taskA(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai Task A")
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai Task A")
}

func taskB(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai Task B")
	time.Sleep(3 * time.Second)
	fmt.Println("Selesai Task B")
}

func main() {
	runtime.GOMAXPROCS(4)
	// runtime.GOMAXPROCS(runtime.NumCPU())

	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(2) // Menambahkan 2 goroutine yang akan ditunggu

	go taskA(&wg)
	go taskB(&wg)

	wg.Wait() // Menunggu semua goroutine selesai

	duration := time.Since(start)
	fmt.Println("Total waktu (dengan concurrency):", duration)
}
