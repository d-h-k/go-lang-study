package main

import (
	"fmt"
	"unsafe"
)

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// 이거랑 :	contactInfo   contactInfo
	// 이거랑 같음 : contactInfo
}

type contactInfo struct {
	eamil   string
	zipCode int
}

func main() {

	dongContact := contactInfo{"email", 123}
	dong := person{"김", "동훈", dongContact}
	fmt.Println(dong)

	//rContact := contactInfo{"email",456}
	ranhee := person{lastName: "rani", firstName: "son"}
	fmt.Println("순서 상관없이 이름 지정해서 할당할수 있다 : ", ranhee)

	var zeroValue person // 제로벨류 할당됨
	alexContact := contactInfo{"email", 123}
	alex := person{"man", "All", alexContact}
	fmt.Println("제로벨류가 할당되며, null 이런거 없다. 0, \"\" 같은 제로벨류가 들어감 : ", zeroValue)
	fmt.Printf("printf 출력용법2 : %+v \n", zeroValue)
	fmt.Printf("printf 출력용법2 : %+v \n", alex)

	jimmy := jim()
	jimmyPointer := &jimmy
	fmt.Print("\n\n")
	fmt.Println("안되야한다")
	jimmy.updateFirstName("goodspeed")
	jimmy.print()
	fmt.Println("되야한다")
	jimmyPointer.updateFirstNamePointer("goodMan")
	jimmy.print()

	// unsafe
	fmt.Println("Size of 'jimmy' variable:", unsafe.Sizeof(jimmy))
	fmt.Println("Size of 'jimmyPointer' variable:", unsafe.Sizeof(jimmyPointer))

}

func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName // 동작을 안함! 미친 !
}

func (personPointer *person) updateFirstNamePointer(newFirstName string) {
	(*personPointer).firstName = newFirstName
}

func jim() person {
	jim := person{
		firstName: "Jim",
		lastName:  "Tonic",
		contact:   contactInfo{"email", 123},
	}
	fmt.Printf("%+v", jim)
	return jim
}

func (p person) print() {
	// 리시버를 사용해봅니다
	fmt.Printf("%+v\n", p)
}
