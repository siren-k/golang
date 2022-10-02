package main

import (
	"golang/internal/01_io_filesystem/filedirs"
)

// 파일 객체를 사용하는 작업은 인메모리 스트림을 활용한 작업과 매우 유사하다. 파일은
// 또한 chown, stat, truncate와 같은 다양한 편의 기능을 직접 제공한다.
func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}
	if err := filedirs.CapitalizerExample(); err != nil {
		panic(err)
	}
}
