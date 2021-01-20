package main

import (
    "fmt"
    "sort"
)
	
func main() {
    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)
    
    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints: ", ints)
    
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}	

//Go 언어의 sort 패키지를 이용해서 내장 자료와 유저 정의 자료를 정렬할 수 있다. 먼저 내장 자료를 정렬하는 것을 살펴보자.



/*
string 배열과 int형 배열을 정렬하는 간단한 예제로 In-place(제자리) 정렬이기 때문에, 정렬된 값을 반환하지 않는다. sort.Strings()으로 string 슬라이스를, sort.Ints로 int형 슬라이스를 정렬하고 있다.
AreSorted계열 메서드로 정렬된 배열(혹은 슬라이스)인지 확인 할 수 있다. 정렬이 됐다면 true를 그렇지 않다면 false를 반환한다.

*/