//
//  main.go
//  Go
//
//  Created by Glory on 2020/11/04/
//  Copyright © 2020 Glory. All rights reserved.
//

// Functions
package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

/*
plus 함수는 두 개의 int형 매개변수를 받고, int형 값을 반환한다.
함수의 형태가 C/C++ 과는 약간 다르다.
C/C++의 경우 반환 값이 함수 이름 앞에 나오는데, Go는 매개변수 뒤에 나온다.
*/

func plusPlus(a, b, c int) int {
	return a + b
	/*
		ruby 같은 언어들은 return을 사용하지 않아도 자동으로 가장 마지막 값을 반환하지만, Go 언어는 그런 거 없다.
		반드시 return 문을 이용해서 반환할 값을 명시해줘야 한다.

		동일한 데이터 타입의 매개변수가 사용 될 경우에는 데이터 타입을 생략할 수 있다.
		그러면 마지막 데이터타입으로 앞의 매개변수의 데이터 타입이 자동으로 결정된다.
	*/
}

func main() {
	res := plus(1, 2) // 함수 호출은 있는 그대로 함수이름(args)하면 된다.
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}
