package tags

import (
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

// DeSerializeStructureStrings 함수는 사용자 정의 직렬화 포맷을 사용해
// 직렬화된 문자열을 구조체로 변환한다.
func DeSerializeStructStrings(s string, res interface{}) error {
	r := reflect.TypeOf(res)

	// 포인터를 사용하도록 설정했기 때문에 항상 포인터가 전달돼야 한다.
	if r.Kind() != reflect.Ptr {
		return errors.New("res must be a pointer")
	}

	// 포인터를 역참조한다.
	r = r.Elem()
	value := reflect.ValueOf(res).Elem()

	// 직렬화된 문자열을 분할(분리)해 맵(map)에 저장한다.
	vals := strings.Split(s, ";")
	valMap := make(map[string]string)
	for _, v := range vals {
		keyval := strings.Split(v, ":")
		if len(keyval) != 2 {
			continue
		}
		valMap[keyval[0]] = keyval[1]
	}

	// 모든 필드에 대해 루프로 처리한다.
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)

		// serialize가 설정됐는지 확인한다.
		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			// "-"를 무시한다. 그렇지 않으면 전체 값이 '키(key)'가 된다.
			if serialize == "-" {
				continue
			}
			// 맵에 있는지 확인한다.
			if val, ok := valMap[serialize]; ok {
				value.Field(i).SetString(val)
			}
		} else if val, ok := valMap[field.Name]; ok {
			// 필드 이름이 맵에 있는지 확인한다.
			value.Field(i).SetString(val)
		}
	}

	return nil
}
