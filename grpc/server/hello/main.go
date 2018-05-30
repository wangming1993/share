package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/wangming1993/share/grpc/discovery/consul"
	"github.com/wangming1993/share/grpc/proto"
	"github.com/wangming1993/share/grpc/proto/common"
)

var (
	service = flag.String("service", "HelloService", "service name")
	port    = flag.Int("port", 1703, "listening port")
	reg     = flag.String("reg", "127.0.0.1:8500", "register address")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		panic(err)
	}

	err = consul.Register(*service, "127.0.0.1", *port, *reg, time.Second*10, 15)
	if err != nil {
		panic(err)
	}

	log.Printf("starting hello service at %d", *port)
	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, &helloServer{})
	s.Serve(lis)
}

type helloServer struct{}

func (helloServer) SayHello(ctx context.Context, req *common.Int64Message) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{}, nil
}

func (helloServer) Greeting(ctx context.Context, req *common.EmptyMessage) (*common.EmptyMessage, error) {
	return &common.EmptyMessage{}, nil
}
