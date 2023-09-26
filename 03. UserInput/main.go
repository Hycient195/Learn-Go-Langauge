package main

import (
	"bufio"
	"os"
	"fmt"
)

func main ()  {
	fmt.Println("Welcome to user input");
	fmt.Println("Give the user a rating");

	reader := bufio.NewReader(os.Stdin);

	userInput, _ := reader.ReadString('\n');
	fmt.Print("User received a rating of ", userInput);
	fmt.Printf("Input is of type %T ", userInput);
}