//Go는 포인터(pointer)를 제공한다. Go에서 모든 데이터는 값에 의한 전달(passed by value)로 작동한다. 이는 전달할 데이터를 복사한다는 의미다. 포인터는 값을 전달하지 않고, 레퍼런스(데이터가 저장된 위치)를 전달한다.

package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 100
    fmt.Println("초기 값 : ", i)
    
    zeroval(i)
    fmt.Println("zeroval : ", i)
    
    zeroptr(&i)
    fmt.Println("zeroptr : ", i)
   
   fmt.Println("포인터 :", &i) 
}


/*
예제 코드에는 zeroval과 zeroptr이라는 두 개의 함수가 있다. 두 함수 모두 int 타입의 매개변수를 받은 다음 "0"을 할당한다. 다른 점은 zeroval은 값을 복사한다는 점이고, zeroptr은 포인터를 넘긴다는 점이다.
zeroptr의 매개변수는 *int(C/C++ 경험이 있다면 익숙 할 거다)으로 데이터 타입 앞에 별표가 붙어있다(. 이는 이 데이터가 int형 포인터 타입임을 의미한다. *ptr은 포인터를 dereferences하라는 의미다. 포인터는 값이 아닌, 값이 저장된 위치다. 따라서, 값이 저장된 위치에 데이터를 저장하려면 dereferences과정을 거쳐서 메모리 상의 위치에 접근을 해야한다. *iptr = 0은 iptr이 가리키는 메모리에 0을 저장하라는 의미가 된다.
7 줄에서 zeroval 함수를 실행 했다. zeroval 함수는 ival에 0을 할당 했지만, 원본이 아닌 복사된 영역에 저장을 한다. main 함수의 i에 데이터를 쓰는게 아니기 때문에 i의 값은 그대로다.
11 줄에서 zeroptr 함수를 실행했다. zeroptr은 포인터를 매개변수로 가진다. 변수 i는 포인터 데이터 타입이 아니므로, &를 이용해서 reference를 넘겨야 한다. zeroptr 함수는 *iptr에 0을 할당한다. main 함수의 i가 가리키는 메모리 위치에 값을 쓰기 때문에, i의 값이 바뀌는 걸 확인 할 수 있다.
만약 매개변수의 값을 수정 해야 한다면, 포인터를 사용하면 된다. 그렇지 않을 경우에는 그냥 값 타입으로 정의해서 사용하는게 명확하다.
*/