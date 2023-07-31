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

## 환경설정 에러

- 아래와 같은 에러 메시지가 출력되지만, 정상적으로 실행되었습니다

```
GOPATH is an environment variable used by the Go programming language to determine the location of the Go workspace on your system. The Go workspace is a directory hierarchy where Go source code and its dependencies are stored.

In the past, the GOPATH environment variable was commonly used to specify the root of the Go workspace. It could contain multiple directories, such as bin, src, and pkg, which served different purposes. The src directory typically contained the source code of Go packages, pkg stored compiled package objects, and bin held executable binaries.

However, since Go version 1.11, the Go module system was introduced, which changed the way dependencies are managed and eliminated the need for a fixed GOPATH. With modules, you can organize your Go code in any directory, and the Go tooling automatically manages dependencies based on the go.mod file present in the project.

Nowadays, it's recommended to use Go modules and avoid relying on the GOPATH environment variable for most projects. You can initialize a new Go module in a directory outside of the GOPATH by using the go mod init command.

In summary, while GOPATH used to be an important environment variable in older versions of Go, it has become less relevant with the introduction of the Go module system. It's recommended to use Go modules and let the Go tooling handle dependencies without relying on a fixed GOPATH.
```

```
go work init
```

- 그리고 `cmd + ,` 키 눌러서 go settings.json 파일에 추가

```json
{
  "..." : "..."
  "gopls": {
    "experimentalWorkspaceModule": true
  }
}
```

- 아 모듈문제인거같은데 정확히 모르겠음

- https://github.com/golang/tools/blob/master/gopls/doc/workspace.md << 이거보고 해결함

## 카드 미니 프로젝트

### 셔플기능

- RAND 패키지 : 공식문서 >>

```
func Intn(n int) int
```

## 고 모듈

```
go.mod 파일 생성
다가오는 강의에서 우리는 첫 번째 테스트를 실행할 것입니다. 이렇게 하려면 go.mod 파일을 만들어야 합니다. 그렇지 않으면 다음과 같은 오류가 표시됩니다.

go: go.mod 파일이 현재 디렉토리나 상위 디렉토리에 없습니다. '이동 도움말 모듈' 참조

카드 프로젝트 디렉토리 내에서 다음을 실행합니다.

go mod init cards

go test그런 다음 VSCode 내에서 테스트 실행 기능을 사용하거나 터미널에서 실행할 수 있습니다 .
```

- 내가쓴거

```
go mod init cards

go: creating new go.mod: module cards
go: to add module requirements and sums:

          go mod tidy
```

## Errorf 사용에 관하여

```
Errorf 호출에 인수가 있지만 형식화 지시문은 없습니다.
다음 강의에서는 Errorf를 사용하여 몇 가지 테스트를 실행할 것입니다 . 서식 지시문을 생략하면 이제 테스트가 실패하므로 즉시 추가해야 합니다.

참고 - 이 지시문은 원래 강의 끝에 추가되었습니다.

8:00 타임스탬프부터 시작하여 테스트를 실행할 때 다음을 변경합니다.

t.Errorf("Expected deck length of 16, but got", len(d))

이에:

t.Errorf("Expected deck length of 16, but got %v", len(d))
```

### 고에서의 테스트

```
원본파일 : deck.go
테스트파일 : deck_test.go

테스트파일 규칙 : {원본파일명}_test.go
```

- 테스트 실행 명령어 : `go test`

### 나의 궁금함

```
1) t *testing.T  << 이거뭐임.. 포인터야?
2) assert 쓰고싶다
3) 모듈 패키지 개념 정리 필요해
4) 리시버 개념
5) 리턴파라미터 몇개까지 가능?
```

### go Struct (구조체)

## go 포인터

- 이코드가 동작하지 않는 원인은?

```go
func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName // 동작을 안함! 미친 !
}
```

## 포인터 배우기 전에 RAM

```
jim := person ...
jimmy := person ...

Address  Value
0000
0001      jim
0002
0003      jimmy
0004
```

- 패스바이 벨류 : 방어적 복사가 기본인 언어

### 포인터 개념과 연산자

- 엠퍼센드 '&' : 메모리 주소를 주는 연산자
- 스타 '\*' : 그 주소에 있는 메모리들의 값을 내놔라
- 코드보면서 설명

