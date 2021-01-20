package main
// 애플리케이션에서는 다양한 방식으로 시간정보를 출력하기 마련이다. Go의 time 패키지는 패턴을 파싱해서 시간을 출력하는 기능을 제공한다.
import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println
    t := time.Now()
    p(t.Format(time.RFC3339))
    
    t1, _:=time.Parse(
        time.RFC3339, "2012-11-01T22:08:41+00:00")
    p(t1)
    
    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, e:= time.Parse(form, "8 41 PM")
    p(t2)
    
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())
        
    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)
}
/*
RFC3339에서 정의하는 형식으로 출력했다. 시간 표기에 대한 국제 표준은 ISO8601 인데, RFC3339는 인터넷상에서 어떻게 다룰지를 규약하고 있다. ISO8601과 거의 비슷하다. time.Parse()메서드를 이용 하면, 입력된 데이이터(여기에서는 두번째 매개변수)를 주어진 형식에 따라서 파싱해서 시간 구조체로 변경할 수 있다. 여기에서는 RFC3339 형식으로 파싱했다.
Format와 Parse는 예제 기반으로 시간 출력을 위한 레이아웃을 만들 수 있다. 보통은 time 에서 가져오는 값을 출력하는 것만으로 충분하지만 사용자 애플리케이션을 위한 다양한 시간 레이아웃을 만들기 위해서 이들 메서드를 사용 할 수 있다.
출력을 명시적으로 하고 싶거나 일반적이지 않은 형식으로 시간을 사용하길 원한다면, Year(), Month(), Day() ... 와 같은 메서드를 이용해서 각 시간 값을 가져와서 fmt.Printf 메서드로 출력하면 된다.

*/