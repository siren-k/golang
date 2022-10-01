package main

import "golang/chap6/mongodb"

// mongo-go-driver 패키지는 연결 풀링과 MongoDB 데이터베이스에 대한 연결을 조정하고 구성하는 여러 방법을
// 제공한다. 이 예제에서 보여준 예시는 매우 기본적이지만, 문서 기반 데이터베이스에 데이터를 쓰고 읽는 과정이
// 얼마나 쉬운지를 보여준다. 이 패키지는 BSON 데이터 타입을 구현하며, 마샬링은 JSON을 활용하는 경우와
// 비슷하다.
func main() {
	if err := mongodb.Exec("mongodb://localhost"); err != nil {
		panic(err)
	}
}
