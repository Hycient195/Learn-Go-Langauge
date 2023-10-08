package main;

import (
	"fmt"
)
func main()  {
	fmt.Println("Learning about Arrays");

	// Array Declaration
	var cities [3]string;
	// Assigning values to the different index positions of the array;
	cities[0] = "Cairo";
	cities[2] = "Cotonou"
	fmt.Println(cities);
	fmt.Println(len(cities));

	// Declaring and initializing an array on the go;
	var scores = [4]int{23, 56, 23, 56};
	fmt.Println(scores);

	var equalScore = [4]int{23, 56, 23, 56};
	fmt.Println(scores == equalScore);

	// Passing an array and the pointer to an array to a function to see their respective behaviours.
	printArray(scores);
	fmt.Println(scores);

	mutateArray(&equalScore);
	fmt.Println(equalScore);
}

func printArray(arg [4]int)  {
	arg[2] = 100;
	fmt.Println(arg);
}

func mutateArray(arg *[4]int) {
	(*arg)[2] = 200;
	fmt.Println(*arg);
}