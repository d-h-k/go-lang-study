# go-lang-study

go-lang-study

## helloworld 예제에서 배울점

```go
package main
//

import "fmt"
// fmt
// 임포트는 꼭 더블쿼트 " 문자를 써야함

func main() {
	fmt.Println("hello world!")
}
```

- 코드 실행법 : go run main.go

## 주요 명령어

```
go run : 알아서 컴파일하고 실행
go build : 컴파일만 하고 실행하지는 않음. 실행직전까지 빌드만 함
go fmt :
go install :
go test :
```

## package 개념

- package : 여러개의 기능을 여러개의 파일, 하나의 폴더 안에 담아두는것이 패키지임
  - package main >> main.go, support.go, subtask.go
  - 두가지 타입이 있음 : Excutable, Reusable(디펜던시나, 라이브러리로 만들 수 있음)
  - Excutable 패키지 만드는법
    - package main >> main.go 파일이 있어서 >> go build 를 하면 이게 그냥 실행가능한거임
  - Reusable 패키지 만드는법
    - package main >> a기능.go, b기능.go >> 패키지 이름과 파일이름이 드라면 리유저블 패키지임
  - 결론 : 패키지 이름은 매우 중요하다

## import 개념

- import "fmt" : 터미널 stdio.h 같은 거임
- 이외 고랭의 패키지들
  - debug, main, fmt, io, crypto, encoding, math
- 패키지 공식내용

## func 개념

- 함수임.. ㅎ..
- 로직넣어라

## 고랭 코드 구성

```go
// 1 패키지
package AAA

// 2 임포트
import "BBB"

// 3 함수
func CCC () {

}
```

## 정적타입언어

- 다이내믹 타입

  - 자바스크립ㅌ, 루비, 파이썬
  - 변수에 타입이 없음

- ## 스테틱 타입(정적타입)

## 자료구조

- Array
  - 길이가 고정된 리스트
- Slice
  - 가변길이 리스트 (grow, shrink 가능)
  - 모든 원소들의 타입이 같아야함
