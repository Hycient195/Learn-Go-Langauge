package main;

import (
	"io"
	"net/http"
	"strings"
	"fmt"
)

func main() {
	fmt.Println("Making a POST request in Go");

	var url string = "https://jsonplaceholder.typicode.com/posts";

	var formData = strings.NewReader(`{
		"name": "Hycient",
		"age": 16,
		"isOnline": false
	}`);

	response, _ := http.Post(url, "application/json", formData);

	dataBytes, _ := io.ReadAll(response.Body);
	defer response.Body.Close();


	fmt.Println(string(dataBytes));
}