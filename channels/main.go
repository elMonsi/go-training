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
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for {
	// 	go checkLink(<-c, c)
	// }
	for l := range c {
		go func(link string) { // we receive the value
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // here we pass the value to the function literal
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error with website:", link)
		c <- link
	} else {
		fmt.Println("Website is OK:", link)
		c <- link
	}
}
