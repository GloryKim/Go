package main
/*
Command Ling Flag는 명령행 프로그램의 옵션을 처리하기 위해서 사용한다. 예를 들어 wc -l에서 -l 이 command line flag다.
Go언어는 command line flag의 처리를 위한 flag 패키지를 제공한다.
*/
import (
    "fmt"
    "flag"
)

func main() {
    wordPtr := flag.String("word", "foo", "a string")
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")
   
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var") 
    
    flag.Parse()
    
    fmt.Println("word: ", *wordPtr)
    fmt.Println("numb: ", *numbPtr)
    fmt.Println("bool: ", *boolPtr)
    fmt.Println("svar: ", svar) 
    fmt.Println("tail: ", flag.Args())
}
/*
이 코드는 실행 할 수 있지만 아직은 명령행 인자를 받지 못하기 때문에 기본 값만을 출력한다.
flag패키지는 string, integer, boolean 타입의 옵션을 처리 할 수 있다. flag.String("word","foo", "a string") 코드를 보자. "world"가 옵션이고 "foo"는 기본 값이다. world 옵션을 생략 할 경우 world에 "foo"가 입력된다. 마지막 값 "a string"는 도움말이다. flag 패키지는 도움 말을 위한 -h 옵션을 기본으로 지원한다. flag 메서드들은 모두 포인터를 반환한다. 따라서 코드에서 값을 사용하기 위해서는 포인터를 사용해야 한다. fmt.Println 메서드를 보면 모두 포인터를 사용하고 있음을 알 수 있다.
flag.Int 는 integer 값을 옵션으로 받기 위해서, flag.Boolean은 불리언 값을 옵션으로 받기 위해서 사용한다.
아래 flag 플래그 사용 예제들이다.


# go build command-line-flags.go

# ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# ./command-line-flag -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
*/
		