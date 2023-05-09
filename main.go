package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	val, err := os.Open(args[1])

	if err != nil {
		fmt.Print("Error")
	}

	io.Copy(os.Stdout, val)
	fmt.Print(args)
}
