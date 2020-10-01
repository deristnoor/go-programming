package main

import (
	"fmt"
	"sync"
)

func main() {
	increment := 0
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := increment

			v++
			increment = v
			fmt.Println(increment)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("finish:", increment)
}
