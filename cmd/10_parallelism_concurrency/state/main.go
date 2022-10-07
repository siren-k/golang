package main

import (
	"context"
	"golang/internal/10_parallelism_concurrency/state"
	"log"
)

// 이 에제의 Processor() 함수는 명시적으로 cancel을 호출하거나 타임아웃을 통해 컨텍스트가 취소될 때까지 게속 반복되는 함수다.
// 매개변수로 받은 다양한 연산에 따라 다른 함수를 실행할 수 있도록 모든 작업을 Process()로 전달한다. 각 기능을 훨씬 더 모듈화된
// 코드를 만들기 위해 각기 다른 함수를 전달하도록 구성할 수도 있다.
//
// 최종적으로 응답 채널에 응답이 반환되며, 루프를 통해 모든 응답을 확인한 후에 맨 마지막에 출력한다. 또한 0으로 나누기(divide by 0)를
// 예로 들면서 오류 사례를 보여준다.
func main() {
	in := make(chan *state.WorkRequest, 10)
	out := make(chan *state.WorkResponse, 10)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go state.Processor(ctx, in, out)
	req := state.WorkRequest{state.Add, 3, 4}
	in <- &req

	req2 := state.WorkRequest{state.Subtract, 5, 2}
	in <- &req2

	req3 := state.WorkRequest{state.Multiply, 9, 9}
	in <- &req3

	req4 := state.WorkRequest{state.Divide, 8, 2}
	in <- &req4

	req5 := state.WorkRequest{state.Divide, 8, 0}
	in <- &req5

	for i := 0; i < 5; i++ {
		resp := <-out
		log.Printf("Request: %v; Result: %v, Error: %v\n", resp.Wr, resp.Result, resp.Err)
	}
}
