/*
여러 개의 고루틴을 관리하다 보면, 이들을 동기화 해야 하는 필요가 생긴다. 여기에서는 고루틴이 끝날 때까지 기다리는 방법을 살펴본다.s
*/

package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Println("working...")
    time.Sleep(time.Second * 1)
    fmt.Println("done")
    
    done <- true
}

func main() {
    done := make(chan bool, 1)
    go worker(done)
    <- done
}
//worker는 매개변수로 채널을 받는다. 소모되는 작업시간은 sleep로 대신했다. 작업이 끝나면 채널에 작업이 끝냈음을 알리는 메시지를 전송한다.