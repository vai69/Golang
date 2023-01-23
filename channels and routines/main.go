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
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	c:= make(chan interface{})

	for _, link := range links {
		go checkLink(link,c)
	}

	for l := range c{
		go func(i interface{}){
			time.Sleep(time.Second*3)
			checkLink(fmt.Sprint(i),c)
		}(l)
	}
	
}

func checkLink(link string, c chan interface{}) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is down!")
		c <- link
		return
	}

	fmt.Println(link, "is Up!")

	c <- link
}
