package main

// main 은 컴파일을 위해서 필요하다. 공유 라이브러리를 만들거면 main xx
// 익스텐션을 설치하면 자동 줄바꿈 자동 타입체크 등등 다해줌 겁나 편한가?
// 함수 이름의 시작이 소문자면 내부함수 대문자면 import 가능
// 가끔씩 편집기 창이 겁나 춤춤... 이상함 ㅠ
import (
	"fmt"
	"mydict"
	"strings"
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
	dictionary := mydict.Dictionary{}
	dictionary["super"] = "monster"
	addErr := dictionary.Add("word", "description")
	if addErr == nil {
		fmt.Println(dictionary)
	}
	updateErr := dictionary.Update("word", "메롱")
	if updateErr == nil {
		fmt.Println(dictionary["word"])
	}
	dictionary.Delete("word")
	fmt.Println(dictionary)
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
