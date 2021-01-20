package main
//웹 애플리케이션을 만들다 보면 URL 문자열을 조작해야 하는 경우가 많다. Go net/url 패키지로 URL을 조작할 수 있다.
//URL :https://www.joinc.co.kr/w/man/12/URI
import (
    "fmt"
    "net"
    "net/url"
)

func main() {
    s := "postgres://yundream:mypassword@joinc.co.kr:5432/path?k=v#f"
    
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }
    
    fmt.Println(u.Scheme)
    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)
    
    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)
    
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)
    fmt.Println(u.RawQuery)
    
    m, _:= url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}
	
/*
테스트에 사용 할 URL은 "postgres://yundream:mypassword@joinc.co.kr:5432/path?k=v#f"이다. 스키마, 유저, 패스워드, 호스트이름, 포트, 패스, 쿼리 파라메터 등 URL을 위한 모든 요소를 가지고 있다.
스키마 : 자원의 타입 정보다. http, ftp, mail, postgres 등 모든 종류의 자원을 표현할 수 있다. http는 가장 유명한 스키마 중 하나일 뿐이다.
url.Parse 메서드로 URL을 파싱할 수 있다. 파싱하고 나면 url.URL 구조체를 반환한다. 주요한 메서드와 값들을 정리했다.
URL.Scheme : URL 스키마.
URL.User : 유저 이름과 유저 패스워드.
URL.Host : 호스트이름과 포트번호를 포함한 문자열을 반환한다. 형식은 "host:port"다. strings.Split메서드로 host와 port를 가져올 수 있는데, net.SplitHostPort메서드를 이용하는게 더 편하다.
URL.Path : 패스. 보통 호스트내에서 자원의 위치를 기술하기 위해서 사용한다.
URL.Fragment : # 뒤에 있는 값
URL.RawQuery : ?뒤에 있는 key=value 스타일의 값이다.
URL.RawQuery는 key=value 스타일로 돼 있다. url.ParseQuery메서드로 RawQuery를 map으로 만들어서 사용 할 수 있다.
*/