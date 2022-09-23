package main

import (
	templates2 "golang/chap1/templates"
)

// Go는 text/template과 html/template이라는 두 개의 템픓릿 패키지를 제공한다.
// 두 패키지는 다양한 기능과 함수를 공유한다. 일반적으로 웹 사이트를 렌더링할 때는 html/template을
// 사용하고 다른 경우에는 text/template을 사용하는 것이 좋다. 템플릿은 일반 텍스트이지만
// 중괄호 블록 안에서 변수와 함수를 사용할 수 있다.
// 또한, 템플릿 패키지는 파일 작업에 편리한 기능도 제공한다. 예제에서 사용했던 예는 임시 디렉터리에서
// 여러 템플릿을 생성한 다음, 한 중의 코드로 모든 템플릿을 읽는다.
func main() {
	if err := templates2.RunTemplate(); err != nil {
		panic(err)
	}

	if err := templates2.InitTemplates(); err != nil {
		panic(err)
	}

	if err := templates2.HTMLDifferences(); err != nil {
		panic(err)
	}
}
