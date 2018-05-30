package tests

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/wangming1993/share/grpc/proto"
	"github.com/wangming1993/share/grpc/proto/common"
	"google.golang.org/grpc"
)

type MockGreeterClient struct {
	ctrl     *gomock.Controller
	recorder *MockGreeterClientRecorder
	proto.HelloServiceClient
}

// Recorder for MockGreeterClient (not exported)
type MockGreeterClientRecorder struct {
	mock *MockGreeterClient
}

func NewMockGreeterClient(ctrl *gomock.Controller) *MockGreeterClient {
	mock := &MockGreeterClient{ctrl: ctrl}
	mock.recorder = &MockGreeterClientRecorder{mock}
	return mock
}

func (m *MockGreeterClient) EXPECT() *MockGreeterClientRecorder {
	return m.recorder
}

func (m *MockGreeterClient) SayHello(ctx context.Context, req *common.Int64Message, params ...grpc.CallOption) (*proto.HelloResponse, error) {
	s := []interface{}{ctx, req}
	for _, p := range params {
		s = append(s, p)
	}
	ret := m.ctrl.Call(m, "SayHello", s...)
	ret0, _ := ret[0].(*proto.HelloResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGreeterClientRecorder) SayHello(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	s := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCall(mr.mock, "SayHello", s...)
}
