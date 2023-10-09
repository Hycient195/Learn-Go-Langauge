package main;

import (
	"time"
	"math/rand"
	"fmt"
)

func main() {
	fmt.Println("Working with switch statement");


	// Basis Syntax of 
	rand.NewSource(time.Now().UnixNano())
	randNum := rand.Intn(6) + 1;

	fmt.Println(randNum);

	switch randNum {
	case 1:
		fmt.Println("Random number is 1"); 
	case 2:
		fmt.Println("Random number is 2");
		fallthrough;
	case 3:
		fmt.Println("Random number is 3");
		fallthrough;
	case 4:
		fmt.Println("Random number is 4"); 
	case 5:
		fmt.Println("Random number is 5"); 
	case 6:
		fmt.Println("Random number is 6"); 
	default:
		fmt.Println("Number is invalid");
	}
}