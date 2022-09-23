package confformat

import (
	"bytes"
	"github.com/BurntSushi/toml"
)

// TOMLData 구조체는 TOML 구조체 태그를 갖는 일반적인 데이터 구조체다.
type TOMLData struct {
	Name string `toml:"name"`
	Age  int    `toml:"age"`
}

// ToTOML 함수는 TOMLData 구조체를 TOML 포멧의 bytes.Buffer로 덤프(반환)한다.
func (t *TOMLData) toTOML() (*bytes.Buffer, error) {
	b := &bytes.Buffer{}
	encoder := toml.NewEncoder(b)
	if err := encoder.Encode(t); err != nil {
		return nil, err
	}
	return b, nil
}

func (t *TOMLData) Decode(data []byte) (toml.MetaData, error) {
	return toml.Decode(string(data), t)
}
