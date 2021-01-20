//Go에서 제공하는 time 패키지의 NewTimer메서드를 호출하면, 채널을 반환한다. NewTimer는 설정된 시간이 지난 후 채널에 메시지를 보내는데, 이를 이용해서 타이머를 구현 할 수 있다.

package main
import (
    "time"
    "fmt"
)

func main() {
    time1 := time.NewTimer(time.Second * 2)
    
    <- time1.C
    fmt.Println("Timer 1 expired")
    
    timer2 := time.NewTimer(time.Second * 1)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}

//2초 후에 메시지를 전달하는 채널 데이터 timer1을 만들었다. 이 코드는 <-timer1.C에서 블럭된다. 2초 후 타이머가 작동해서 메시지를 수신하면 블럭이 풀리고 다음 코드로 진행한다.
//Stop() 메서드를 이용해서 타이머를 중단 할 수도 있다. 예제코드에서는 timer2.Stop()를 호출해서 타이머를 중단했다. 타이머가 중단되면 채널로 메시가 전송되지 않는다. 따라서 worker 고루틴은 <-timer2.C 에서 블락된다. Worker 고루틴이 메시지를 받기 전에 main() 고루틴이 종료되므로, worker 고루틴은 코드를 완료하지 못하게 된다. Reset메서드를 이용해서 타이머를 재 설정 할 수도 있다.