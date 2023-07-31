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
		"http://youtube.com",
		"http://golang.org",
	}

	c := make(chan string) // channel

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c { // infinit loop
		go func(link string) {
			time.Sleep(5 * time.Second) // pause one second
			checkLink(link, c)
		}(l) // function literal

	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
