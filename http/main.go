package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//fmt.Println(resp)
	// bs := make([]byte, 99999)

	// resp.Body.Read(bs)

	//fmt.Println(string(bs))

	// Another approach
	//io.Copy(os.Stdout, resp.Body)

	//Custom Writer example.

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Written :", len(bs), "bytes")
	return len(bs), nil
	//return 1, nil
}
