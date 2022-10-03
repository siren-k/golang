package http

import (
	"encoding/json"
	"github.com/apex/log"
	"golang/internal/08_microservice/grpcjson/keyvalue"
	"net/http"
)

// Controller는 내부 KeyValue 객체를 가진다.
type Controller struct {
	*keyvalue.KeyValue
}

// SetHandler 함수는 gRPC Set 함수를 래핑한다.
func (c *Controller) SetHandler(w http.ResponseWriter, r *http.Request) {
	var kv keyvalue.SetKeyValueRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&kv); err != nil {
		log.Errorf("failed to decode: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	gresp, err := c.Set(r.Context(), &kv)
	if err != nil {
		log.Errorf("failed to set: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(gresp)
	if err != nil {
		log.Errorf("failed to marshal: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
