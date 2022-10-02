package nulls

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type nullInt64 sql.NullInt64

// ExampleNullInt 구조체도 동일하지만
// sql.NullInt64를 사용한다는 점만 다르다.
type ExampleNullInt struct {
	Age  *nullInt64 `json:"age,omitempty"`
	Name string     `json:"name"`
}

func (v *nullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}
	return json.Marshal(nil)
}

func (v *nullInt64) UnmarshalJSON(b []byte) error {
	v.Valid = false
	if b != nil {
		v.Valid = true
		return json.Unmarshal(b, &v.Int64)
	}
	return nil
}

// NullEncoding 함수는 nil/omitted 값을 처리하는
// 다른 방법을 보여준다.
func NullEncoding() error {
	e := ExampleNullInt{}

	// no는 유효하지 않은 값을 의미한다는 점에 주의한다.
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("nullInt64 Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("nullInt64 Marshal, with no age:", string(value))

	if err = json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("nullInt64 Unmarshal, with age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	fmt.Println("nullInt64 Marshal, with age = 0:", string(value))

	return nil
}
