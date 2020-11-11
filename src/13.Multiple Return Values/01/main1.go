//
//  main1.go
//  Go
//
//  Created by Glory on 2020/11/11/
//  Copyright © 2020 Glory. All rights reserved.
//

// Multiple Return Values

package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	/*
		Go 언어는 한 번에 두 개 이상의 값을 반환하는 걸 허용한다.
		이 기능은 특히 처리 결과와 에러를 함께 리턴하기 위한 용도로 널리 사용한다.
	*/

	_, c := vals()
	fmt.Println(c)
	/*
		vals 함수는 두 개의 int 값을 반환한다.
		개발자는 multiple assignment를 이용해서 여러 개의 값을 할당할 수 있다.
		만약 반환 값을 할당하고 싶지 않다면 언더바(_)를 이용하면 된다.
	*/
}
