package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("Beginning of main function");

	go printHello()
	go greetUser();
	go helloSecondUser()

	time.Sleep(2 * time.Second); // This delay function pauses execution for 2 seconds for all the goroutines to execute, before exiting the main function;

	fmt.Println("End of main function");
}

func printHello() {
	fmt.Println("Hello People");
}

func greetUser() {
	fmt.Println("Hello new user");
}

func helloSecondUser() {
	fmt.Println("Greetings to the second user")
}