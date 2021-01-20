package main
//많은 Go 프로그램들이 파일읽기와 쓰기는 기본 기능이다. 예제를 보자.

import (
    "fmt"
    "bufio"
    "io"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main() {
    dat, err := ioutil.ReadFile("/tmp/runner/dat.txt")
    check(err)
    
    fmt.Print(string(dat))
    
    f, err := os.Open("/tmp/runner/dat.txt")
    check(err)
    
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))
    
    o2, err := f.Seek(6,0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2,string(b2))
    
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
    
    _, err = f.Seek(0, 0)
    check(err)
    
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))
    
    f.Close()
}

/*
에러처리를 위해서 check 메서드를 만들었다. error가 nil이 아니면, 프로세스를 종료한다. 예제코드는 /tmp/runner/dat.txt 파일을 읽어서 테스트한다. 이 파일은 아래의 데이터를 가지고 있다.
ioutil.ReadFile메서드로 파일의 전체 데이터를 읽을 수 있다. 파일의 내용만 필요하다만 간단하게 사용 할 수 있는 편한 함수다. D
파일 데이터를 세밀하게 다루고 싶다면, Open메서드를 이용하면 된다. Open메서드는 호출하면 열린 파일에 대한 정보와 메서드들을 담고 있는 os.File 구조체를 반환한다.
Read 메서드를 이용해서 파일로 부터 정해진 크기만큼 데이터를 읽을 수 있다. 예제 코드는 5 바이트 크기의 데이터를 읽었다.
Seek 메서드로 파일의 위치를 이동 할 수 있다. Seek 메서드를 호출하기 위해서는 offset과 whence 두 개의 매개변수를 입력해야 한다. Seek 메서드를 호출하면, 파일의 whence에서 offset 만큼 떨어진 위치로 파일의 포인터를 이동한다. 예제코드에 f.Seek(6,0)을 실행하면 파일의 처음(0)에서 6만큼 위치로 포인터를 이동한다. 따라서 f.Read 메서드를 호출하면 w부터 읽기 시작할 것이다.
io 패키지는 파일 읽기 작업을 위한 유용한 메서드들을 제공한다. 예를 들어 ReadAtLeast같은 메서드를 이용해서 데이터를 읽을 수 있다. Read 메서드와 달리, 읽을 바이트 수를 매개변수로 받기 때문에 견고한 코드를 만들 수 있다.
파일 작업을 마친다음에는 Close 메서드를 이용해서 파일을 닫아주자. 파일 Open이 성공한 뒤에 defer를 사용하면 좋다. bufio패키지를 이용하면 buffered 연산을 할 수 있다.
*/