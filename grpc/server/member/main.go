package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync/atomic"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/wangming1993/share/grpc/discovery/consul"
	pb "github.com/wangming1993/share/grpc/proto"
)

var (
	serv = flag.String("service", "MemberService", "service name")
	port = flag.Int("port", 1702, "listening port")
	reg  = flag.String("reg", "127.0.0.1:8500", "register address")
)

var CallTimes int64 = 0

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		panic(err)
	}

	err = consul.Register(*serv, "127.0.0.1", *port, *reg, time.Second*10, 15)
	if err != nil {
		panic(err)
	}

	log.Printf("starting member service at %d", *port)
	s := grpc.NewServer()
	pb.RegisterMemberServiceServer(s, &memberServer{})
	s.Serve(lis)
}

type memberServer struct{}

func (memberServer) GetMember(ctx context.Context, req *pb.MemberInfoRequest) (*pb.MemberInfoResponse, error) {
	atomic.AddInt64(&CallTimes, 1)
	fmt.Printf("Received request: %d times \n", atomic.LoadInt64(&CallTimes))
	return &pb.MemberInfoResponse{
		Name:  "mike",
		Age:   25,
		Likes: []string{"sports", "reading", "sleep"},
	}, nil
}
