//time 패키지의 NewTicker를 이용해서 일정시간 간격으로 채널 메시지를 전달할 수 있다. 일정 간격으로 주기적으로 코드를 실행하는 프로그램 개발에 유용하게 사용 할 수 있다.

package main

import (
    "time" 
    "fmt"
)

func main() {
    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at ", t)
        }
        fmt.Println("Ending")
    }()
    
    time.Sleep(time.Millisecond * 1500)
    ticker.Stop()
    fmt.Println("Ticker stopped")
    
}



/*
time.NewTicker()메서드를 이용해서 0.5초 마다 메시지를 전송하는 채널을 만들었다. worker 고루틴은 0.5초 간격으로 채널로 부터 메시지를 받아서 코드를 실행한다. main() 고루틴은 1.6초를 기다린다. 따라서 worker 고루틴은 대략 2-3회 메시지를 받게 될 것이다.
Stop메서드를 이용해서 ticker를 중단 할 수 있다.
*/