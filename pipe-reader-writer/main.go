package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//Use of Pipe
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		_, err := fmt.Fprintln(pw, "Wrting data using Pipe!!")
		if err != nil {
			panic(err.Error())
		}
	}()

	_, err := io.Copy(os.Stdout, pr)
	if err != nil {
		panic(err.Error())
	}

	// Use of MultiReader API
	start := strings.NewReader("<msg>")
	body := strings.NewReader("Use of MultiReader API!!")
	end := strings.NewReader("</msg>")
	r := io.MultiReader(start, body, end)
	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		panic(err.Error())
	}

	// Use of MultiWriterAPI

	buf := new(bytes.Buffer)
	mw := io.MultiWriter(os.Stdout, os.Stderr, buf)
	fmt.Fprintln(mw, "Writing to Multiple Writer!!")
	fmt.Println("From buffer: ", buf)
}
