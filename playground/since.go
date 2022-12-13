package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now().Unix()
	time.Sleep(2 * time.Second)
	fmt.Printf("%v\n", float64(time.Since(time.Unix(nowTime, 0)))/1000000000)
}
