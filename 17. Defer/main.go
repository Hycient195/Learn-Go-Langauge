package main;

import (
	"fmt"
)

func main() {
	fmt.Println("Studying the 'Defer' Keyword");

	defer fmt.Println("Third")
	defer fmt.Println("Second")
	defer fmt.Println("First")
	fmt.Println("Welcome")
	PrintNumbers()
}

func PrintNumbers() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}