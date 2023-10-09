package main;

import (
	"fmt"
)

func main() {
	fmt.Println("Working with structs");

	// Creating an instance of the "Player" struct.
	var Hycient Player = Player{"Hycient", 10, true, []string{"Readiing", "Swimming", "Adventures"}}
	fmt.Printf("%+v \n", Hycient)

	// You can also initialize a struct using the syntax below by specifying the name of the properties
	var John Player = Player{
		Name: "John",
		Age: 17,
		IsRegistered: false,
		Hobbies: []string{"Jumping", "Frying", "Roasting"},
	}
	fmt.Printf("%+v \n", John);

	// Accessing a particular value in an instanc of a struct using the dot "." notation
	fmt.Printf("John's hobbies are %v \n\n", John.Hobbies);
	John.Hobbies = []string{"Hiking", "SkyDiving"};
	fmt.Printf("John's new hobbies are %v \n\n", John.Hobbies);

	// You can also initialize a struct using the syntax below
	var newUser User = User{
		Name: "Chidera Onugwu",
		IsActive: true,
		Account: struct{
			Number int64
			Bank string
		}{
			Number: 2523525,
			Bank: "ECO",
		},
	}
	fmt.Printf("%+v \n", newUser)
}

// Delcaring my Struc
type Player struct {
	Name string
	Age int
	IsRegistered bool
	Hobbies []string
}

type User struct {
	Name string
	IsActive bool
	Account struct {
		Number int64
		Bank string
	}
}