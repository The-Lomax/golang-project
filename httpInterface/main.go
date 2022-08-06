package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}
