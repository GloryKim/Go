package main

import (
    s "strings"
    "fmt"
)

var p = fmt.Println

func main() {
    p("Contents: ", s.Contains("test","es"))
    p("Count: ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index: ", s.HasSuffix("test", "e"))
    p("Join: ", s.Join([]string{"a","b"}, "-"))
    p("Repeat: ", s.Repeat("a", 5))
    p("Replace: ", s.Replace("foo", "o", "0", -1))
    p("Replace: ", s.Replace("foo", "o", "0", 1))
    p("Split: ", s.Split("a-b-c-d-e", "-"))
    p("ToLower: ", s.ToLower("TEST"))
    p("ToUpper: ", s.ToUpper("test"))
    p()
    p("Len: ", len("Hello"))
    p("Char: ", "hello"[1]) 
}
//Go는 문자열을 다루기 위해서 표준 라이브러리 형식의 strings 패키지를 제공한다. strings 패키지의 몇 가지 사용방법을 살펴보려 한다.




//자주 사용하는 fmt.Println를 대신하기 위해서 짧은 이름의 변수를 만들었다. strings 패키지가 제공하는 함수는 여기에서 확인 할 수 있다.
//https://golang.org/pkg/strings/