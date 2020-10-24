package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("3 + 1 = ", 3+1)
	fmt.Println("2.0 / 3.0 = ", 2.0/3.0)
	fmt.Println("6.0 * 3.0 = ", 6.0*3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}

//Go는 strings, integers, floats, booleans 등 다양한 값(value)들을 지원한다.
//예제를 통해서 기본적인 값들을 살펴보도록 하자.
//+ 연산자를 이용해서 스트링을 더 할 수 있다.
//Integer과 floats 값들의 연산을 확인 할 수 있다.
//Boolean 연산의 경우 예상한 대로 결과가 나오는 걸 확인 할 수 있다.

//컴파일 방법 : F5클릭 후 'Select Debug'에서 'Go' 선택후 'DEBUG CONSOLE'창에서 'hello world'출력 확인
