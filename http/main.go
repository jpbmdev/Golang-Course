package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	return len(bs), nil
}

func main() {
	res, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	bs := make([]byte, 99999)

	res.Body.Read(bs)

	fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, res.Body)
}
