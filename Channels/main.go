package main

import (
	"fmt"
	"net/http"
	"time"
)

type siteChecker struct {
	URLs []string
}

type siteResult struct {
	Site string
	error
}

func (s siteChecker) checkSite(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "   ==>   ", err)
		c <- site
		return
	}
	fmt.Println(site, "   ==>   Success")
	c <- site
}

func main() {
	checker := siteChecker{
		URLs: []string{
			"https://www.google.com",
			"https://www.facebook.com",
			"https://stackoverflow.com",
			"https://golang.org",
			"https://amazon.com",
			"https://github.com",
		},
	}

	c := make(chan string)

	for _, site := range checker.URLs {
		go checker.checkSite(site, c)
	}

	for u := range c {
		go func() {
			time.Sleep(time.Second * 3)
			checker.checkSite(u, c)
		}()
	}
}
