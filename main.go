package main

// main 은 컴파일을 위해서 필요하다. 공유 라이브러리를 만들거면 main xx
// 익스텐션을 설치하면 자동 줄바꿈 자동 타입체크 등등 다해줌 겁나 편한가?
// 함수 이름의 시작이 소문자면 내부함수 대문자면 import 가능
// 가끔씩 편집기 창이 겁나 춤춤... 이상함 ㅠ
import (
	"fmt"
	"strings"
	"time"
)

func mul(a int, b int) int { // 각각각 타입을 지정해 줘야 함
	return a * b
}
func returnType(name string) (length int, uppercase string) { // legnth uppercase는 위에서 선언된것
	defer fmt.Print("\n imDone\n") // fnc가 끝나면 동작함
	length = len(name)
	uppercase = strings.ToUpper(name)
	return // 요런 신기한것도 지원함 ㄷㄷ 이걸 naked return 이라고 함
}
func main() {
	const name string = "yh" // 정적 타입 지정 언어
	myName := "leeyh"        // 축약형을 쓰면 첫번째로 지정한 타입으로 할당됨 지금은 str
	myName = "cococo"        // 타입을 지정하지 않으면 프프린트도 안댐
	fmt.Println("hello, world")
	fmt.Println("\n할로 월드?")
	fmt.Println(name + "\n\n")
	fmt.Println(myName)
	fmt.Println(mul(2, 3))
	fmt.Println(returnType("sass")) //뭔가 달고 주석추가 안하면 코딩창 지랄남?
	fmt.Println(time.Now())
	result := 0
	for i := 0; i < 1000000; i++ {
		result += i
	}
	fmt.Println(result)
	fmt.Println(time.Now())
	loop(1, 2, 3, 4, 5, 6, 7)
}

func loop(numbers ...int) int {

	for index, number := range numbers { // 이렇게 쓰면 foreach뮨

		fmt.Println(index, number)
	}

	for _, number := range numbers { // 이렇게 쓰면 foreach문인데 index ignore 가능!number ignore도 가능 ~

		fmt.Println(number)
	}
	fmt.Println("아래는 포문~")
	for i := 0; i < len(numbers); i++ { // 문법적으로 틀린게 있을때 ctrl + s 누르면 해결해줌 ㄷ
		fmt.Println(i)
	}
	return 1 // go에서는 for문 하나로 모든게 가능하다.
}
