package main
//Go의 math/rand패키지로 의사난수(pseudorandom number)를 만들 수 있다.
import (
    "time"
    "fmt"
    "math/rand"
)

func main() {
    fmt.Print(rand.Intn(100), ",")
    fmt.Print(rand.Intn(100))
    fmt.Println()
    
    fmt.Println(rand.Float64())
    
    fmt.Print((rand.Float64() * 5)+5,",")
    fmt.Print((rand.Float64() * 5)+5)
    fmt.Println()
    
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    
    fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))
    fmt.Println()
    
    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    
    fmt.Print(r2.Intn(100), ",")
    fmt.Print(r2.Intn(100))
    fmt.Println()
    
    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    
    fmt.Print(r3.Intn(100), ",")
    fmt.Print(r3.Intn(100))
    fmt.Println()
}


/*
rand.Intn(n)은 0 <= n < 100 사이의 interger형 난수를 반환한다. rand.Float64는 0.0 <= f < 1.0 사이의 float형 난수를 반환한다. 난수의 범위를 변경하려면 사칙연산을 적절하게 사용해야 한다. 예제 코드의 경우 5.0 <= f < 10.0 사이의 난수를 반환 한다.
math/rand는 고정된 난수 시퀀스를 반환한다. 다른 난수 시퀀스를 얻기 위해서는 rand.NewSource메서드로 새로운 랜덤시드(random seed)를 설정해야 한다. 랜덤시드가 같으면 난수 시퀀스가 같기 때문에, 안전한 난수를 얻기가 까다롭다. 보통 유닉스 시간을 랜덤시드로 사용한다. 안전한(예측하기 힘든) 난수를 얻고 싶다면 crypt/rand패키지를 사용해야 한다.
예제 코드를 실행하면, 랜덤시드에 따라서 난수 시퀀스를 반환하는 것을 확인 할 수 있다.
*/