package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("There was an error:", err)
	} else {
		io.Copy(os.Stdout, file)
	}
}
