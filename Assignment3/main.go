package main

import (
	"io"
	"os"
)

func readToStd(fname string) {
	mf, _ := os.Open(fname)
	io.Copy(os.Stdout, mf)
}

func main() {
	readToStd(os.Args[1])
}
