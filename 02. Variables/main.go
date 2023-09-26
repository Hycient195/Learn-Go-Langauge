/* Variables */
/* 
	Variables are simply containers or placeholders to hold specific type
	of data for future manipulation, computation or reference in a program.
*/

package main

import "fmt";

// No var style is not allowed outside function body or method
// randomNumber := 323.34;
// fmt.Println(randomNumber);

const IPAddress = "234.238492.2093";
var GlobalUrl = "https://hycient.vercel.app";

func main() {
	var username string = "John Doe";
	fmt.Println("This is user " + username);
	fmt.Printf("Username is of type %T \n", username);

	var isLoggedIn bool = false;
	fmt.Printf("Type of %T \n", isLoggedIn);

	var age int = 257;
	fmt.Printf("Age is of type %T \n", age);

	var temperature float64 = 56.0902378238787686;
	fmt.Println("Temperature is", temperature);
	fmt.Printf("Temperature is of type %T \n", temperature);

	// Default values for initializing a variable withough assigning any value to it
	var emptyNum float32;
	fmt.Println(emptyNum);

	var emptyString string;
	fmt.Println(emptyString);

	// Implicit type
	var playerName = "Yoshi \n";
	fmt.Println(playerName);

	// No var style
	numberOfPlayers := 3;
	fmt.Println("This is the number of players", numberOfPlayers);
	fmt.Printf("Type of number of players is %T \n", numberOfPlayers)
	// It is worthy to note that the no-var style of variable declaration cannot be used outside of a method or function scope;

	// Accessing and printing a Global constant
	fmt.Println("The global IP Address is", IPAddress);
	// Constants cannot be re-assigned. The below commented code would throw an error
	// IPAddress = "New IP"

	// Accessing and printing a global variable;
	fmt.Println("The global URL is: ", GlobalUrl);

	const localIpAddress = "9842.234.3424";
	fmt.Println("The local IP address is", localIpAddress);
	fmt.Printf("The local IP address is of type %T \n", localIpAddress);
	
}