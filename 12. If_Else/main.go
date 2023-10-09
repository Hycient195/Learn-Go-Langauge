package main

import (
	"fmt"
)

func main() {
	fmt.Println("Working with if and else statements");

	num := 34;

	if (num < 20) {
		fmt.Println("Numeber is less than 20");
	} else {
		fmt.Println("Number is not less than 20");
	}

	// Declaring, initializing and checking a vale;
	if count := 3; count > 3 {
		fmt.Println("Count is greater than 3")
	} else {
		fmt.Println("Count is not greater than 3");
	}
}