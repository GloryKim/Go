package main

//Go의 time 패키지를 이용해서 시간가져오고 측정 할 수 있다.

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

    now := time.Now()
    p(now)

    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    p(then.Weekday())

    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    diff := now.Sub(then)
    p(diff)

    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    p(then.Add(diff))
    p(then.Add(-diff))
}
	

/*
time.Now()메서드로 현재 시간을 가져오고 time.Date()로 시간을 설정 할 수 있다. 두 메서드 모두 time.Time 구조체를 반환하며, 구조체에서 제공하는 메서드로 다양한 연산을 할 수 있다.
*/