package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

//컴파일 방법 1 : F5클릭 후 'Select Debug'에서 'Go' 선택후 'DEBUG CONSOLE'창에서 'hello world'출력 확인
//컴파일 방법 2 : 터미널 경로를 따라 실행하기 원하는 경로까지 cd 해준다. (EX:cd Github/Go_1042/01.helloworld) ->그후 명령어 창에 원하는 파일을 go run 해준다. (EX:go run main.go)