package bytestrings

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"os"
	"strings"
)

// SearchString 함수는 문자열을 검색하는
// 여러 방법을 보여준다.
func SearchString() {
	s := "this is a test"
	// s는 단어 this를 포함하기 때문에
	// true를 반환한다.
	fmt.Println(strings.Contains(s, "this"))
	// s는 철자 a를 포함하기 때문에 true를 반환하며
	// b나 c를 포함하는 경우에도 같은 결과를 반환한다.
	fmt.Println(strings.ContainsAny(s, "abc"))
	// s가 this로 시작하기 때문에 true를 반환한다.
	fmt.Println(strings.HasPrefix(s, "this"))
	// s가 테스트로 끝나기 때문에 true를 반환한다.
	fmt.Println(strings.HasSuffix(s, "test"))
}

// ModifyString 함수는 문자열을 다양한 방식으로 수정한다.
func ModifyString() {
	s := "simple string"
	// [simple string]을 출력한다.
	fmt.Println(strings.Split(s, " "))
	// "Simple String"을 출력한다.
	fmt.Println(strings.Title(s))
	fmt.Println(cases.Title(language.English))
	// "simple string;";을 출력한다.
	// 앞뒤의 모든 공백은 제거된다.
	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
}

// StringReader 함수는 문자열로 io.Reader 인터페이스를
// 빠르게 생성하는 방법을 보여준다.
func StringReader() {
	s := "simple string"
	r := strings.NewReader(s)
	// 표준 출력(Stdout)에 출력한다.
	io.Copy(os.Stdout, r)
}
