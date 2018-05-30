package tests

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/wangming1993/share/grpc/proto"
	"github.com/wangming1993/share/grpc/proto/common"
	grpc_tests "github.com/wangming1993/share/grpc/tests"
	"gotest.tools/assert"
)

func TestSayHello(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGreeterClient := grpc_tests.NewMockGreeterClient(ctrl)

	mockGreeterClient.EXPECT().SayHello(
		gomock.Any(),
		&common.Int64Message{Id: 1},
	).Return(&proto.HelloResponse{Reply: "mocked"}, nil)
	testSayHello(t, mockGreeterClient)
}

func testSayHello(t *testing.T, client proto.HelloServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SayHello(ctx, &common.Int64Message{Id: 1})

	assert.NilError(t, err)
	assert.Equal(t, res.Reply, "mocked")
}
