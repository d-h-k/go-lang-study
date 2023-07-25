package main

import "fmt"

func main() {
	myslice := []string{"Hi", "1", "2", "3"}
	updateSlice(myslice)

}

func updateSlice(myslice []string) {
	myslice[0] = "bye"
}

func testBefore() {
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
