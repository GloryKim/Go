package main
//시스템 프로그램을 개발하다보면 Unix epoch[1] 로 부터 시작되는 seconds, miliseconds, nanoseconds 등을 계산해야 될 때가 있다. Go를 이용해서 이들 값을 가져오는 방법을 알아보자.
//[1] : ^Unix time이라고 부르기도 한다. UTC 1970년 1월 1일 00:00:00시를 기준으로 지금까지 흐른 시간이다.
import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()    
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)
    
    millis := nanos / 100000
    fmt.Println("Sec    : ", secs)
    fmt.Println("Millis : ", millis)
    fmt.Println("Nanos  : ", nanos)
    
    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}
/*
Now()메서드를 호출하면, 현재의 시간정보를 담은 time.Time구조체를 반환한다. 이 구조체에서 지원하는 Unix()와 UnixNano() 메서드를 이용해서, Unix epoch 값을 가져올 수 있다. 전자는 초 단위로 후자는 나노 세컨드단위의 정보를 반환한다.
*/