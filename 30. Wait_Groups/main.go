package main;

import (
	"sync"
	"net/http"
	"fmt"
)

var wg sync.WaitGroup;
var URLS []string = []string { "https://google.com", "https://hycient.vercel.app", "https://scrabble-modern.onrender.com", "https://facebook.com", "https://github.com" };

func main() {
	for _, url := range URLS {
		wg.Add(1)
		go checkWebStatus(url);
	}
	wg.Wait();
}

func checkWebStatus(url string) {
	resp, err := http.Get(url);
	HandlePanicError(err);
	defer resp.Body.Close();
	fmt.Printf("%v is status for %v\n", resp.Status, url);
	defer wg.Done();
}

func HandlePanicError(e error) {
	if e != nil {
		panic(e);
	}
}