package main

import "fmt"

func main() {
	var arr []int
	for i := 0; i <= 10; i++ {
		arr = append(arr, i)
	}
	for _, el := range arr {
		if el%2 == 0 {
			fmt.Println(el, "is even")
		} else {
			fmt.Println(el, "is odd")
		}
	}
}
