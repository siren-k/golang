package main

import (
	"golang/internal/01_io_filesystem/tempfiles"
)

// ioutil 패키지를 사용해 임시 파일과 디렉터리를 생성할 수 있다. 여전히 파일을 직접
// 삭제해야 하지만, RemoveAll 한 줄만 추가하면 나머지 작업은 알아서 처리될 것이다.
// 테스트를 작성핼 때는 임시 파일을 사용할 것을 적극 권장한다. 임시 파일을 사용하면
// 빌드 아티패트(build artifact) 등에도 유용하다. Go ioutil 패키지는 OS 환경 설정을
// 기본값으로 사용하지만, 필요한 경우에는 다른 디렉터리를 설정할 수 있다.
func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
