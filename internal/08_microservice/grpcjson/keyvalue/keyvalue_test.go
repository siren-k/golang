package keyvalue

import (
	"reflect"
	"testing"

	"golang.org/x/net/context"
)

func TestNewKeyValue(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewKeyValue()
		})
	}
}

func TestKeyValue_Set(t *testing.T) {
	type fields struct {
		m map[string]string
	}
	type args struct {
		ctx context.Context
		r   *SetKeyValueRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *KeyValueResponse
		wantErr bool
	}{
		{"base-case", fields{make(map[string]string)}, args{context.Background(), &SetKeyValueRequest{Key: "key", Value: "value"}}, &KeyValueResponse{Value: "value"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KeyValue{
				m: tt.fields.m,
			}
			got, err := k.Set(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyValue.Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyValue.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyValue_Get(t *testing.T) {
	type fields struct {
		m map[string]string
	}
	type args struct {
		ctx context.Context
		r   *GetKeyValueRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *KeyValueResponse
		wantErr bool
	}{
		{"no key", fields{make(map[string]string)}, args{context.Background(), &GetKeyValueRequest{Key: "key"}}, nil, true},
		{"key", fields{map[string]string{"key": "value"}}, args{context.Background(), &GetKeyValueRequest{Key: "key"}}, &KeyValueResponse{Value: "value"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KeyValue{
				m: tt.fields.m,
			}
			got, err := k.Get(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyValue.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyValue.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
