package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLinkStatus(link, c)
	}

	for l := range c {

		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLinkStatus(link, c)
		}(l)
		//checkLinkStatus(l, c)
	}

}

func checkLinkStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!!", err)
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link

}
