package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 123, 23444, 657}

	for _, elem := range a {
		if elem%2 == 0 {
			fmt.Printf("%d is even\n", elem)
		}
		if elem%2 == 1 {
			fmt.Printf("%d is odd\n", elem)
		}
	}
}
