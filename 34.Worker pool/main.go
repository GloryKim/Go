//고루틴과 채널을 이용해서 워커풀을 만들어보자.

package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, result chan<-int) {
    for j := range jobs {
        fmt.Println("워커 ", id, " 작업처리 중 : ", j)
        time.Sleep(time.Millisecond * 100)
        result <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    for w := 1; w <=3; w++ {
        go worker(w, jobs, results)
    }
    
    for j := 1; j <=9; j++ {
        jobs <- j
    }
    
    close(jobs)
    for a := 1; a <=9; a++ {
        <-results
    }
}

/*
worker 함수는 고루틴을 이용해서 cuncurrent 하게 실행한다. 작업은 jobs 채널로 받고, 결과는 results 채널로 보낸다. job 실행에 드는 비용은 time.Sleep()메서드로 시뮬레이션 했다.
main() 고루틴은 jobs 채널과 results 채널을 만들고, 이들을 매개변수로 해서 worker 함수를 고루틴으로 실행한다. 코드에서는 3개를 고루틴을 실행했다.
for 문을 이용해서 9개의 작업을 할당했다. 코드의 마지막에 작업의 결과를 기다리고 프로그램을 종료한다. 프로그램을 실행하면, 3개의 워커들이 작업을 분산해서 실행하는 걸 확인 할 수 있을 것이다.
*/