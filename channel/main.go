package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://amazon.com",
		"http://amazon.com",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://golang.org",
		"http://facebook.com",
		"http://google.com",
		"http://google.com",
		"http://kakaoenterprise.com",
	}

	// for _, link := range links {
	//// v1 : 이거는 반응이 없음
	// 	go checkLink(link)
	// }

	// for _, link := range links {
	//// v2 : 이거는 그냥 동기적으로 동작함
	// 	go fmt.Println(checkLinkV2(link))
	// }

	c := make(chan string)
	for _, link := range links {
		go checkLinkChannel(link, c)
		//fmt.Println(<-c) // 이런식으로 포문안에 넣으면 블로킹으로 동작함
	}
	//for i, _ := range links {
	for i := 0; i < len(links); i++ {
		fmt.Print(i)
		fmt.Println(<-c) // 체널에서 데이터를 받는건 블러킹이다
	}

}

func checkLinkChannel(link string, c chan string) {
	_, myErr := http.Get(link)

	if myErr != nil {
		c <- link + "cannot connect"
		return
	}
	c <- link + " is GOOD!"
}

func checkLink(link string) {
	_, myErr := http.Get(link)

	if myErr != nil {
		fmt.Println(link, "cannot connect")
		return
	}
	println(link + " is GOOD!")
}

func checkLinkV2(link string) string {
	_, myErr := http.Get(link)

	if myErr != nil {
		return link + "cannot connect"
	}
	return link + " is GOOD!"
}
