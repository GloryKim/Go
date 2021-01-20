package main
//파일 쓰기도 파일 읽기와 비슷하다.
import(
    "io/ioutil"
    "os"
    "fmt"
    "bufio"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile("/tmp/dat1.txt", d1, 0644)
    check(err)
    
    f, err := os.Create("/tmp/dat2.txt")
    check(err)
    
    d2 := []byte{115, 111, 109, 101, 10}
    n2,err := f.Write(d2)
    check(err)
    fmt.Printf("Wrote %d byte\n", n2)
    
    n3, err := f.WriteString("writes\n")
    fmt.Printf("Wrote %d bytes\n", n3)
    
    f.Sync()
    
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("Buffered\n")
    fmt.Printf("Wrote %d byte\n", n4)
    
    w.Flush()
}
	
/*
ioutil.WriteFile 메서드로 바이트 데이터를 파일에 저장했다. 데이터의 크기를 알고 있으며, 크기가 별로 크지 않다면 간단하게 사용 할 수 있다.
저수준에서 파일을 쓰려면 os.Create로 파일을 열고, Write, WriteString과 같은 저수준 메서드를 이용해서 쓰기작업을 해야 한다. Sync와 Flush는 데이터를 안전하게 저장하기 위해서 사용한다.
bufio 패키지를 이용하면 buffered write를 할 수 있다.
*/