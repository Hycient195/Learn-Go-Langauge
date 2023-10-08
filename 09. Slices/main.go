package main;

import (
	"fmt"
)

func main() {
	fmt.Println("Learning about slices");

	// Declaring and initializing a slice;
	cars := []string{"Toyota", "Bentley", "Mercedes"};
	fmt.Println(cars);

	cars = append(cars, "Kia", "Honda");
	fmt.Println(cars);

	// Second method of declaring and intializing an array using the "make" keyword;
	points := make([]int, 5);
	points[0] = 23;
	points[1] = 89;
	points = append(points, 78, 90);
	fmt.Println(points);
	fmt.Println(len(points));
	fmt.Println(cap(points));
	/*
		The advantage of this approach is that under the hood, the slice has a fixed size at initialization like an array, but unlike an array,
		values can still be appended to it if need be, and the size would then be changed. Since it is fixed at initialization like an array, when,
		values are appended to it, the slice is moved out of its memory location to a new memory location that can now acccomodate it's new size, and the
		space allocation for it's previous address is garbage collected. This approach of creating a slice is better because it takes into account
		memory management, and moves the memory location to a new address if need be, on the event of appending a new value to the slice.
	 a specific size is allocate 
	*/

	// Mutability of slices;
	trees := []string{"Acacia", "Iroko", "Bamboo"};
	modifySlice(trees);
	fmt.Println(trees);

	// Slicing a slice (This also works for arrays also)
	numbers := []int{10,20,30,40,50,60};
	fmt.Println(numbers[2:4]) // This returns the subset of the <trees> array, start index included, and end index excluded.


	/* Copying Slices */
	initialSlice := []int{1,2,3,4,5};
	copySclice := make([]int, len(initialSlice));
	copy(copySclice, initialSlice);
	fmt.Println(copySclice)

	/* ====================================== */
	/* Removing an item from a slice by index */
	/* ====================================== */
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July"}
	removeFromSlice(months, 2);
	fmt.Println(months);
}

func modifySlice(arg []string) {
	arg[1] = "Mahogony";
	fmt.Println(arg);
}

func removeFromSlice(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...);
}