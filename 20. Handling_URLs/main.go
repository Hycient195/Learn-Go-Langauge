package main;

import (
	"net/url"
	"fmt"
)

var testURl string = "https://example:3000/test-path?category=course&level=beginner";

func main() {
	fmt.Println("Handling URLs in Go Langauge");

	parsedURL, _ := url.Parse(testURl);

	// These different properties can be used to acces different parts of the URL;
	fmt.Println(parsedURL.Scheme);
	fmt.Println(parsedURL.Host);
	fmt.Println(parsedURL.Port());
	fmt.Println(parsedURL.Path);
	fmt.Println(parsedURL.RawQuery);
	fmt.Println(parsedURL.Query());

	// On the other hand, a URL can be composed up of its parts as shown below.
	var constructedURL *url.URL = &url.URL{
		Scheme: "https",
		Host: "fake-api",
		Path: "users",
		RawQuery: "age=19&region=Africa",
	}

	fmt.Println(constructedURL.String());

	/*
		It is good to know that the two main methods used to carry out the operations of
		getting parts of a URL, and also composing up a URL using its parts are the 
		"Parse" and "URL" methods from the "net/url" package.
	*/
}