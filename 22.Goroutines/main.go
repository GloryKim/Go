//고루틴은 실행 할 수 있는 경량 스레드(lightweight thread)다. 스레드와는 다르다. Go 언어를 사용하는 이유 중 하나다.

package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {
    f("direct")
    
    go f("goroutine")
    
    go func(msg string) {
        time.Sleep(time.Second * 1)
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second * 2)
    fmt.Println("done")
}

/*
19 줄에서 f(s)함수를 동기적으로(synchronously) 호출했다. 지금까지 함수를 호출한 방법이다.
고루틴은 go를 이용해서 만들 수 있다. 고루틴으로 만들어진 코드는 마치 스레드 처럼 독립적으로 실행이 된다. 21줄에서 고루틴으로 f(s) 함수를 호출 하고 있다.
23 줄에서는 go의 anonymous 함수를 이용해서 고루틴 코드를 만들었다. 이 함수는 1초를 기다린 다음에 "going" 메시지를 출력하는 일을 한다. 고루틴으로 실행이 되기 때문에, 코드가 블락(block)되지 않고 즉시 다음 코드로 넘어간다.
만약 26줄에서 time.Sleep로 기다리지 않으면 main 함수가 종료되면서 고루틴이 작업을 완료하지 못하게 된다. 26줄을 주석 처리 한 다음에 실행을 해보자.
물론 실제 코드에서는 time.Sleep나 fmt.Scanln 등으로 고루틴의 종료를 기다리지 않는다. 대신 채널(channels)이나 sync.Wait()등의 함수를 사용한다.

*/