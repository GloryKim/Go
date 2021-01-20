package main

import (
    "fmt"
    "runtime"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        fmt.Println(s)
        runtime.Gosched()
    }
}

func main() {
    go say("world")    
    say("hello")
}

/*
고루틴을 보면 runtime.Gosched() 메서드를 호출하는 것을 볼 수 있는데, 이 메서드는 다른 고루틴이 수행되도록 스케쥴을 허용하는 일을 한다. 아래 코드를 실행해보자.
*/


/*
"hello"와 "world"를 번갈아가면서 출력하는 걸 확인 할 수 있을 것이다. 이제 runtime.Gosched()를 주석처리 하고 실행하면, hello만 출력이 된다.
*/