package main;

import (
	"fmt"
)

func main() {
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug"}

	// First syntax of using the for loop to access itens iteratively using their index
	for i := 0; i < len(months); i++ {
		// fmt.Println(months[i]);
	}

	// for i := range months {
	// 	// fmt.Println(months[i])
	// }

	for i, month := range months {
		fmt.Printf("Month %v has index of %v \n", month, i);
	}

	// Using break statement
	for i := 0; i < len(months); i++ {
		if (i == 4) {
			break; // execution breaks out of the loop if the condition above is satisfied.
		}
		// fmt.Println(months[i]);
	}

	// Using Continue Statement
	for i := 0; i < len(months); i++ {
		if (i == 4) {
			continue; // execution skips the block of code in the for loop if the condition above is satisfied.
		}
		// fmt.Println(months[i]);
	}

	// Using the "goto" statement
	for i := 0; i < len(months); i++ {
		if (i == 4) {
			goto breakout; // execution breaks out of the loop if the condition above is satisfied.
		}
		// fmt.Println(months[i]);
	}

	breakout:
		fmt.Println("This is a breakout statement");

	// Using the for loop like a while loop in other langauges;
	count := 0;
	for count < 10 {
		if (count == 5) {
			count ++;
			continue;
		}
		fmt.Printf("Count is %v \n", count);
		count ++;
	}


} 