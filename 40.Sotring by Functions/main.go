package main

import (
    "sort"
    "fmt"
)

type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}

func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j] , s[i]   
}

func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}

/*
프로그램을 만들다 보면, 내장된 sort 패키지로는 원하는 정렬을 할 수 없을 때가 있다. 예를 들어 문자열을 알파벳 순서가 아닌 문자열의 길이로 정렬해야 한다면, custom 함수를 구현해야 한다.
*/



/*
Custom 정렬 함수를 만들기 위해서는 그에 맞는 타입을 준비해야 한다. 예제에서는 []string 타입의 alias인 ByLength타입을 만들었다.
이제 sort의 interface인 Len, Less, Swap를 ByLength에 구현하면 된다. 일종의 제너릭(generic) 함수의 구현이라고 보면 된다. 보통 문자열을 정렬 할경우 Less() 메서드는 단순 string 비교를 할 것이다. 여기에서는 len() 내장 함수를 이용해서 문자열의 길이를 가져오고 이것을 비교하고 있다.
main() 함수에서 ByLength 슬라이스인 fruits를 만들어서 정렬을 했다. 오름차순으로 정렬되는 걸 확인 할 수 있을 것이다. Less메서드를 Len(s[i]) > len(s[j])로 바꾸는 것만으로 내림차순 정렬로 바꿀 수 있다. 테스트해보자.
*/