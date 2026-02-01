package main

import (
	"fmt"
	"net/http"
)

import "rsc.io/quote"
func helloHandler(w http.ResponseWriter, r *http.Request) {

	resultCh := make(chan int64)
	

	go func() {
		var sum int64 = 0
		for i := int64(0); i < 1_000_000_000; i++ {
			sum += i
		}
		resultCh <- sum
	}()

	sum := <-resultCh
	fmt.Fprintln(w, "Hello", sum)
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println(quote.Go(),"Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
