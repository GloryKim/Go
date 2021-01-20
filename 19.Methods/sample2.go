package main

import "fmt"

type ByteSlice []byte

func (self ByteSlice) Append(data []byte) []byte {
    return append(self, data...)
}

func (self *ByteSlice) MAppend(data []byte) {
    *self = append(*self, data...)
}

func main() {
    var b ByteSlice
    c := b.Append([]byte("Hello World"))
    fmt.Println("value: ", string(c))
   
   var f ByteSlice 
   f.MAppend([]byte("Hello World"))
   fmt.Println("pointer: ",string(f))
    
}
//Append()는 밸류 리시버 타입으로 ByteSlice 구조체를 변경할 수 없기 때문에, 값을 리턴해야만 한다. 포인터 리시버 타입 메서드인 MAppend()는 Byteslice를 변경 할 수 있기 때문에 리턴을 할 필요가 없다.