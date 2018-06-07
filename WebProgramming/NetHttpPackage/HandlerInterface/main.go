package main

import (
	"fmt"
	"net/http"
)

// The concept is basically any type which implement ServeHTTP method will be treatd as a Handler
type handlerInterface int

func main() {
	var handler handlerInterface
	http.ListenAndServe(":8080", handler)
}

func (handler handlerInterface) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving http request")
}
