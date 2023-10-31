package main

import (
	"sync"
	"fmt"
)

func main() {
	ch := make([]chan int, 3);
	wg := &sync.WaitGroup{};

	for i := range ch {
		ch[i] = make(chan int)
	}

	wg.Add(4);
	go sender(ch, 34, wg);

	go receiverOne(ch[0], wg);
	go receiverTwo(ch[1], wg);
	go receiverThree(ch[2], wg);

	/* Receiving the broadcast via a loop */
	// for index, val := range ch {
	// 	wg.Add(1)
	// 	go func(arg chan int, id int) {
	// 		res := <- arg;
	// 		fmt.Printf("Routine %d received value %v \n", id + 1, res);
	// 		wg.Done();
	// 	}(val, index)
	// }

	wg.Wait()
}

func sender(cha []chan int, data int, wg *sync.WaitGroup) {
	for _, val := range cha {
		val <- data
		defer close(val);
	}
	defer wg.Done();
}

func receiverOne(arg chan int, wg *sync.WaitGroup) {
	res := <- arg;
	fmt.Printf("Receiver 1 got value %v\n", res);
	wg.Done();
}

func receiverTwo(arg chan int, wg *sync.WaitGroup) {
	res := <- arg;
	fmt.Printf("Receiver 2 got value %v\n", res);
	wg.Done();
}

func receiverThree(arg chan int, wg *sync.WaitGroup) {
	res := <- arg;
	fmt.Printf("Receiver 3 got value %v\n", res);
	wg.Done();
}