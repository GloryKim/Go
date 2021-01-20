package main
/*
base64는 바이너리 데이터를 문자코드에 영향을 받지 않도록 ASCII 코드들로 인코딩하는 기술을 가리키는 개념이다. 이메일을 통한 데이터통신에 많이 사용하고 있다. encoding/base64패키지를 이용해서 데이터를 base64로 인코딩 혹은 디코딩 할 수 있다.
*/
import (
    b64 "encoding/base64"
    "fmt"
)

func main() {
    data := "abc123!?$*()'-=@~~"
    
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)
    
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()
    
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _:= b64.URLEncoding.DecodeString(uEnc)
    
    fmt.Println(string(uDec))
}

/*
base64를 모두 쓰기 귀찮아서 별칭인 b64를 쓰기로 했다. 예제코드는 string 데이터를 인코딩/디코딩 하고 있다. 먼저 데이터를 인코딩 한 다음, 그 결과를 다시 디코딩 해서 서로 일치하는지 확인한다.
Go는 표준 base64인코딩과 URL-compatible base64 인코딩을 모두 지원한다. StdEncoding로 표준 인코딩을 URLEncoding로 URL-compatible 인코딩을 할 수 있다. 인코딩이 에러(error)를 반환했다면, 입력값과 출력값이 같은지 확인해보자.
표준 base64와 URL-compatible 인코딩은 결과에 약간의 차이가 있다. base64는 끝에 +가 URL-comaptible 인코딩은 -가 붙는다.
*/		