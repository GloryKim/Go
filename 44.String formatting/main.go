package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {
    p := point{1, 2}
    fmt.Printf("%v\n", p)
    fmt.Printf("%+v\n", p)
    fmt.Printf("%#v\n", p)
    
    fmt.Printf("%T\n", p)
    fmt.Printf("%t\n", true)
    fmt.Printf("%d\n", 123)
    
    fmt.Printf("%b\n", 14)
    fmt.Printf("%c\n", 33)
    
    fmt.Printf("%x\n", 456)
    fmt.Printf("%f\n", 78.9)
    
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)
    
    fmt.Printf("%s\n", "\"String\"")
    fmt.Printf("%q\n", "\"String\"")
    fmt.Printf("%x\n", "hex this")
    
    fmt.Printf("%p\n", &p)
    fmt.Printf("|%6d|%6d|\n", 12, 345)
    
    fmt.Printf("|%-6.2f|%-6.2f|", 1.2, 3.45)
    fmt.Printf("|%6s|%6s|", "foo", "b")
    
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)
    
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}

//fmt패키지에서 제공하는 Printf메서드를 이용해서, 형식화된 출력을 할 수 있다.


		/*
fmt 패키지의 형식화된 출력에는 디버깅 목적으로 사용 할 수 있는 여러가지 출력 옵션들을 가지고 있다.
%v : 구조체의 값을 출력한다.
%+ek. : 구조체의 필드이름과 값을 출력한다.
%#g : 호출된 함수이름과 구조체의 이름 같은 소스코드 정보까지 함께 출력한다.
%T : 타입을 출력한다.
%t : 불리언의 값을 true 혹은 false 문자열로 출력한다.
%d : Integer 값을 출력하기 위해서 사용한다. 10자리 크기의 정수를 포함한 문자열 형식을 가진다.
%b : 이진(바이너리)값을 출력한다.
%c : 정수에 해당하는 문자를 출력한다.
%x : Hex 인코딩 값을 출력한다.
%f : 부동 소숫 점 값을 출력한다.
%e : 과학적 표기법으로 출력한다.
%E : 과학적 표기법으로 출력한다. %e와 다른점은 e가 대문자인지 소문자인지
%s : 문자열을 출력할 때 사용한다. %d와 더불어 가장 많이 사용하는 옵션
%q : 문자열에 있는 쌍다움표를 그대로 출력한다.
%x : 먼저 값을 integer로 변환한 다음 16비트 문자열로 변환해서 출력한다.
%p : 포인터의 주소 값을 출력한다.
%nd : 포맷팅에 사용하는 숫자 n을 이용해서 출력 할 넓이를 설정할 수 있다. %6d인 경우 6칸의 넓이를 가진다. 채우지 못한 곳은 왼쪽 부터 스페이스 문자로 채워진다.
%-nd : %nd와 달리 왼쪽부터 데이터가 채워지고 나머지 공간을 스페이스 문자가 채운다.
io.Writers : fmt 패키지는 데이터를 표준출력(os.Stdout)한다. 표준에러나 다른 파일로 출력을 보낼 수 있다. 여기에서는 표준에러(os.Stderr)로 보냈다.
		*/