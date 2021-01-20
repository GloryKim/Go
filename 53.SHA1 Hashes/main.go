// /종종 SHA1 hashes는 바이너리 파일이나 텍스트파일을 식별하기 위한 문자열을 만들기 위해서 사용한다. 예를들어서 git revision control system은 SHA1을 이용해서 파일과 디렉토리의 버전을 만든다. Go로 SHA1 해시를 만들어보자.
//Go의 crypto/*패키지는 다양한 해시 함수 구현을 가지고 있다.

package main

import (
    "crypto/sha1"
    "fmt"
)

func main() {
    s := "sha1 this string"
    h := sha1.New()
    
    h.Write([]byte(s))
    
    bs := h.Sum(nil)
    
    fmt.Println(s)
    fmt.Printf("%x\n", bs)
    
}

//sha1.New()메서드를 호출하면 hash.Hash를 반환한다. h.Write()메서드로 해시할 데이터를 입력 받고, h.Sum([]byte{})로 입력받은 데이터에 대한 SHA1 값을 만든다. h.Sum 메서드는 매개변수로 해시할 데이터의 슬라이스를 설정 할 수 있다. 보통 전체 데이터를 해시하기 때문에, 매개변수는 nil로 설정하는 경우가 많다.
//Git을 비롯한 많은 애플리케이션이 SHA1값을 hex 값으로 출력해서 사용한다. fmt.Printf 메서드의 %x를 이용해서 hex 문자열로 변환 할 수 있다. 이 코드를 실행하면 사람이 읽을 수 있는 문자열을 출력한다. crypto에는 md5 해시도 지원한다. md5.New()메서드를 이용해서 해시 값을 만들 수 있다.
	
		