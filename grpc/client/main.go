package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/wangming1993/share/grpc/discovery/consul"
	pb "github.com/wangming1993/share/grpc/proto"
	"github.com/wangming1993/share/grpc/proto/common"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	reg = flag.String("reg", "127.0.0.1:8500", "register address")
)

func main() {
	flag.Parse()

	go sayHello()
	getMember()
}

func sayHello() {
	r := consul.NewResolver("HelloService")
	b := grpc.RoundRobin(r)

	conn, err := grpc.Dial(*reg, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(2 * time.Second)
	for t := range ticker.C {
		client := pb.NewHelloServiceClient(conn)
		resp, err := client.SayHello(context.Background(), &common.Int64Message{Id: 1})
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v: Reply is %s\n", t, resp.Reply)
	}
}

func getMember() {
	r := consul.NewResolver("MemberService")
	b := grpc.RoundRobin(r)

	conn, err := grpc.Dial(*reg, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(2 * time.Second)
	for t := range ticker.C {
		client := pb.NewMemberServiceClient(conn)
		resp, err := client.GetMember(context.Background(), &pb.MemberInfoRequest{Id: "Mike"})
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v: Reply is %s\n", t, resp)
	}
}
