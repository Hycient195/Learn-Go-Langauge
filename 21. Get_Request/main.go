package main;

import (
	"strings"
	"io"
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("Getting Started with making GET requests");
	MakeGetRequest();
}

func MakeGetRequest() {
	// Making the request to the API endpoint
	response, _ := http.Get("https://jsonplaceholder.typicode.com/todos/1");

	// Obtaining the bytes of data from the response body
	dataBytes, _ := io.ReadAll(response.Body);
	defer response.Body.Close(); // To finally close the network connection after the request has been made.

	// Converting the response bytes to string
	var responseString = string(dataBytes);

	// Alternate method of converting the databytes recieved to string.
	var stringBuilder strings.Builder;
	byteCount, _ := stringBuilder.Write(dataBytes);
	var responseString2 string = stringBuilder.String();
	fmt.Println("Response byte count is ", byteCount)

	// Printing the response in the terminal
	fmt.Println(responseString) 
	fmt.Println(responseString2)
}