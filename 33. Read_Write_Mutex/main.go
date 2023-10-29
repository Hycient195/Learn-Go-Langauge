package main;

import (
	"fmt"
	"sync"
)

var (
	count       int = 0
	rwmu sync.RWMutex
	wg				 sync.WaitGroup
)

func main() {

	// Reading count (concurrent)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			rwmu.RLock()
			fmt.Println("Read count:", count)
			rwmu.RUnlock()
			wg.Done();
		}()
	}


	// Writing count
	rwmu.Lock()
	count = 42
	rwmu.Unlock()

	// Wait for goroutines to finish
	wg.Wait();
}