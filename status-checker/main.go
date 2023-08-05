package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://facebook.com",
		"http://google.com",
		"http://kakaoenterprise.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLinkChannel(link, c)
	}

	for l := range c {
		//v3
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLinkChannel(link, c)
		}(l)

		//v1
		//go checkLinkChannel(<-c, c)

		//v2
		//time.Sleep(2 * time.Second)
		//go checkLinkChannel(l, c)

	}

}

func checkLinkChannel(link string, c chan string) {
	_, myErr := http.Get(link)
	if myErr != nil {
		fmt.Println(link + " is DOWN")
		c <- link
		return
	}
	fmt.Println(link + " is UP")
	c <- link
}
