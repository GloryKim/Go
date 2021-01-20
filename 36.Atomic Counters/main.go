package main

import (
    "fmt"
    "time"
    "sync/atomic"
    "runtime"
)

func main() {
    var ops uint64 = 0
    
    for i := 0; i < 50; i++ {
        go func() {
            for {
                atomic.AddUint64(&ops, 1)
                runtime.Gosched()
            }
        }()
    }
    
    time.Sleep(time.Second)
    opsFinal := atomic.LoadUint64(&ops)
    fmt.Println("ops: ", opsFinal)
}


//지금까지 고루틴간 통신에는 채널을 사용했다. 실제 채널은 고루틴간 커뮤니케이션을 하기 위해서 가장 널리 사용하는 메커니즘이다. 하지만 채널외에도 사용 할 수 있는 몇 가지 옵션들이 있다. aync/atomic을 이용해서 고루틴간 atomic counter을 구현 해보도록 하자.



/*
이 코드는 ops를 카운트 한다. 50개의 고루틴이 for 루프를 돌면서 ops++ 연산을 한다. 이렇게 두 개 이상의 고루틴이 하나의 자원에 접근 하려고 하면, 원자적 연산을 위한 장치를 마련해야 한다. 즉, 어떤 고루틴이 ops를 읽고 있다면, 연산이 완전히 끝나기 전까지 다른 고루틴은 ops를 읽지 못하게 막아야 한다.
AddUint64를 이용해서 원자적 연산을 구현할 수 있다. atomic.AddUint64(&ops, 1)은 ops변수에 1씩 증가하는 연산을 원자적으로 수행 한다. main() 고루틴은 1초를 쉰 다음에 atomic.LoadUint64()메서드를 이용해서 ops에 있는 값을 읽어온다.
고루틴을 보면 runtime.Gosched() 메서드를 호출하는 것을 볼 수 있는데, 이 메서드는 다른 고루틴이 수행되도록 스케쥴을 허용하는 일을 한다. 아래 코드를 실행해보자.
*/