package main;

import (
	"sync"
	"fmt"
)

func main() {
	fmt.Println("Working with select statement in Concurrent Go Applications");

	chs := make([]chan string, 3);
	wg := &sync.WaitGroup{};

	for i := range chs {
		chs[i] = make(chan string, 1);
	}

	wg.Add(3);
	go sender1(chs[0], wg);
	go sender2(chs[1], wg);
	go sender3(chs[2], wg);

	wg.Wait();

	var Globval string;

	select{
	case Globval = <-chs[0]:
		fmt.Printf("Variable is %v\n", Globval);
	case Globval = <-chs[1]:
	case Globval = <-chs[2]:
	default:
		fmt.Println("Waiting for channels");
	}

	fmt.Println(Globval);
}

func sender1(ch chan string, wg *sync.WaitGroup) {
	ch <- "From channel one";
	close(ch);
	wg.Done()
}

func sender2(ch chan string, wg *sync.WaitGroup) {
	ch <- "From channel two";
	close(ch);
	defer wg.Done();
}

func sender3(ch chan string, wg *sync.WaitGroup) {
	ch <- "Fron channel three";
	close(ch);
	defer wg.Done();
}