package main

import "fmt"

/*
Go는 클로져(혹은 anonymous function이라고 부르는) 지원한다.
익명 함수(Anonymous function)는 코드 중간에 이름 없이 함수를 정의하기 위해서 사용한다.
이름이 없기 때문에 외부에서 사용 할 수 없다.

실용적인 관점에서는 데이터가 오직 하나의 메서드만 가지고 있을 때,
데이터가 정의된 함수 내에서 처리 하고자 할 때 사용한다.

데이터와 데이터를 조작하는 코드를 같은 코드 영역에 묶을 수 있다.
코드의 흐름을 깨지않으며, 가독성도 높아진다.
데이터의 처리에 두 개 이상의 함수가 사용 될 경우 익명 함수는 그리 좋은 선택이 아니다.
*/

func intSeq() func() int {//intSeq함수는 함수를 반환한다. 반환하는 함수는 intSeq 내부에서 익명 함수로 정의 했다. 변수 i는 익명 함수로 전달 된다.
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main() {
	nextInt := intSeq()
	/*
		27 줄에서 intSeq 함수의 리턴 값을 nextInt에 할당했다.
		이제 nextInt는 익명 함수를 가리키게 되고 이 후에는 함수처럼 사용 할 수 있다.
		38 번째 줄에서 intSeq()를 새로 할당해서 사용하고 있다.
	*/
    
    fmt.Println(nextInt())//nexInt를 함수처럼 사용하는 것을 볼 수 있다.
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    
    newInt := intSeq()
    fmt.Println(newInt())
}
		