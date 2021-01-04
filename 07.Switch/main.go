//
//  main.go
//  Go
//
//  Created by Glory on 2020/10/29/
//  Copyright © 2020 Glory. All rights reserved.
//

// 조건문

//Switch 문은 여러 개의 조건을 분기하기 위해서 사용한다.
//if의 목적과 상당 부분 유사하지만 약간의 차이점이 있으니 명심하자
package main

import (
	"fmt"
	"time" // time은 현 시점의 시간을 파악하기 위함이다.
)

func main() {

	//case 문을 이용해서 여러 개의 조건을 한번에 처리하는 것을 확인 할 수 있다.
	//만약 모든 case 조건을 만족하지 않는 값들에 대한 기본처리를 원하고 싶다면 default문을 이용하면 된다.
	i := 4
	fmt.Print("다음 숫자 ", i, "는 이렇게 읽습니다. ")
	switch i {
	case 3:
		fmt.Println("삼")
	case 4:
		fmt.Println("사")
	case 5:
		fmt.Println("오")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("오늘은 주말이에요!")
	default:
		fmt.Println("오늘은 평일이에요!")
	}

	//35-40 코드에서 default 문을 사용하고 있다.
	//오늘이 토요일(Saturday) 혹은 일요일(Sunday)가 아닌 경우에 대해서는 "오늘은 평일이에요!"를 출력한다.

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("지금은 오전입니다!")
	default:
		fmt.Println("지금은 오후입니다!")
	}

	// 보통 if와 switch의 차이는 조건문의 갯수와 관련이 있다.
	// 조건이 적다면 if 많다면 switch를 작성한다.
}
