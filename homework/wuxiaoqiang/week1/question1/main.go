package main

import (
	"fmt"
	"sync"
)

// 两个协程交替打印数字1-100
var wg sync.WaitGroup

func main() {
	fmt.Println("beginning of main")
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go printNumbers()
	}
	wg.Wait()

	fmt.Println("end of main")
}

func printNumbers() {
	fmt.Println("beginning of print")
	for j := 0; j < 100; j++ {
		fmt.Println(j)
	}
	fmt.Println("end of print")
	wg.Done()
}
