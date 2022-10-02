package handlers

import (
	"encoding/json"
	"net/http"
)

// GreetingResponse 구조체는 GreetingHandler가 반환하는 JSON 응답이다.
type GreetingResponse struct {
	Payload struct {
		Greeting string `json:"greeting,omitempty"`
		Name     string `json:"name,omitempty"`
		Error    string `json:"error,omitempty"`
	} `json:"payload"`
	Successful bool `json:"successful"`
}

// GreetingHandler 함수는 오류나 유용한 정보(payload)를 갖는
// GreetingResponse를 반환한다.
func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var gr GreetingResponse
	if err := r.ParseForm(); err != nil {
		gr.Payload.Error = "bad request"
		if payload, err := json.Marshal(gr); err == nil {
			w.Write(payload)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	name := r.FormValue("name")
	greeting := r.FormValue("greeting")

	w.WriteHeader(http.StatusOK)
	gr.Successful = true
	gr.Payload.Name = name
	gr.Payload.Greeting = greeting
	if payload, err := json.Marshal(gr); err == nil {
		w.Write(payload)
	}
}
