package interfaces

import (
	"io"
	"os"
)

// PipeExample 함수는 io 인터페이스를 사용하는 더 많은 에제를 제공한다.
func PipeExample() error {
	// pipe reader와 pipe writer는
	// io.Reader와 io.Writer를 구현한다.
	r, w := io.Pipe()
	// 이 코드는 reader가 정리를 위해 마지막에 닫을 때까지
	// 대기하기 위해 블록(block) 방식으로 동작하기 때문에
	// 별도의 go 루틴에서 실행해야 한다.
	go func() {
		// 여기서는 단순한 내용을 작성한다.
		// 이 내용 역시 json, base64 등으로 인코딩 가능한다.
		w.Write([]byte("test\n"))
		w.Close()
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
