package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	s := os.Args
	if len(s) < 2 {
		fmt.Println("file name not provided")
		os.Exit(1)
	}

	f, err := os.Open(s[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
