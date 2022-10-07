package main

import (
	"golang/internal/11_distributed_system/docker"
	"log"
	"net/http"
	"strconv"
	"time"
)

// 이 값들은 빌드 시점에 설정된다.
var (
	version   string
	builddate string
)

var versioninfo docker.VersionInfo

func init() {
	// 빌드 시간 변수를 해석(파싱)한다.
	versioninfo.Version = version
	i, err := strconv.ParseInt(builddate, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	versioninfo.BuildDate = tm
}

// 이 예제는 리눅스 아키텍처용으로 Go 바이너리를 컴파일하고 main.go 파이에 다양한 private 변수를 설정하는 스크립트를 작성한다. 버전 엔드포인트에서
// 버전 정보를 반환하기 위해 이 변수들을 사용한다. 바이너리가 컴파일되면 이 바이너리를 포함하는 도커 컨테이너가 생성된다. 바이너리에 자체 Go 런타임이
// 포함되기 때문에 아주 작은 컨테이너가 생성된다. 그런 다음, 컨테이너가 HTTP 트래픽을 수신하는 포트를 노출하면서 컨테이너를 실행한다. 마지막으로, 포트를
// localhost로 설정해 curl 명령을 실행함으로써 버전 정보가 반환되는지 확인한다.
func main() {
	http.HandleFunc("/version", docker.VersionHandler(&versioninfo))

	log.Printf("version %s listening on :8000\n", versioninfo.Version)
	panic(http.ListenAndServe(":8000", nil))
}
