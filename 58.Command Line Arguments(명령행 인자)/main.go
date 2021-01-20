package main
/*
명령행 인자(Command-line arguments)는 프로그램에 데이터를 넘기기 위해서 사용한다. 예를 들어 go run hello.go의 경우 run 과 hello.go 를 go 프로그램의 명령행 인자로 사용하고 있다.
*/
import (
    "os"
    "fmt"
)

func main() {
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]
    
    arg := os.Args[3]
    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}
/*
os.Args는 슬라이스 자료구조로, 명령행 인자들을 저장한다. 슬라이스의 첫번째에는 명령행 인자가 아닌 실행 프로그램의 이름이 들어간다. 따라서 os.Args[1:] 로 명령행 인자에 접근 할 수 있다.
아래는 명령행 인자의 처리 예제다.
*/
		