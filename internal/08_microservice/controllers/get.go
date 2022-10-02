package controllers

import (
	"encoding/json"
	"net/http"
)

// GetValue 함수는 HandlerFunc를 래핑하는 클로저다
// UseDefault 값이 true면 value에 저장되는 값은 항상 "default"이며,
// UseDefault 값이 false면 저장소에 저장된 값이 value에 저장된다.
func (c *Controller) GetValue(UseDefault bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		value := "default"
		if !UseDefault {
			value = c.storage.Get()
		}

		w.WriteHeader(http.StatusOK)
		p := Payload{Value: value}
		if payload, err := json.Marshal(p); err == nil {
			w.Write(payload)
		}
	}
}
