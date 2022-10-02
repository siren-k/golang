package filedirs

import (
	"errors"
	"io"
	"os"
)

func Operate() error {
	// 이 0755는 명령줄의 chown에서 볼 수 있는 것과 비슷하다.
	// 다음 명령은 /tmp/example 디렉터리를 생성한다.
	// 또한 상태 경로 대시 절대 경로를 사용할 수도 있다.
	if err := os.Mkdir("example_dir", os.FileMode(0755)); err != nil {
		return err
	}
	// /tmp 디렉토리로 이동한다.
	if err := os.Chdir("example_dir"); err != nil {
		return err
	}
	// f는 일반적인 파일 객체이며 여러 인터페이스를 구현한다.
	// 파일을 열 때 적절한 비트를 설정하면 reader 또는 writer로 사용할 수 있다.
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}
	// 파일에 알고 있는 길이(known-length) 값을 쓰고 쓰기가 제대로 완료됐는지 검증한다.
	value := []byte("hello")
	count, err := f.Write(value)
	if err != nil {
		return err
	}
	if count != len(value) {
		return errors.New("incorrect length returned from write")
	}
	if err = f.Close(); err != nil {
		return err
	}
	// 파일을 읽는다.
	f, err = os.Open("test.txt")
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, f)
	if err = f.Close(); err != nil {
		return err
	}
	// /tmp 디렉토리로 이동한다.
	if err = os.Chdir(".."); err != nil {
		return err
	}
	// cleanup, os.RemoveAll은 잘못된 디렉터리를 가리키거나 사용자 입력을 사용하는 경우에
	// 위험할 수 있다. 특히 root에서 실행하는 경우 위험할 수 있다.
	if err = os.RemoveAll("example_dir"); err != nil {
		return err
	}
	return nil
}
