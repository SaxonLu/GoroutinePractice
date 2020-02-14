package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("start")
	wg.Add(1)
	go InsertA(&wg)
	wg.Add(1)
	go InsertB(&wg)
	wg.Wait()
	fmt.Println("end")
}

func InsertA(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	for i := 0; i < 100; i++ {
		fmt.Println("AAA")
	}
	wg.Done()
}

func InsertB(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	for i := 0; i < 100; i++ {
		fmt.Println("BBB")
	}
	wg.Done()
}
