package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://www.facebook.com/",
		"http://stackoverflow.com",
		"http://go.dev",
		"http://amazon.com",
	}

	//create channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "🔴")
		//c <- link + "🔴"
		c <- link
		return
	}
	fmt.Println(link, "🟢")
	//c <- link + " 🟢"
	c <- link
}
