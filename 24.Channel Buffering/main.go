//기본적으로 채널은 unbuffered하게 만들어진다. 이것은 채널(chan <-)에 단지 하나의 데이터만 보낼 수 있음을 의미한다. 만약 채널이 데이터를 버퍼링 하기 원한다면, 채널을 만들 때 채널의 크기를 설정하면 된다.

package main

import "fmt"

func main() {
    messages := make(chan string, 2)
    
    messages <- "buffered"
    messages <- "channel"
    
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}

/*
예제에서 크기가 2인 messages 채널을 만들었다. 이 message 채널은 버퍼링이 되기 때문에, 두 개의 메시지를 보낼 수 있다.
버퍼링채널은 master & worker모델의 프로그램에서 널리 사용한다. 이미지에 대한 섬네일을 만드는 프로그램이 있다고 가정해 보자. main 고루틴은 섬네일 변환 요청을 받아서, worker 고루틴에 요청을 전달한다. 요청을 받은 worker는 이미지에 대한 섬네일을 만든다. 요청을 받아서 전달하는 작업은 매우 빠르게 이루어지지만 요청을 처리하는 일은 상대적으로 느리게 작동한다. 따라서 하나의 채널만 가지고 있다면, main 고루틴은 worker 고루틴이 작업을 끝 낼때까지 요청을 받지 못하게 될 것이다.
버퍼링 채널을 이용하면, 유저 요청을 기다리지 않고 받을 수 있다.*/