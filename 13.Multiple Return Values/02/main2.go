//
//  main2.go
//  Go
//
//  Created by Glory on 2020/11/11/
//  Copyright © 2020 Glory. All rights reserved.
//

// Multiple Return Values

package main

import (
	"errors"
	"fmt"
)

/*
Go 언어로 프로그램을 개발한다면,
많은 경우 에러와 값을 반환하기 위해서 multiple return values를 사용 할 것이다.
아래는 처리 값과 error를 함께 반환하는 예제 프로그램이다.
*/
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("분모로 0을 사용 할 수 없습니다.")
	}
	return a / b, nil
}

func main() {
	res, err := divide(4, 0)
	if err != nil {
		fmt.Println("ERROR : ", err.Error())
	} else {
		fmt.Println("4 /2 = ", res)
	}
}

/*
//
//  main2.c
//  Go
//
//  Created by Glory on 2020/11/11/
//  Copyright © 2020 Glory. All rights reserved.
//

// Multiple Return Values

#include <stdio.h>

int divide(int a, int b) {
    return a / b;
}
int main(int argc, char **argv) {
    int a = 5;
    int b = 0;
    if (b == 0) {
        printf("분모로 0을 사용 할 수 없습니다.\n");
    } else {
        printf("a / 4 = %d\n", a/b);
    }
    return 0;
}
/*
간단한 나눗셈 프로그램이다.
divide 함수는 분자와 분모로 쓰일 두 개의 매개 변수를 필요로 한다.
분모가 0일 경우에는 error를 반환한다.
C/C++의 경우에는 함수를 호출하기 전에 매개 변수를 미리 검사하는 식을 해결해야 한다.
*/

/*
위의 C 코드와 Go 코드를 비교해 보면,
Go 코드가 훨씬 깔끔하다는 걸 느낄 수 있을 것이다.
이유는 Go의 코드가 자기 완결성을 가지고 있기 때문이다.
*/
