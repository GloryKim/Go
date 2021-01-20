package main
//Go는 JSON 인코딩과 디코딩을 위한 메서드를 내장하고 있다.
import (
    "encoding/json"
    "fmt"
    "os"
)

type Response1 struct {
    Page int
    Fruits []string
}

type Response2 struct {
    Page int `json:"page"`
    Fruits []string `json:"fruits"`
    Tag string `json:"tag,omitempty"`
    Comment string `json:"-"`
}

func main() {
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))
    
    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))
    
    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))
    
    strB, _ := json.Marshal("gohher")
    fmt.Println(string(strB))
    
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))
   
    mapD := map[string]int{"apple":5, "lettuce":7} 
    mapB, _:= json.Marshal(mapD)
    fmt.Println(string(mapB))
    
    res1D := &Response1 {
        Page: 1,
        Fruits: []string{"apple", "peach", "pear"},
    }
    res1B, _:= json.Marshal(res1D)
    fmt.Println(string(res1B))
    
    res2D := &Response2 {
        Page: 1,
        Fruits: []string{"apple", "peach", "pear"},
    }
    res2B, _:= json.Marshal(res2D)
    fmt.Println(string(res2B))
    
    byt := []byte(`{"num":6.13, "str":["a", "b"]}`)
    
    var dat map[string]interface{}
    
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println()
    
    os.Exit(0)
}
	
/*
Marshal() 메서드를 이용해서 go의 데이터들을 json 으로 변환할 수 있다. 리턴 값은 바이트배열이므로 문자열로 사용하기 위해서는 변환을 해줘야 한다.
integer, string, float같은 원시데이터 타입에 대한 변환도 물론 가능하지만 배열과 맵, 구조체를 변환하기 위해서 주로 사용한다. 32~38에서 배열과 맵의 변환을 확인 할 수 있다. 우리가 예상한 그대로의 형태로 변환된다.
구조체는 약간 다르다. 구조체의 경우 각 필드 이름이, json의 name이 된다. 개발자는 struct tag를 이용해서 각 필드에 정보를 추가 할 수 있다. JSON 파서는 필드에 있는 struct tag를 읽어서 구문에 대한 정보를 얻을 수 있다.
struct tag는 역 따움표로 설정 할 수 있다. Response2 에서 json:"page", json:"fruits"로 설정했는데, 각 필드에 대한 json name이 page와 fruits로 재 설정된다. omitempty를 설정하면, 필드에 값이 없을 경우 필드를 json name으로 변환하지 않는다. 필드를 json으로 변환하지 않고 싶다면 "-"를 설정하면 된다.
JSON 데이터는 Unmarshal()메서드를 이용해서 Go 데이터 타입으로 변환할 수 있다. HTTP 기반의 REST API에 대한 응답을 처리 할 때 널리 사용한다. JSON 패키지에 대한 자세한 내용은 Json and Go나 JSON package docs 문서를 참고하자.

https://blog.golang.org/json


https://golang.org/pkg/encoding/json/

*/