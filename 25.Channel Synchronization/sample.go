//main함수는 worker 고루틴을 실행하고 채널로 부터 메시지를 기다린다. 만약 <- done을 제거하면, 이 프로그램은 worker가 작업을 종료하기 전에 종료되 버릴 거다. 이 코드는 우아한 종료에 응용 할 수 있을 거다.
package main

import (
    "fmt"
    "time"
)

type job struct {
    Name string
    Result string 
}

func worker(j chan job) {
    for {
        myjob := <-j
        if myjob.Name == "close" {
            time.Sleep(time.Second * 1)
            myjob.Result = "종료"
        } else {
            fmt.Println("작업 수행: ", myjob.Name)
            myjob.Result = "완료"
        }
        j <- myjob 
    }
}

func main() {
    done := make(chan job, 1)
    go worker(done) 
    for _,i := range []string{"job 1", "job 2"} {
        done <- job{Name:i}
        result := <-done
        fmt.Println("Job: ", result.Name, result.Result)
    }
    done <- job{Name:"close"}
    result := <- done
    fmt.Println("결과: ", result)
}

//위 코드는 우아한 종료의 예를 보여주고 있다. worker함수는 작업의 상세 내용을 담고 있는 job구조체 타입의 채널을 매개변수로 받는다. worker는 작업의 이름이 "close"일 경우에는 자원을 정리한다. main 함수는 작업이 종료되기를 기다렸다가 프로그램을 종료한다.