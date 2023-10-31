// package main;

// import (
// 	"time"
// 	"fmt"
// )

// func main() {
// 	go spinner(100 * time.Millisecond)
// 	const n = 45
// 	fibN := fib(n) // slow
// 	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
// }
// func spinner(delay time.Duration) {
// 	for {
// 		for _, r := range `-\|/` {
// 			fmt.Printf("\r%c", r)
// 			time.Sleep(delay)
// 		}
// 	}
// }

// func fib(x int) int {
// 	if x < 2 {
// 		return x
// 	}
// 	return fib(x-1) + fib(x-2)
// }

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Working with select statement in Concurrent Go Applications")

	chs := make([]chan string, 3)
	wg := &sync.WaitGroup{}

	for i := range chs {
		chs[i] = make(chan string, 1)
	}

	wg.Add(3)
	go sender1(chs[0], wg)
	go sender2(chs[1], wg)
	go sender3(chs[2], wg)

	var Globval string

	for i := range chs {
		select {
		case msg := <-chs[i]:
			Globval = msg
			chs[i] <- "Response from sender" + fmt.Sprint(i+1)
		}
	}

	wg.Wait()

	fmt.Println(Globval)
}

func sender1(ch chan string, wg *sync.WaitGroup) {
	ch <- "From channel one"
	msg := <-ch
	fmt.Printf("%s: %s\n", ch, msg)
	wg.Done()
}

func sender2(ch chan string, wg *sync.WaitGroup) {
	ch <- "From channel two"
	msg := <-ch
	fmt.Printf("%s: %s\n", ch, msg)
	wg.Done()
}

func sender3(ch chan string, wg *sync.WaitGroup) {
	ch <- "From channel three"
	msg := <-ch
	fmt.Printf("%s: %s\n", ch, msg)
	wg.Done()
}
