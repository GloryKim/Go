//for & range 를 이용해서 채널을 체크 할 수 있다.

package main

import "fmt"

func main() {
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)
    
    for elem := range queue {
        fmt.Println(elem)
    }
}

/*
string 메시지타입의 크기가 2인 채널을 만들었다. 여기에 두 개의 메시지를 전송하고 채널을 close 했다. 채널을 close 한다고 해서 즉시, 채널을 close하는 건 아니다. 채널이 비워진 후에 비로서 close 된다.
for & range 문을 이용해서 채널로 부터 메시지를 읽는다. 채널에는 두 개의 메시지가 있으므로 두 번의 루프를 돌면 채널이 완전히 비워지고, 3번째 루프를 돌때 for 문을 빠져나온다.
*/