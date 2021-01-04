package main

import "fmt"

func sum(nums ...int) { //This sum function takes several parameters and adds them and outputs the results.
	//이 sum 함수는 여러 개의 매개변수를 받아서 더하고 그 결과를 출력하는 일을 한다.
	//여러개의 변수를 받는 방법은 간단하다. 데이터 타입 앞에 "..."을 붙여주기만 하면 된다.
	//sum 함수내에서 매개변수는 슬라이스로 처리하면 된다.
	fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {
    sum (1,2)
	sum (1,2,3)
	/*
	 sum 함수의 사용 코드를 보자.
	 매개변수의 개수에 신경쓰지 않고 자유롭게 입력 할 수 있다.
	*/
    
    nums := []int{1,2,3,4,5}
    sum(nums...)//	 Variadic 함수에서 매개변수는 슬라이스로 다루기 때문에, 19 줄에서 처럼 슬라이스를 직접 사용해도 된다.
}


/*
Variadic Function is function that counts of variables is not defined. 



*/