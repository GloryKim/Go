package main
// 문자열을 숫자로 변환하는 것은 기본적인 작업이지만, 많은 프로그래머들이 어려워하는 작업이기도 하다. Go를 이용해서 문자열을 숫자로 바꿔보자.
import (
    "strconv"
    "fmt"
)

func main() {
    f, _:= strconv.ParseFloat("1.234", 64)
    fmt.Println(f)
    
    i, _:= strconv.ParseInt("123", 0, 64)
    fmt.Println(i)
    
    d, _:= strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)
    
    u, _:= strconv.ParseUint("789", 0, 64)
    fmt.Println(u)
    
    k, _:= strconv.Atoi("135")
    fmt.Println(k)
    
    _, e:= strconv.Atoi("wat")
    fmt.Println(e)
}

/*
strconv 패키지로 문자열을 숫자로 변환할 수 있다. ParseFloat는 문자열을 float타입으로 인식해서 변환한다. ParseInt는 integer 타입으로, ParseUint는 unsigned int로 변환한다. ParseInt는 Hex 타입도 인식한다. 각 메서드의 두번째 매개변수는 몇 비트를 변환할지를 결정한다.
Atoi는 문자열을 interger 타입으로 변환한다. ParseInt와 비슷한일을 하지만, 더 간단하게 사용 할 수 있다.
인식할 수 없는 문자열에 대해서는 error를 반환한다.
*/
		