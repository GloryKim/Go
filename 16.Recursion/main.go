//Go는 재귀함수(Recursive functions)를 지원한다. 아래는 고전적인 factorial 예제코드다.

package main	

import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(5))
}
        
//fact 함수는 fact(0)이 호출 될 때까지 자기 자신을 계속 호출한다.