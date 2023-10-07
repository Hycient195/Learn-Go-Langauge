package main;

import (
	"fmt"
)

func main() {
	// Getting the memory location of a variable
	age := 30;
	agePtr := &age;
	fmt.Println("Location of age is ", &agePtr);

	// Obtaining value from a pointer
	ageVal := *agePtr
	fmt.Println("The real value of age is ", ageVal);

	// Storing a value address in a pointer
	var count *int = agePtr;
	fmt.Println("Counter is ", count);

	var score int = 120;
	var scorePtr *int = &score;
	fmt.Println("Score pointer is ", scorePtr)

	var name string = "Hycient";
	fmt.Println(name);
	modifyValue(&name);
	fmt.Println(name);

	var nameToAccess string = "Jennifer";
	accessValue(nameToAccess);
	accessValueLocaion(&nameToAccess);
}

func modifyValue(name *string) {
	*name = ("Hello " + *name);
}

// This function modifies its own copy of the "name" variable passed to it, and cannot overwrite the source of the value itself
func accessValue(name string)  {
	name = "Bob";
	fmt.Println(name)
}

// On the other hand this function has access to the "name" variable's memory address, and as such can modify the original value itself
// Such that even any other function making use of the original "name" variable, would be affected by this
func accessValueLocaion(name *string) {
	*name = "Laura";
	fmt.Println(*name);
}