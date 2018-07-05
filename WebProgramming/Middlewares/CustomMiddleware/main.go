/*
	The code is just demonstration of implementing a Middleware functionality in GoLang, we will write simple
	Middleware which will be called before and after executing actual handler method.

	In context of Java, you can think of Middleware as a ServletFilter, or AOP intercepter we write using Spring AOP
	framework.
*/

package main

import (
	"io"
	"log"
	"net/http"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Before: Log Request Attributes here you want for debugging purpose")
		log.Printf("Method: %s", req.Method)
		log.Printf("Remote Address: %s", req.RemoteAddr)
		next.ServeHTTP(w, req)
		log.Printf("After: Log Request/Response Attributes here you want for debugging purpose")
		log.Printf("Content-Type of the Response: %s", w.Header().Get("Content-Type"))
	}
}

func authMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Before: Authenticate")
		next.ServeHTTP(w, req)
		log.Printf("After: Authenticate")
	}
}

/*
	You can also create a chain of Middleware, below is the function of implementing chain of Middleware
	Below function has been done recursivily, the same can be implemented as an Iterative way.
*/
func chainMiddlewaresRecursive(f http.HandlerFunc, m ...middleware) http.HandlerFunc {
	// if chain is done, return original handlerfunc
	if len(m) == 0 {
		return f
	}
	// otherwise nest the handlerfuncs
	return m[0](chainMiddlewaresRecursive(f, m[1:cap(m)]...))
}

/*
	Chain middlewares using Iterative way
*/
func chainMiddlewaresIterative(f http.HandlerFunc, m ...middleware) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		last := f
		for i := len(m) - 1; i >= 0; i-- {
			last = m[i](last)
		}

		last(w, req)
	}
}

func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `Welcome to Home page!!`)
}
func audit(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `Logging activities!!`)
}

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {

	http.HandleFunc("/", authMiddleWare(home))
	http.HandleFunc("/log", loggingMiddleware(audit))

	// implementing chaining middleware recursive and Iterative
	http.HandleFunc("/middleware-recursive", chainMiddlewaresRecursive(home, authMiddleWare, loggingMiddleware))
	http.HandleFunc("/middleware-iterative", chainMiddlewaresIterative(home, authMiddleWare, loggingMiddleware))

	http.ListenAndServe(":8080", nil)
}
