//
//  main.go
//  Go
//
//  Created by Glory on 2020/10/30/
//  Copyright © 2020 Glory. All rights reserved.
//

// 배열(Array)

//배열(Array)는 하나 이상의 연속된 자료들의 집단으로 구성된 자료구조다.

// In Go, an _array_ is a numbered sequence of elements of a
// specific length.

package main

import "fmt"

func main() {

	// 27 번째 줄에서 크기가 5인 int형 배열 a를 만들었다.
	// 이 배열은 5개의 int 값을 저장할 수 있다.
	// 28 번째 줄에서 a를 출력했다.
	// 배열을 선언하고 초기화하지 않을 경우에는 각 자료형의 초기 값으로 자동으로 초기화 된다.
	// int, float는 0, string는 "", boolean은 false다.
	var a [5]int
	fmt.Println("emp:", a)

	// 배열에 있는 값은 인덱스 값으로 읽거나 쓸 수 있다.
	// 33번째 줄에서 a의 다섯번째 위치에 100을 입력했다.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	//len은 go에서 제공하는 내장(builtin)함수로 배열의 크기를 반환한다.
	fmt.Println("len:", len(a))

	// 43-52에서는 2차원 배열의 사용법을 보여주고 있다.
	//차원은 얼마든지 늘릴 수 있지만, 3차원 이상의 배열은 사용하지 않는게 정신건강에 좋을 거다.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//3차원 이상 배열
	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", c)

	var threeD [3][2][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 3; k++ {
				threeD[i][j][k] = i + j + k
			}
		}
	}
	fmt.Println("2d: ", threeD)

}
