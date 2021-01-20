//채널은 기본적으로 양방향으로 사용할 수 있다. 반면 개발자는 채널을 읽기 전용으로 혹은 쓰기 전용으로 설정 할 수 있다.

package main
import "fmt"

func ping(pings chan<-string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<-string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}

/*
ping 함수는 쓰기 전용의 채널을 매개변수로 넘겨 받았다. 채널 pings는 쓰기만 가능하기 때문에, 읽으려고 하면 컴파일시 invalid operation에러가 발생한다.
pong 함수는 읽기 전용의 채널을 매개변수로 넘겨 받았다. 읽기만 가능하기 때문에, 쓰려고 하면 역시 컴파일 에러가 발생한다. pong 함수는 읽기 전용의 ping 으로 부터 메시지를 읽어서, 쓰기 전용의 pong채널에 데이터를 쓴다. 채널을 (읽기 혹은 쓰기에 대한)방향을 설정하면, 컴파일시 타입을 검사하기 때문에 타입을 안전하게 사용 할 수 있다.
*/