package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}
var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer

func main() {
	fmt.Println("Hello from GoRoutines")

	// go greeter("Hello")
	// greeter("World")
	websitelist := []string{
		"https://rudrakshkanungo.xyz",
		"https://dribbble.com/tags/webiste",
		"https://github.com",
	}
	for _, website := range websitelist {
		go getStatusCode(website)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("its error")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
