package main

import (
	"fmt"
	"io"
	"os"

	"./imgcat"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "missing paths to image")
		os.Exit(2)
	}

	for _, path := range os.Args {
		err := cat(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to cat given %s: %v", path, err)
			panic(err)
		}
	}

}

func cat(path string) error {
	f, err := os.Open(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open a given File for the given path %s", path)
		return err
	}
	defer f.Close()
	wc := imgcat.NewWriter(os.Stdout)
	if _, err = io.Copy(wc, f); err != nil {
		return err
	}
	return wc.Close()
}
