package pipes

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// WordCount 함수는 파일을 입력받아 각 단어를 키(key)로 하고
// 단어의 등자 수를 값(value)으로 하는 맵(map)을 반환한다.
func WordCount(f io.Reader) map[string]int {
	result := make(map[string]int)

	// 파일 io.Reader 인터페이스에서 동작할 스캐너를 만든다.
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		result[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return result
}
