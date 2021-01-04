package main

import (
	"fmt"
	"math"
)

const s string = "constant"

/*상수는 const를 이용해서 선언 할 수 있다.
상수(常數)는 어느 관계를 통하여 변하지 않는 일정한 값을 가진 수나 양을 의미한다.
코드내에서 변수처럼 사용 할 수는 있지만 값을 변경 할 수는 없다.
값을 변경 하려고 하면 컴파일 에러가 발생한다.
30번째 줄에 n = 1000 코드를 추가하고 실행해보자.
컴파일 에러를 출력할 것이다.
*/
func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	fmt.Println(math.Cos(n))
	//	const n = 1000
	fmt.Println(math.Tanh(n))
}

/*
컴파일러는 상수를 사용 할 때,
사용하는 데이터 타입에 맞게 타입을 변환한다.
위 코드에서 n은 integer형 상수이지만 math.Sin에서 사용 할 때는 float64로 자동으로 변환된다.
*/
