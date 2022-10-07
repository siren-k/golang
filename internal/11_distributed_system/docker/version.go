package docker

import (
	"encoding/json"
	"net/http"
	"time"
)

// VersionInfo는 빌드 시점에 전달되는 값들을 저장한다.
type VersionInfo struct {
	Version   string
	BuildDate time.Time
	Uptime    time.Duration
}

// VersionHandler 함수는 최신 버전 정보를 쓴다.
func VersionHandler(v *VersionInfo) http.HandlerFunc {
	t := time.Now()
	return func(w http.ResponseWriter, r *http.Request) {
		v.Uptime = time.Since(t)
		version, err := json.Marshal(v)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(version)
	}
}
