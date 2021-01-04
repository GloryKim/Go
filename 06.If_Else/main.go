//
//  main.go
//  Go
//
//  Created by Glory on 2020/10/28/
//  Copyright © 2020 Glory. All rights reserved.
//

// 조건문

//조건에 의해서 실행문을 분기하기 위해서 사용한다.
//"만약 A 조건을 만족하면 A` 문을 실행하고 그렇지 않으면 B 문을 실행하라"라는 식으로 작동한다.

package main

import "fmt"

func main() {

	if 9%2 == 0 {
		fmt.Println("9는 짝수 입니다.")
	} else {
		fmt.Println("9는 홀수 입니다.")
	}
	//20-24 코드는 짝수와 홀수를 구분하는 일을 한다.
	//만약(if) 9을 2로 나눈 나머지가 0이면 "9는 짝수 입니다."
	//그렇지 않으면(else) "9는 홀수 입니다."를 출력한다.
	// You can have an `if` statement without an else.
	if 12%4 == 0 {
		fmt.Println("12는 4로 나누어집니다.")
	}
	//29-31 코드는 12가 4의 배수인지를 검사한다. 4로 나눈 나머지가 0이면 "12는 4로 나누어집니다."를 출력한다.

	if num := 9; num < 0 {
		fmt.Println(num, "음수입니다!")
	} else if num < 10 {
		fmt.Println(num, "0이상 10 미만의 수 입니다.")
	} else {
		fmt.Println(num, "10 이상의 수 입니다.")
	}
	//34-40 코드를 보면 if와 else if를 여러 번 사용하고 있다.
}
