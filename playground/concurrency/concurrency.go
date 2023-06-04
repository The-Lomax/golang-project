package main

import (
	"fmt"
	"time"
)

func returnAfterDelay(sec float64, c1 chan float64) {
	for i := 0.0; i < sec; i = i + 0.5 {
		time.Sleep(time.Duration(1) * time.Second / 2)
		fmt.Printf("loop step: %v\n", i)
		c1 <- i
	}
}

func main() {
	c1 := make(chan float64, 10)
	go returnAfterDelay(3.0, c1)
	fmt.Println("Hello, playground")
	time.Sleep(time.Duration(2) * time.Second)
	slice := []float64{}

loop:
	for {
		select {
		case v, ok := <-c1:
			slice = append(slice, v)
			fmt.Printf("Received: %v, more data in channel: %v\n", v, ok)
		default:
			fmt.Println("No value received")
			break loop
		}
	}
	fmt.Printf("Slice: %v\n", slice)
}
