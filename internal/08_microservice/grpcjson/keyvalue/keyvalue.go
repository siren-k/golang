package keyvalue

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"sync"
)

// KeyValue는 맵을 저장하는 구조체다.
type KeyValue struct {
	mutex sync.RWMutex
	m     map[string]string
	UnimplementedKeyValueServer
}

// NewKeyValue는 KeyValue 구조체와 KeyValue 구조체의 맵을 초기화한다.
func NewKeyValue() *KeyValue {
	return &KeyValue{
		m: make(map[string]string),
	}
}

// Set 함수는 값을 Key에 설정한 다음 이 값을 반환한다.
func (k *KeyValue) Set(ctx context.Context, r *SetKeyValueRequest) (*KeyValueResponse, error) {
	k.mutex.Lock()
	k.m[r.GetKey()] = r.GetValue()
	k.mutex.Unlock()
	return &KeyValueResponse{Value: r.GetValue()}, nil
}

// Get 함수는 입력받은 키로 값을 가져오거나, 키가 존재하지 않으면
// '찾을 수 없다(not found)'는 오류를 반환한다.
func (k *KeyValue) Get(ctx context.Context, r *GetKeyValueRequest) (*KeyValueResponse, error) {
	k.mutex.RLock()
	defer k.mutex.RUnlock()

	val, ok := k.m[r.GetKey()]
	if !ok {
		return nil, grpc.Errorf(codes.NotFound, "key not set")
	}

	return &KeyValueResponse{Value: val}, nil
}
