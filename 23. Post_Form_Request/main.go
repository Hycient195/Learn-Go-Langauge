package main;

import (
	"io"
	"net/http"
	"net/url"
	"fmt"
)

func main() {
	fmt.Println("Making formData POST request in Go");

	var link string = "https://jsonplaceholder.typicode.com/posts";

	var formData = url.Values{}
	formData.Add("player", "Marques");
	formData.Add("sport", "Frisbee");
	formData.Add("isPlaying", "true");

	response, _ := http.PostForm(link, formData);

	dataBytes, _ := io.ReadAll(response.Body);
	defer response.Body.Close();
	
	fmt.Println(string(dataBytes));
}