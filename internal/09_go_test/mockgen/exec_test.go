package mockgen

import (
	"errors"
	"github.com/golang/mock/gomock"
	"golang/internal/09_go_test/mockgen/internal"
	"testing"
)

func TestController_Set(t *testing.T) {
	tests := []struct {
		name         string
		getReturnVal string
		getReturnErr error
		setReturnErr error
		wantErr      bool
	}{
		{"get error", "value", errors.New("failed"), nil, true},
		{"value match", "value", nil, nil, false},
		{"no errors", "not set", nil, nil, false},
		{"set error", "not set", nil, errors.New("failed"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockGetSetter := internal.NewMockGetSetter(ctrl)
			mockGetSetter.EXPECT().
				Get("key").
				AnyTimes().
				Return(tt.getReturnVal, tt.getReturnErr)
			mockGetSetter.EXPECT().
				Set("key", gomock.Any()).
				AnyTimes().
				Return(tt.setReturnErr)

			c := &Controller{
				GetSetter: mockGetSetter,
			}
			if err := c.GetThenSet("key", "value"); (err != nil) != tt.wantErr {
				t.Errorf("Controller.Set() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
