package main

import (
	"context"
	"fmt"
	"net/http"
)

type ctxKey string

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/print-ctx", print)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	u := ctxKey("userName")
	r := ctxKey("role")
	// Use Context values for request-scoped data only, not for passing optional parameters to functions
	ctx = context.WithValue(ctx, u, "Deepak")
	ctx = context.WithValue(ctx, r, "Admin")

	ctxValue := acessCtxData(ctx, u)

	useofBackgroundContext(w)
	fmt.Println("Use of Request Scoped Context")
	fmt.Fprintln(w, ctxValue)
}

func print(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	fmt.Fprintln(w, ctx)
}

func acessCtxData(ctx context.Context, k ctxKey) string {
	userName := ctx.Value(k)
	//fmt.Println("userName")
	return userName.(string)
}

func useofBackgroundContext(w http.ResponseWriter) {
	language := ctxKey("language")
	ctx := context.WithValue(context.Background(), language, "GoLang")
	fmt.Println("Use of Background Context")
	fmt.Fprintln(w, ctx)
	fmt.Fprintf(w, "Programming Language: %s\n", ctx.Value(language))
}
