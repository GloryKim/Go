//Go의 select를 이용하면 여러 개의 채널로 부터 메시지를 기다릴 수 있다.


package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)
    
    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "첫번째 고루틴"
    }()
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "두번째 고루틴"
    }()
    
    for i:=0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("Received: ", msg1)
        case msg2 := <-c2:
            fmt.Println("Received: ", msg2)
        }
    }
}

/*
예제 코드는 두 개의 고루틴을 가지고 있다. main 함수는 이들 고루틴으로 부터 메시지를 기다리는데, 고루틴에 따라서 서로 다른 처리를 해야 한다.
24 번째 줄에서 select를 이용해서 복수의 채널로 부터의 메시지를 기다리며 블럭된다. 어떤 채널로 부터 메시지가 발생하면, 블럭이 풀리고 case 문을 이용해서 메시지가 발생한 채널로 부터 메시지를 읽을 수 있다. 채널 다중화(multiplexing)를 위한 툴로 볼 수 있겠다.
*/