package main

import(
    "fmt" 
    "math/rand"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)

func main() {
   var state = make(map[int]int)
   var mutex = &sync.Mutex{}
   
   var ops int64 = 0 
   
   for r :=0; r < 100; r++ {
       go func() {
        total := 0
           for {
               key := rand.Intn(5)
               mutex.Lock()
               total += state[key]
               mutex.Unlock()
               atomic.AddInt64(&ops, 1)
               
               runtime.Gosched()
           }
       }()
   }
   
   for w := 0; w < 10; w++ {
       go func() {
           for {
               key := rand.Intn(5)
               val := rand.Intn(100)
               mutex.Lock()
               state[key] = val
               mutex.Unlock()
               atomic.AddInt64(&ops, 1)
               runtime.Gosched()
           }
       }()
   }
   
   time.Sleep(time.Second)
   
   opsFinal := atomic.LoadInt64(&ops)
   fmt.Println("ops:", opsFinal)
   
   mutex.Lock()
   fmt.Println("state:", state)
   mutex.Unlock()
}	


//앞 장에서 원자연산을 이용한 카운터의 구현을 살펴봤다. 여기에서는 뮤텍스(Mutex)를 이용해서 비슷한 구현을 해보려 한다. 애초에 원자적 연산은 뮤텍스를 주로 이용하니, 일반적인 방법이라고 할 수 있겠다.

/*
뮤텍스는 자원에 대한 접근을 한번에 하나씩제어하기 위해서 사용한다. 이 코드에서 관리할 자원은 map자료구조인 state다. 뮤텍스 객체는 sync.Mutex{}로 만들 수 있다.
18 ~ 31. 100개의 고루틴을 만들었다. 이 고루틴은 state로 부터 값을 읽어서 더하는 일을 한다. total += state[key]코드를 실행하기 위해서는 1. state로 부터 값을 읽어서 2. total과 더해서 저장해야 한다. 코드에 대한 원자적 연산을 보장하지 않는다면 state의 값을 읽어서 저장하기 전에, 다른 고루틴이 state의 값을 읽을 수 있다. 중복연산이 되는 셈이다. 이를 막기 위해서 mutex.Loc()와 mutex.Unlock()를 이용해서 하나의 고루틴만 진입하도록 설정했다.
33 ~ 45. 10개의 고루틴이 state에 값을 쓴다. 역시 뮤텍스를 이용해서 한번에 하나의 고루틴만 접근 할 수 있게 했다.
main 고루틴은 1초를 쉰 다음에, 그동안 연산했던 값들을 출력한다.
*/