```go
func main() {
	jimmy := jim()
	jimmyPointer := &jimmy
	fmt.Print("\n\n")
	fmt.Println("안되야한다")
	jimmy.updateFirstName("goodspeed")
	jimmy.print()
	fmt.Println("되야한다")
	jimmyPointer.updateFirstNamePointer("goodMan")
	jimmy.print()
}

func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName // 동작을 안함! 미친 !
}

// *person : 타입을 설명하는건데, person 형식의 포인터로 동작한다는 타입을 설명해주는거임
func (personPointer *person) updateFirstNamePointer(newFirstName string) {
	(*personPointer).firstName = newFirstName

  // (*personPointer) : 이건 연산자에요, 이 포인트가 가르치는 메모리 주소에 담겨있는 값을 가져와라
  // 포인터에 담겨있는
}
```

### 포인터

```
personPointer = *person
```

```
personPointer := &jim
```

### 포인터 관점에서 Array 와 Slices

#### 배경설명

- 슬라이스

- 어레이

- 슬라이스는 일종의 어레이다 : 3가지 부분으로 나눠짐
  - 포인터(ptr to head)
  - 사이즈(lenght)
  - 어쩌고(Cap) >> 데이터의 어레이가 저장되는것?

#### 서론2 : 데이터타입

- 벨류타입 (자바에서의 프리미티브 타입)
  - int, float, bool, string, structs ....
- 레퍼런스 타입
  - slices, map, channels, 등등..
  - 여전히 데

#### 본론

- 슬라이스 데이터타입은 레퍼런스타입이다

### 고랭에서의 Map 자료구조

- K-V 구조임 그건 여전히 맞음

### 맵

- 레퍼런스타입!!!
- 키는 같은 타입
- 벨류는 같은타입
- 키는 인덱스되어있음
-

### 구조체(struct)

- 벨류타입!!!
- 인덱싱을 지원하지않는다
- 서로 다른 다양한 타입들이 들어있을수 있다
- 구조체를 넘기면 구조체는 복사한다 왜냐하면 >> 벨류타입이라
  - 맵을 넘기며면 복사하지않음. 왜냐하면 레퍼런스타입이라 참조만 넘기기 때문
- 모든 필드에 대한 you need know all field a디프런트 필드 앳 컴파일타임 ": >> 참조되서 돌아다니기때문에 여기저기서 접근이 가능하다
- 유스 리프리젠스 ㅌ어 띵! 윗 어 랏오브 ㅇ디프런트 프로퍼티스
-

## 인터페이스 개념 인 고랭

### 인터페이스가 필요한 이유

- 지금까지의 코드에서의 문제점
  - 모든 값들은 타입이 존재한다
  - 모든 펑션들은 정해진 타입으로만 주고받는다 : 아규먼트(입력출력), 리시버 등... 타입이 확실해야만한다
  - 결론 : 로직이 정말 같은데, 타입이 다르면 재사용이 불가능한가??

### 인터페이스의 역할

### 인터페이스에 대한 규칙

### 자바에서의 인터페이스와의 차이점, 고랭의 인터페이스

- 함수 타입만 맞으면 뭐든 호출해서 사용 : \
- 그래서 인터페이스를 쓰면 뭐가 좋은걸까??
- 나쁜 코드를 만든다음 >> 리팩토링 하는 과정으로 설명

  - 인터페이스 기반으로 리팩토링!

- 결론 : 인터페이스는 메시지를 전송한다!!!!

### 고랭의 인터페이스

#### 콘크리트 타입

- 맵, 스트럭트, 인트, 스트링 등등.. >> 변화할수 없는 타입

#### 인터페이스 타입

- 실체화할수 없고, 꼭 이 형식으로 구체클래스(콘크리트 클래스 타입)를 작성해줘야하는

#### 인터페이스 배경설명, 고랭의 인터페이스 설명

- 인터페에스는 제네릭타입이 아니다! : 고랭은 제네릭타입을 지원하지 않음. 앞으로 지원할꺼긴 한데..

  - 고랭은 제네릭 지원안하는걸루 유명하지,.,

- 암묵적 인터페이스를 지원한다 : 자바에서는 명시적인 인터페이스만 존재했는데 >

  - bot 인터페이스, englishbot 구체타입이랑 명시적으로 꼭 작성해줄필요는 없다
  - 축복일수도 있고 저주일수도 있다.

- 인터페이스는 contract(계약) 혹은 타입관리자(help us manage types) 이다

  - 리턴타입과 입력타입이 맞는지만 점검해주는,
  - 함수(계약) 을 잘 지켰는지만 확인해주는거다

- interface are tough : 인터페이스는 터프하다
  - 이해해라 스탠다드 라이브러리같은거 읽어 보면서 인터페이스가 어떻게 동작하는지
  - 너만의 인터페이스를 작성하는건 어려울수 있다

#### 인터페이스의 고랭 구현
