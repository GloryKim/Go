package main

import (
    "fmt"
    "os"
)

func main() {
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("Creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("Writing")
    fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
    fmt.Println("Closing")
    f.Close()
}

//Defer은 코드가 종료된 뒤에 수행할 함수를 설정하기 위해서 사용한다. 자원을 정리(cleanup)하기 위한 목적으로 주로 사용한다. 다른 언어의 encure와 finally와 비슷한 일을 한다.
/*
이 프로그램은 파일을 /tmp 디렉토리에 파일을 만들고, 간단한 문자열을 쓴 다음에 닫는 일을 한다. 10번째 줄에 defer를 이용해서 closeFile() 함수를 등록했다. 호출이 아니고 등록이라는 것에 주의하자. 지금 실행되지 않는다. writeFile() 함수의 호출이 끝나고 main()함수가 종료 할 때, defler로 등록한 closeFile() 함수가 실행된다.

*/