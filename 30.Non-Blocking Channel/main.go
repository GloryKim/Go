//채널을 이용해서 worker 고루틴에 작업 메시지를 보낸다고 가정해 보자. 만약 더 이상 작업이 없다면, worker고루틴을 종료해야 할 것이다. 이를 위해서 worker는 더 이상 메시지가 없음을 측정 할 수 있어야 한다. 아래 예제를 보자.

package main

import "fmt"

func main() {
    jobs := make(chan int)
    done := make(chan bool)
    
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("작업 수행 : ", j)
            } else {
                fmt.Println("더 이상 작업이 없음")
                done <- true
                return
            }
        }
    }()
  
    for j := 0; j <= 3; j++ {
        jobs <- j
        fmt.Println("작업 요청 : ", j)
    }
    close(jobs)
    fmt.Println("모든 작업 마침")
    <- done
}
/*
예제 프로그램의 main() 고루틴은 jobs채널을 이용해서 worker 고루틴과 통신을 한다. main 고루틴은 더 이상 작업이 없을 때, jobs 채널을 닫는다.
worker 고루틴은 채널로 부터 메시지를 기다린다. 기존의 채널 기다리는 코드와 다른 점이 있는데, j, more := <- jobs로 두 개의 값을 기다린다. more는 채널의 상태를 측정 값이 담겨있다. 채널이 열려 있다면 true를 채널이 닫혔다면 false를 반환한다. woker는 more 값으로 채널의 닫힘을 알 수 있다.
*/
		