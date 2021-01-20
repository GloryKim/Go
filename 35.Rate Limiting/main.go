package main

import (
    "fmt"
    "time"
)

func main() {
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    
    close(requests)
    
    limiter := time.Tick(time.Millisecond * 200)
    
    for req := range requests {
        <- limiter
        fmt.Println("request", req, time.Now())
    } 
    
    burstyLimiter := make(chan time.Time, 3)
    
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }
    
    go func() {
        for t := range time.Tick(time.Millisecond * 200) {
            burstyLimiter <- t
        }
    }()
    
    burstyRequests := make(chan int, 5)
    for i := 1; i <=5 ; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    
    for req := range burstyRequests {
        <- burstyLimiter
        fmt.Println("request", req, time.Now())
    }
     
}



/*
Rate Limitng는 주로 NIC(Network Interface Controller)에서 입력과 출력 트래픽의 양(rate)을 제어하기 하기 위해서 사용한다. 이 외에 컴퓨팅 자원을 나눠서 사용해야 하거나 QoS를 위해서 사용 한다. Go에서 제공하는 채널과 tickers를 이용하면 고루틴에 대한 Rate limiting을 손쉽게 구현할 수 있다.
*/



/*
9~12에서 채널을 이용, 5개의 요청을 만들었다. 이 요청에 대한 rate limiting을 설정하는게 목표다.
Rate limiting를 설정하기 위해서 limiter 채널을 만들었다. time.Ticker는 time.Time타입의 채널을 반환한다. 그리고 매개변수로 주어진 시간간격으로 메시지를 전송한다. 따라서 이 채널은 200밀리세턴드 간격으로 메시지를 전송한다. 이제 rate limiting이 필요한 영역에서 채널을 읽으면 된다.
18~19에서 limiter 채널을 이용해서 rate limiting를 적용했다. 예상대로라면 200미리세컨드 간격으로 time.Now()를 실행 할 것이다.
23줄에서 raite limit를 해제하기 위해서 burstyLimiter채널을 만들었다. 크기가 3이므로 3개의 요청은 리미터가 해제된 상태로 처리될 것이다.
29줄에서 rate limit를 다시 설정했다. burstyLimiter의 크기가 3이므로, 3개의 요청이후로는 200밀리세컨드의 rate limit이 설정된다. 결과적으로 처음 5개는 200밀리세턴드의 rate limt로, 그 다음 3개는 limit 없이, 마지막 2개는 다시 200미리세컨드의 rate limt가 적용 될 것이다. 대략 아래와 같은 결과를 출력할 것이다.
*/