//
//  main.go
//  Go
//
//  Created by Glory on 2020/11/1/
//  Copyright © 2020 Glory. All rights reserved.
//

// 슬라이스(slices)

//Slices는 go의 중요한 데이터타입이다.
//데이터를 정렬하는 관점에서 배열과 상당히 유사하다.
//차이점은 배열과 다르게 슬라이스의 크기를 변경할 수 있다.
//이를 통해 더욱 강력한 기능을 제공 하기 때문에,
//배열 대신 슬라이스를 주로 사용하게 된다.

package main

import "fmt"

/*
0 보다 큰 슬라이스를 만들기 위해서는 make를 사용해야 한다.
예제코드에서는 스트링을 저장 할 수 있는 크기가 3인 슬라이스를 만들었다.
슬라이스의 값은 zero-value로 초기화 된다.
이렇게 만든 슬라이스는 배열처럼 값을 넣고, 꺼낼 수 있다.
마찬가지로 len을 이용해서 슬라이스의 크기를 가져올 수 있다.
*/

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	/*
		값을 꺼내고 넣는 기본적인 오퍼레이션 외에, 슬라이스는 배열이 지원하지 않는 여러가지 기능들을 지원한다.
		그 중 하나가 append 내장 함수다. 이 함수를 이용해서, 슬라이스에 하나 이상의 원소를 추가 할 수 있다.
		append 함수는 슬라이스를 반환한다.
		또한 copy 함수를 이용해서 새로운 슬라이스로 복사 할 수도 있다.
		코드에서는 s 슬라이스와 크기가 같은 c 슬라이스를 만들어서, s 슬라이스를 복사했다.
	*/

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1", l)

	l = s[:5]
	fmt.Println("sl2", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	/*
	   슬라이스는 slice[low:high]문법을 이용해서 슬라이스의 일정 부분을 가리킬 수 있다.
	   코드에서 l := s[2:5]에서 l은 슬라이스 s[2],s[3],s[4]를 가리킨다. 일종의 포인터라고 볼 수 있다.
	   s[:5]는 s[0] ~ s[4]를 가리킨다. s[2:]은 s[2] ~ s[6]을 가리킨다.
	*/
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	/*
		71줄에서는 슬라이스를 초기화 하고 있다. "g", "h", "i"를 포함한 크기가 3인 슬라이스 t를 만들었다.
		배열과 마찬가지로 슬라이스를 원소로 하는 다차원 슬라이스를 만들 수 있다.
		74줄에서 81줄까지 2차원 슬라이스를 만들고 사용하는 방법을 보여주고 있다.
	*/
	fmt.Println("2d: ", twoD)
	/*

	   먼저 크기가 3인 슬라이스를 만들었다. 이 슬라이스는 슬라이스를 원소로 한다.
	   그리고 루프를 돌면서, 새로운 슬라이스를 만들어서 삽입을 했다.

	*/

}
