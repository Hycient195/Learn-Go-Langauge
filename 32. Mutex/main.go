package main;

import (
	"fmt"
	"sync"
)

var Balance int = 200;
var mu sync.Mutex;
var wg sync.WaitGroup;

func main() {
	println("Getting started with mutex");

	wg.Add(1)
	go deposit();

	wg.Add(1);
	go withdraw();

	wg.Wait();
	fmt.Printf("Balance after transactions is %d", Balance);
}

func deposit() {
	for i := 0; i < 5000; i++ {
		mu.Lock();
		Balance += 10;
		mu.Unlock();
	}
	println("Deposit completed");
	wg.Done();
}

func withdraw() {
	for i := 0; i < 5000; i++ {
		mu.Lock();
		Balance -= 10;
		mu.Unlock();
	}
	println("Withdrawal completed");
	wg.Done();
}
