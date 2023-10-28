/* ============================================= */
/* Passing Wait Groups As Arguments To Functions */
/* ============================================= */

package main;

import (
	"sync"
	"net/http"
	"fmt"
)

func main() {
	var wg sync.WaitGroup;
	var URLS []string = []string { "https://google.com", "https://hycient.vercel.app", "https://scrabble-modern.onrender.com", "https://facebook.com", "https://github.com" };

	for _, url := range URLS {
		wg.Add(1)
		go checkWebStatus(url, &wg);
	}
	wg.Wait();
}

func checkWebStatus(url string, wg *sync.WaitGroup) {
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