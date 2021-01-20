//Go의 채널은 기본적으로 블럭킹(Blocking)모드로 작동한다. 하지만 select 의 default문을 이용해서 논 블록킹 채널을 구현할 수 있다. 논 블럭킹 채널을 이용하면 두 개 이상의 selects를 기다릴 수 있다.

package main

import "fmt"

func main() {
    message := make(chan string)
    signals := make(chan bool)
    
    select {
    case msg := <-message:
        fmt.Println("Received message ", msg)
    default:
        fmt.Println("no message received")
        
    }
    
    msg:="hi"
  
    select {
    case message <- msg:
        fmt.Println("Send message ", msg)
    default:
        fmt.Println("no message sent")
    } 
    
    select {
    case msg := <-message:
        fmt.Println("Receivced message ", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}

/*
9~15 줄을 보자. message 채널에서 메시지를 기다리고 있다. 원래대로라면 메시지를 기다리면서 블럭돼야 겠지만, default 문이 수행된다.
19~26 줄에서 message 채널에 메시지를 보내고 있다. 채널을 기다리는 대상이 없기 때문에 default문이 수행된다. default 문을 삭제하고 코드를 실행하면 fatal error: all goroutines are asleep - deadlock! 에러 메시지를 출력하면서 실행 실패한다. 메시지 채널을 기다리는 고루틴이 없어서 영원히 블럭되기 때문이다.
28~34 줄에서 처럼 두 개 이상의 채널을 비 봉쇄로 읽을 수 있다.
*/