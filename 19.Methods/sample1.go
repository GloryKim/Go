//Go는 클래스를 지원하지 않는다. 대신 구조체에 메서드를 연결 하는 식으로 클래스 처럼 사용 할 수 있다.

package main

import "fmt"

type rect struct {
    width, height int
}

func (self *rect) area() int {
    return self.width * self.height
}

func (self rect) perim() int {
    return 2*self.width + 2*self.height
}

func main() {
    r := rect{width: 10, height:5}
    
    fmt.Println("Area: ", r.area())
    fmt.Println("perim: ", r.perim())
    
    rp := &r
    fmt.Println("Area: ", rp.area())
    fmt.Println("perim: ", rp.perim())
   
}

/*
rect라는 이름의 구조체를 정의 하고 있다. 이 스트럭쳐는 width와 height 두 개의 변수를 가지고 있다.
11 줄에서 rect 구조체에 대한 메서드를 정의했다. 메서드는 recever type을 이용해서 특정 구조체에 연결하는 방식으로 정의 할 수 있다. 코드에서 area()는 *rect에 연결된 메서드임을(*rect를 수신하는 메서드) 알 수 있다. 메서드는 연결된 스트럭쳐의 변수들을 사용 할 수 있다. main 함수에서는 메서드를 호출하는 방법을 보여주고 있다. 다른 객체지향 언어처럼 .을 이용해서 메서드를 호출 할 수 있다.
메서드는 포인터 리시버 타입과 밸류 리시버 타입을 가질 수 있다. area()는 포인터 리시버 타입의 메서드, perim()은 밸류 리시버 타입의 메서드다. 포인터 리시버 타입은 값의 복사를 막거나 스트럭쳐를 변경하기 위해서 사용 할 수 있다. 아래 코드는 밸류 리시버 타입과 포인터 리시버 타입의 차이를 보여준다.
*/