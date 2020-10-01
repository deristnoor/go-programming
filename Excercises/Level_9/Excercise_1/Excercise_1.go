package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Begin CPUs\t\t", runtime.NumCPU())
	fmt.Println("Begin Goroutines\t", runtime.NumGoroutine())

	wg.Add(2)
	go func() {
		fmt.Println("haha")
		wg.Done()
	}()

	go func() {
		fmt.Println("you do not know the answer")
		wg.Done()
	}()

	fmt.Println("Mid CPUs\t\t", runtime.NumCPU())
	fmt.Println("Mid Goroutines\t\t", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("nothing to see here")
	fmt.Println("End CPUs\t\t", runtime.NumCPU())
	fmt.Println("End Goroutines\t\t", runtime.NumGoroutine())
}
