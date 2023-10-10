package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exercise on functions");

	// A Basic function syntax;
	fmt.Println(addTwo(12, 23));

	// Function accepting multiple arguments (Variadic functions)
	var nums = []int{1,2,3,4,5}
	fmt.Println(addMultiple(nums...));

	// Function returning multiple values
	res1, res2, res3 := returnMultiple(2, 4)
	fmt.Printf("Addition is %v, multiplication is %v, third is %v\n", res1, res2, res3);

	// Anonymous instantly invoked function expression (IIFE)
	func(arg string) {
		fmt.Println("Hello ", arg)
	}("Jolene")

	// Assigning an anonymous function to a varible and invoking it from the variable.
	namedFuncRes := func() string{ 
		return "Hello there"
	}();

	fmt.Println(namedFuncRes)

	// Function named return values;
	fmt.Println(subAndAdd(10, 5));
}

// A simple function
func addTwo(arg1 int, arg2 int) int {
	return arg1 + arg2;
}

// A function taking multiple arguments
func addMultiple(args ...int, ) int {
	var result int = 0;
	for _, value := range args {
		result += value;
	}
	return result;
}

// Runction returning multiple values that can be used with the comma okay syntax;
func returnMultiple(arg1 int, arg2 int) (int, int, int) {
	return arg1 + arg2, arg1 * arg2, arg1 - arg2;
}

// Function named return values;
func subAndAdd(a, b int) (diff int, sum int) {
	diff = a - b;
	sum = a + b;
	return
}

// A runction taking multiple arguments and returning a new slice;
func addAnotherMultiple(args ...int) (res []int) {
	res = append(res, args...)
	return
}