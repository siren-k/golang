package controllers

import (
	"encoding/json"
	"net/http"
)

// SetValue 함수는 컨트롤러 객체의 기존 저장소를 수정한다.
func (c *Controller) SetValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	value := r.FormValue("value")
	c.storage.Put(value)

	w.WriteHeader(http.StatusOK)
	p := Payload{Value: value}
	if payload, err := json.Marshal(p); err != nil {
		w.Write(payload)
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
