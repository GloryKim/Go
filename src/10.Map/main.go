//
//  main.go
//  Go
//
//  Created by Glory on 2020/11/1/
//  Copyright © 2020 Glory. All rights reserved.
//

// Maps

//Map은 Go에 내장된 associative date type이다.
//다른 언어에서는 hash(ruby)나 dicts(python)라고 부르기도 한다.
//Map == Dictionary

package main

import "fmt"

func main() {

	//내장 함수인 make를 이용해서 비어있는 map을 만들 수 있다.
	//make(map[key-type]val-type)으로 사용한다.

	m := make(map[string]int)

	//key-type는 키의 데이터 타입이고 val-type는 키에 저장할 자료의 데이터 타입이다.
	//map은 key/value 쌍으로 이뤄진 데이터를 저장을 하며, name[key]=value의 문법으로 값을 저장 할 수 있다.

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)
	//Println을 이용해서 map을 출력하면, map에 있는 모든 key/value를 화면에 출력한다.
	//값을 꺼낼 때는 name[key]를 사용하면 된다.
	v1 := m["k1"]
	fmt.Println("v1 : ", v1)

	fmt.Println("len : ", len(m))
	//내장 함수인 len으로 map에 저장된 key/value의 갯수를 알아낼 수 있다.
	//delete로 map에서 key/value를 삭제 할 수 있다.
	delete(m, "k2")
	fmt.Println("map : ", m)

	_, prs := m["k2"]
	fmt.Println("prs : ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)

	/*
		key로 값을 가져올 경우 2개의 값을 반환한다.
		첫 번째 값은 key에 저장된 value이고, 두 번째 값은 bool 값으로 true이면 key가 있다는 것을 의미한다.
		두 번째 값을 이용해서 key가 있는지를 검사 할 수 있다.
		47 줄에서 처럼 map을 만들 때, 값을 초기화 할 수도 있다.
	*/

}
