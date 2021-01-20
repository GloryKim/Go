//고루틴은 코드에 동시성(concurrent)을 주기 위해서 사용한다. 두 개 이상의 코드가 실행이 되는 건데, 이 경우 코드간 통신이 중요해진다. Channels는 고루틴 간에 메시지를 주고 받기 위한 일종의 파이프다. 고루틴은 채널을 이용해서 다른 고루틴으로 데이터를 전송 할 수 있다.

package main

import (
    "fmt"
    "time"
)

func main() {
    messages := make(chan string)
    
    go func() {
        time.Sleep(time.Second * 1)
        messages <- "ping"
    }()
    
    msg := <-messages
    fmt.Println(msg)
}

/*
make(chan val-type)로 val-type의 데이터를 전송하는 채널을 만들 수 있다. 예제 코드에서는 string을 전송하는 채널을 만들었다.
채널은 파이프 처럼 작동 한다. 즉 한쪽에서 데이터를 전송하고 다른 한쪽에서 데이터를 수신하는 방식으로 사용한다. 채널로 데이터를 보내기 위해서는 channel <- data type문법을 사용하면 된다. 데이터를 수신하기 위해서는 <- channel 문법을 사용한다.
15줄에서 messages 채널로 "done"메시지를 전송하고 있다. main 함수는 18줄에서 messages 채널에서 데이터를 기다린다. time.Sleep등을 이용하지 않아도 깔끔하게 고루틴들을 동기화 할 수 있다.


*/