package main

import (
    "fmt"
    "sort"
)

type Person struct {
    Name string    
    Id   int
}

type Persons []Person

func (self Persons) Len() int {
    return len(self)
}

func (self Persons) Swap(i, j int) {
    self[i], self[j] = self[j], self[i]
}

func (self Persons) Less(i, j int) bool {
    return self[i].Id < self[j].Id
}

func (self Persons) Sort() {
    sort.Sort(Persons(self))
    fmt.Println(self)
}

func main() {
    data := Persons{
        {"ab", 33}, 
        {"hong", 21},
        {"kim", 38},
        {"Tei", 29},
        {"Tom", 28},
    }
    data.Sort()
}

/*
main() 함수에서 ByLength 슬라이스인 fruits를 만들어서 정렬을 했다. 오름차순으로 정렬되는 걸 확인 할 수 있을 것이다. Less메서드를 Len(s[i]) > len(s[j])로 바꾸는 것만으로 내림차순 정렬로 바꿀 수 있다. 테스트해보자.
아래는 구조체 타입에 대한 정렬 예제다.
*/



//Person.Id로 정렬하는게 목표다. Persons 슬라이스 타입을 만들고, Len(), Swap(), Less()를 구현했다. Persons.Sort()메서드를 만들어서 객체지향적으로 보이게 만들었다.
		