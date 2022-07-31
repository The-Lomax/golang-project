package main

import "fmt"

func main() {
	employees := map[string]string{
		"tb": "Tim Burton",
		"jb": "Jack Black",
		"bw": "Bob White",
		"jd": "John Doe",
		"bm": "Bloody Mary",
	}
	printMap(employees)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
