/*
Go는 두 개 이상의 값을 반환할 수 있다. 이 점을 이용해서 실행결과와 에러 값을 함께 반환하는 방식으로 에러를 처리 할 수 있다. 이 방식은 상당히 직관적이다. 단지 하나의 값만을 반환하는 다른 언어들은 에러를 측정하는게 명확하지 않을 수 있다. 예를 들어 int 값을 반환하는 C 함수가 있다고 가정해 보자. 이 함수는 실패 할 경우 0을 반환한다. 그런데, -1을 반환하는 함수도 있다. 메뉴얼(man) 페이지를 보지 않는 한 개발자가 값 만을 가지고 에러인지 판단하기는 쉽지 않다.

*/


package main

import "errors"
import "fmt"

func f1(arg int) (int, error) {
    if arg == 42 {

        return -1, errors.New("can't work with 42")

    }

    return arg + 3, nil
}

type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {

        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {

    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}

/*
f1 함수는 두 개의 값을 반환한다. 이 중 두번째 값으로 error를 반환하고 있다. 에러일 경우에는 errors.New 메서드를 이용해서 새로운 에러를 만들고, 에러가 없을 경우에는 nil을 반환한다. 개발자는 error이 nil인지 아닌지를 확인하는 것으로 에러를 처리 할 수 있다. 에러 메시지는 error.Error() 메서드로 확인 할 수 있다.
error는 내장 인터페이스다. 따라서 개발자는 Error()메서드를 구현하는 방식으로 자신만의 errors타입을 만들 수 있다. 예제 코드에서는 argError이라는 구조체를 만들고 Error()를 구현했다. 여기에 에러의 타입을 설정하기 위해서 arg 변수도 추가했다.
f2함수는 내장 error 대신에, argError를 사용하고 있다. 28줄에서 &argError를 이용해서 새로운 구조체를 만들어서 반환했다.
main 함수에서는 f1와 f2함수를 테스트하면서, 어떻게 에러를 처리하는지를 보여주고 있다. if 문을 이용해서 두 번째 값이 nil인지를 판단해서 에러처리를 하고 있다.
57줄을 보자. error인터페이스로 부터 구조체타입을 읽어서 처리하는 방법을 보여주고 있다.

*/