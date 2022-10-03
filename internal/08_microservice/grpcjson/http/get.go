package http

import (
	"encoding/json"
	"github.com/apex/log"
	"golang/internal/08_microservice/grpcjson/keyvalue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"net/http"
)

// GetHandler 함수는 RPC Get 호출을 래핑한다.
func (c *Controller) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	kv := keyvalue.GetKeyValueRequest{Key: key}

	gresp, err := c.Get(r.Context(), &kv)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		log.Errorf("failed to get: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(gresp)
	if err != nil {
		log.Errorf("failed to marshal: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}
