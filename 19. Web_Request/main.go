package main;

import (
	"io"
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("Gettin started with making webreqiests in Go Langauge");

	var URL string = "https://html-resume-silk.vercel.app";
	// Making a GET request to the provided URL
	response, err := http.Get(URL);
	HandleError(err);
	defer response.Body.Close();

	// The response contains so many parameters byt we are only interested in reading content from it's body.
	dateBytes, err := io.ReadAll(response.Body);
	HandleError(err);

	// Since the response body being read is in the form of bytes, it has to be converted to string so that it can be readable
	responseString := string(dateBytes);


	fmt.Println(responseString);

	fmt.Println(response.Header["Cache-Control"]);

	fmt.Printf("Type of response is %T\n", response);

	
	// fmt.Println(response.Body);
}

func HandleError(err error) {
	if err != nil {
		panic(err);
	}
}