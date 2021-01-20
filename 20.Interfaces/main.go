//인터페이스(interface)는 메서드들의 이름의 집합이다. 메서드의 이름만 포함하고 있기 때문에, 구현이 필요하다.

package main

import "fmt"
import "math"

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}

func (r rect)area() float64 {
    return r.width * r.height
}

func (r rect)perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}


func main() {

   fmt.Println("Interfaces를 위한 예제 입니다.")    

}
/*
어떤 도형의 면적과 둘레의 길이를 구하는 함수를 만든다고 가정해 보자. 도형은 삼각형, 사각형 혹은 원형이거나 다른 도형일 수도 있다. 어떤 도형이든지 이들은 면적을 구하는 area()와 둘레를 구하는 perim() 두 개의 메서드를 공통으로 가지고 있을 거다. 이 메서드는 각 도형에서 실제 구현하게 된다.
예제에서는 rect와 circle 구조체에서 area()와 perim을 구현하고 있다.
Go의 인터페이스는 멤버 데이터 없이, 오직 메서드만 존재한다는 점에서 C++의 순수 추상 클래스(pure abstract class)와 유사한 측면이 있다. 반면 인터페이스와 구현이 완전히 분리된다는 차이 점이 있다. 코드를 보면 알겠지만 geometry 인터페이스와 circle, rect 구조체는 서로 연관성이 전혀 없다.
geometry 인터페이스와 circle,rect 구조체의 연관성은 measure 함수에서 만들어진다. measure 함수는 인터페이스를 매개변수로 받는데, 인터페이스에 있는 메서드들만 구현하고 있다면, 어떤 종류의 구조체라도 매개변수로 받을 수 있다.
main 함수에서 measure 함수는 rect와 circle구조체를 매개변수로 받아서 실행을 하고 있다. 이제 measure 함수는 rect와 circle에 있는 메서드를 호출하게 된다.
*/
