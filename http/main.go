package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	//fmt.Println(resp)

	//condense down blow
	//   io.Copy(os.Stdout, resp.Body)

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

// 인터페이스는 계약이다!
func (logWriter) Write(bs []byte) (int, error) {
	//fmt.Println(string(bs))
	fmt.Println("사이즈 = ", len(bs))
	return len(bs), nil
}
