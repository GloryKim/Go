//Go의 스트럭쳐는 여러 개의 필드로 구성된 데이터 타입이다. 데이터의 그룹이나 레코드를 다루는데 유용하다.

package main

import "fmt"

type person struct {
    name string
    age int
}

func main() {
    yundream := person{name: "yundream", age: 25}
    jackson := person{name: "jackson", age: 31}
    fmt.Println(yundream)
    fmt.Println(jackson)
    
    sp := &jackson
    sp.name = "Jackson json"
    sp.age = 48
    fmt.Println(sp.age)
}


/*
name과 age 두 개의 필드를 가지는 person 스트럭처를 만들었다. Go는 대부분의 객체지향 언어들이 가지는 class 키워드를 제공하지 않는다. 대신 struct를 이용해서 객체지향을 구현 할 수 있다. 객체지향을 지원하지 않는 C에서 struct와 함수 포인터를 이용해서 구현하는 것과 매우 유사한 방식이다.
현대적인 언어에 비하면 상당히 투박하게 지원한다. 실제 데이터와 메서드를 가지지만 타입 상속(type hierarchy)는 없다. 실제 Go 언어를 "Light object oriented programming language" 프로그래밍 언어라고 부르는 개발자도 있다.
13, 14 줄에서 yundream과 jackson이라는 스트럭쳐를 만들었다. 객체지향적 관점에서 보자면 클래스로 부터 객체를 만드는 과정으로 볼 수 있겠다. &를 이용해서 스트럭처에 대한 포인터를 만들 수 있다. dot(.)을 이용해서 필드에 접근 할 수 있다.


*/