package csvformat

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

// Book 구조체는 Author와 Title을 가진다.
type Book struct {
	Author string
	Title  string
}

// Books는 Book 구조체 타입의 배열이며, 여러 책의 정보를 저장하는데 사용된다.
type Books []Book

// ToCSV 함수는 여러 개의 Book을 입력받고 io.Writer에 기록한 다음, 오류가 발생할 경우 이를 반환한다.
func (books *Books) ToCSV(w io.Writer) error {
	n := csv.NewWriter(w)
	err := n.Write([]string{"Author", "Title"})
	if err != nil {
		return err
	}
	for _, book := range *books {
		err = n.Write([]string{book.Author, book.Title})
		if err != nil {
			return err
		}
	}
	n.Flush()
	return n.Error()
}

// WriterCSVOutput 함수는 여러 책의 정보를 초기화하고 이를 os.Stdout에 기록한다.
func WriteCSVOutput() error {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title:  "The Great Gatsby",
		},
		Book{
			Author: "J D Salinger",
			Title:  "The Catcher in the Rye",
		},
	}
	return b.ToCSV(os.Stdout)
}

// WriteCSVBuffer 함수는 여러 책의 정보가 담긴 버퍼 csv를 반환한다.
func WriteCSVBuffer() (*bytes.Buffer, error) {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title:  "The Great Gatsby",
		},
		Book{
			Author: "J D Salinger",
			Title:  "The Catcher in the Rye",
		},
	}
	w := &bytes.Buffer{}
	err := b.ToCSV(w)
	return w, err
}
