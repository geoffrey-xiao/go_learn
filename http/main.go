package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		fmt.Println("Error is", err)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

// customer writer

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("just wrote this many byte", len(string(bs)))
	return len(bs), nil
}
