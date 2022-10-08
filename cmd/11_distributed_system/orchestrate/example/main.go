package main

import mongodb "orchestrate"

// 이 구성은 로컬 개발에 좋다. docker-compose up 명령이 실행되면 로컬 디렉터리가 재빌드되며, 도커는 최신 버전을 사용해 MongoDB 인스턴스에
// 연결을 생성하고 동작을 시작한다. 이 예제는 의존성 관리(dependency management)를 위해 go mod vender를 사용한다. 결과적으로 go mod cache를
// 비활성하고, go build 명령에서 예제에서 생성한 vendor 디렉터리를 사용하도록 지정했다.
//
// 이 예제는 외부 서비스에 연결해야 하는 앱 개발을 시작할 때 좋은 기준을 제공한다. 6장, '데이터베이스와 저장소의 모든 것'에서 다룬 예제에서는 데이터베이스의
// 로컬 인스턴스를 생성하는 대신 이 방법을 활용할 수 있다. 생산 단계에서는 도커 컨테이너 뒤에서 데이터 저장소를 실행하지 않으려고 할 수 있지만, 일반적으로
// 구성을 위한 정적 호스트 이름을 사용할 수 있다.
func main() {
	if err := mongodb.Exec("mongodb://mongodb:27017"); err != nil {
		panic(err)
	}
}
