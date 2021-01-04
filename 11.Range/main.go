//
//  main.go
//  Go
//
//  Created by Glory on 2020/11/03/
//  Copyright © 2020 Glory. All rights reserved.
//

// Range

package main

import "fmt"

func main() {

	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("합 : ", sum)
	/*
	   17~22 슬라이스 nums의 값을 range로 가져와서 더하는 코드다.
	   배열 역시 같은 방식으로 값을 가져올 수 있다.
	   range를 이용하면 두 개의 값을 반환한다.
	   첫번째 값은 key이고 두번째 값은 value다.
	   따라서 num 에는 슬라이스의 값이 저장된다.
	*/
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index : ", i)
		}
	}
	/*
		30~ 34 배열과 슬라이스의 경우 range의 첫번째 값으로 인덱스(index)가 반환된다.
		0부터 +1씩 증가할 것이다.
	*/
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
	//39 ~ 46 키와 값을 모두 출력하고 있다.
}

/*
range는 string도 순환 할 수 있다.
이 경우 첫번째 값은 인덱스가 되고, 두번째 값은 rune이 된다.
rune은 Go에서 지원하는 데이터타입의 하나로 UCS-4(혹은 UTF-32라고 부르기도 한다.)
한 글자를 나타내는 타입이다.
*/
