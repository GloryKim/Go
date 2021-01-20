//타임아웃은 외부에서의 네트워크 연결과 데이터베이스 작업등의 수행 시간을 설정하기 위해서 사용한다. Go에서는 select를 이용해서 타임아웃을 구현할 수 있다.

package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()
    
    select {
    case res:= <-c1:
        fmt.Println(res)
    case <- time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }
    
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "result 2"
    }()
    
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }
}


/*
예제 코드를 보자.
예제 코드의 첫번째 고루틴은 2초를 sleep 한 다음에 c1 채널로 메시지를 보낸다.
main 함수는 select 에서 c1로 부터 메시지를 기다린다. 그리고 time.After 함수를 이용해서 1초의 타임타임아웃을 설정했다.
time.After 함수는 매개변수에 설정된 시간이 지난 후에 Time 타입의 채널(<-chan Time)로 메시지를 반환한다.
따라서 c1 채널로 부터 메시지를 받기 전에, time.After 코드가 먼저 실행이 된다. time.After이 실행되고 select 문을 빠져나가기 때문에 c1 채널의 메시지는 받지 못할 것이다.
time.After에 3초를 설정하고 코드를 실행해 보자.
c2 채널을 사용하는 고루틴은 2초를 기다린뒤 메시지를 전송한다. select에서 time.After로 3초를 기다리기 때문에, c2로 부터 메시지를 받을 수 있다.
select를 이용한 타임아웃 패턴은 채널을 이용해서 결과 값을 받는 통신에 널리 사용한다.
*/
