/*
라인 필터(Line Filters)는 표준입력 혹은 파일로 부터 읽은 데이터를 변형해서 출력하기 위해서 사용한다. grep이나 sed 같은 프로그램들이 전형적인 라인필터 프로그램이다. grep 같은 경우 데이터를 줄 단위로 읽어서 일치하는 문자열이 있으면 출력하고, sed는 문자열을 치환하거나 삭제하는 등의 일을 한다.
*/

package main

import(
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
        ucl := strings.ToUpper(scanner.Text())
        fmt.Println(ucl)
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
	
/*
bufio.NewScanner를 이용해서 표준입력(os.Stdin)을 스캔하도록 했다. Scan관련 메서드는 입력 문자열의 다음 토큰을 찾는 일을 한다. 기본 토큰은 개행문자("\n")다. for 문으로 Scan 을 호출하면, 줄단위로 문자열을 읽는다. 그리고 ToUpper 메서드를 이용해서 대문자로 변경 한 다음 출력한다.
*/

//입력에 hello 치고 엔터하면 HELLO하면 된다.