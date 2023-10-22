package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://baidu.com",
		"https://google.com",
		"https://amazon.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	for l := range c {
		// time.Sleep(time.Second * 5)
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		// c <- "it is down"
		c <- link
		return
	}
	fmt.Println(link, "is up")
	// c <- "it is up"
	c <- link
}
