package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("os.Args = ", os.Args)
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)

	//실행법
	fmt.Println("go run main.go my.txt")
	fmt.Println("go build main.go && ./main main.go")
}
