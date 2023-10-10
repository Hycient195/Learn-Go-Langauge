package main;

import (
	"io"
	"os"
	"fmt"
)

func main() {
	fmt.Println("File Operations in Go Langauge");

	// Create the file and handle file creation error case;
	file, err := os.Create("./file.txt");
	CheckNilError(err);

	// Write content to the file.
	size, err := io.WriteString(file, "This is a basic string writting into the text file.");
	CheckNilError(err);
	fmt.Println(size);

	// Reading content storen in the text file
	text, err := os.ReadFile("./file.txt");
	CheckNilError(err);

	// Displaying in the console, the content obtained from the text file;
	fmt.Println(string(text));

	// Close the file open when all operations have been performed on the file.
	defer file.Close();

	// err := os.Remove("./file.txt")
	// CheckNilError(err);
}

func CheckNilError(err error) {
	if err != nil {
		panic(err);
	}
}