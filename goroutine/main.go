package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	list := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
		"http://golang.org",
	}

	c := make(chan string)
	for _, link := range list {
		go checklink(link, c)
	}

	for l := range c {
		go func(x string) {
			time.Sleep(5 * time.Second)
			go checklink(x, c)
		}(l)
	}
}

func checklink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}
