package main
//Go는 정규표현(regular expressions)기능을 내장하고 있다.
import (
    "bytes"
    "fmt"
    "regexp"
)

func main() {
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)
    
    r, _:= regexp.Compile("p([a-z]+)ch")
    
    fmt.Println(r.MatchString("peach"))
    fmt.Println(r.FindString("peach puch"))
    fmt.Println(r.FindStringIndex("peach puch"))
    fmt.Println(r.FindStringSubmatch("peach puch"))
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))
    
    fmt.Println(r.FindAllString("peach punch pinch", -1))
    fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
    fmt.Println(r.FindAllString("peach punch pinch", 2))
    
    fmt.Println(r.Match([]byte("peach")))
    
    r = regexp.MustCompile("p([a-z]+)ch")
    
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
    
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}


/*
Go에서 제공하는 regexp패키지를 이용해서 정규표현식을 이용 할 수 있다. MatchString()메서드는 정규표현에 일치하는 문자열이 있는지를 검사한다. 일치할 경우 true를 반환한다.
MatchString()에서는 문자열에 대한 패턴매칭을 즉시수행 했다. 하지만 많은 경우 패턴을 저장해 놓고 사용하게 될 것이다. Compile()메서드를 이용 해서 정규식이 필요 할 때마다 컴파일 할 수 있다.
FindString()메서드로 일치하는 문자열을 찾을 수 있다. 일치는 한번만 이루어진다. 전체 문자열에서 일치하는 문자열을 찾기 위해서는 FindAllString()메서드를 이용하면 된다.
Index가 붙는 메서드들은 일치하는 문자열의 위치를 반환한다.
Submatch가 붙는 메서드들은 submatch되는 문자열들까지 반환한다. 예를 들어 "p([a-z])+ch"의 경우 "p([a-z])ch"와 "([a-z]+)에 대해서 일치하는 문자열을 반환한다.
All이 붙는 메서드들은 매개변수로 몇 개까지 일치할 것인지를 받는다. -1인 경우에는 일치하는 모든 문자열을 찾는다. 양수일 경우에는 그 수만큼만 일치하는 문자열을 찾는다.
Replace개열 메서드를 이용해서 일치하는 문자열을 치환할 수 있다. ReplaceAllString는 일치하는 모든 문자열을 치환하기 위해서 사용한다. ReplaceAllFunc메서드를 이용해서 치환에 사용할 함수를 설정 할 수도 있다.
*/
