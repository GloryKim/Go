package main

import "fmt"

//Go에서 변수는 명시적으로 선언되며, 컴파일시간에 정확히 선언했는지를 검사하게 된다.
//선언하고 사용하지 않은 변수들에 대해서는 경고가 아닌 에러를 발생하며, 컴파일 실패한다.

func main() {

	//C/C++ 언어와 비교해서 선언 방법에 차이가 있다.
	//C 언어의 경우 int a 와 같이 타입이 먼저 오는데 반해,
	//go는 이름이 먼저오고 그 다음에 타입이 온다.
	var a string = "initial"
	fmt.Println(a)
	//타입을 검사하는 자바 같은 언어에서도 타입이 앞에 오기 때문에 약간은 생소 할 수 있다.
	//이러한 선언 방식에 대해서는 인간의 언어에 좀 더 가깝게 하기 위해서라는 의견이 있다.
	//예컨데 "var a string"은 "변수 a는 스트링이다"라는 식으로 해석 할 수 있기 때문에 좀 더 직관적이라는 주장이다.
	//var a string를 풀어 쓰면, Variable a is string 이 되니 그럴듯 하다고 볼 수 있겠다.

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
	//19줄에서는 타입을 생략하고 있다.
	//:= 를 이용하면 선언을 짧게 할 수 있다.
	//풀어 쓰자면 var f string = "short" 가 된다.
}
