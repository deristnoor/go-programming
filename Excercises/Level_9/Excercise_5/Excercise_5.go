package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var increment int64
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&increment, 1)
			fmt.Println(atomic.LoadInt64(&increment))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("finish:", increment)
}
