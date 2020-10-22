# Go lang 설치방법!
1. http://golang.org/dl 사이트 가서 Go Lang을 설치한다.

## 윈도우 
2. 고언어가 설치되면 환경 변수를 설정해야 한다.
- GOROOT : 고언어가 설치된 폴더 (C:\Go\)
- GOPATH : 프로젝트 소스를 보관할 위치 입력 (D:\go-workspace)
- 환경변수 등록 방법 : 내컴퓨터 - 우클릭 후 속성 - 설정 변경 - 시스템 속성에서 고급 탭 - 환경 변수 - 시스템 변수 - 새로만들기
- 본인은 고언어 설치시 GOPATH의 환경 변수가 자동으로 등록됐다.
- 컴퓨터를 재부팅 후 명령 프롬프트 창에서 go env 또는 go version를 입력해서 고언어가 설치되어 있는지 확인한다

## MAC
2. 설치 후 '터미널'에서 아래의 명령어 3개를 작성한다.
- brew install go
- go env
- go version

## GIT 사용 방법

2. 깃을 설치하기 위해 https://git-scm.com/download/win 접속 후 GIT 설치를 진행하고 다음 과정을 진행한다.
- https://github.com/Microsoft/vscode-go 로 이동한다.
- 읽다 보면 To install the tools manually in the current GOPATH, just paste and run: 내용이 있다.
- 아래의 명령어를 명령 프롬프트 창에 실행한다.
========================================
- go get -u -v github.com/nsf/gocode
- go get -u -v github.com/rogpeppe/godef
- go get -u -v github.com/zmb3/gogetdoc
- go get -u -v github.com/golang/lint/golint
- go get -u -v github.com/lukehoban/go-outline
- go get -u -v sourcegraph.com/sqs/goreturns
- go get -u -v golang.org/x/tools/cmd/gorename
- go get -u -v github.com/tpng/gopkgs
- go get -u -v github.com/newhook/go-symbols
- go get -u -v golang.org/x/tools/cmd/guru
- go get -u -v github.com/cweill/gotests/...
- go get -u -v golang.org/x/tools/cmd/godoc
- go get -u -v github.com/fatih/gomodifytags
========================================
- bin 폴더에 파일이 있는지 확인한다
- 명령 프롬프트에 아래 명령어를 실행하여 go 디버거를 설치한다.
- go get github.com/derekparker/delve/cmd/dlv


## VisualCode로 go 실행!
3. GOPATH로 설정된 경로로 들어가서 bin, pkg, src 3개의 폴더를 생성한다. (맥에서는 자동으로 생성)
- bin : go 에서 사용하는 명령어 저장
- pkg : go get 명령어로 다운 받은 패키지 저장
- src : go 파일 소스
4. vs code를 실행해서 src 폴더에 main.go를 작성한다.
------------------------------------------
package main
 
import (
    "fmt"
)
 
func main() {
    fmt.Println("hello world")
}
------------------------------------------
5. F5를 눌러서 디버그 콘솔에서 결과를 확인한다.
6. dlv 명령어가 불가능하다고 알림이 뜨면, install 버튼을 눌러서 설치하면 된다.