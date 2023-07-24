package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	dong := person{"김", "동훈"}
	ranhee := person{firstName: "son", lastName: "rani"}

	fmt.Println(dong)
	fmt.Println(ranhee)
}
