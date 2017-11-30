package main

import (
	"flag"
	"fmt"

	"github.com/wangming1993/share/grpc/discovery/consul"
	pb "github.com/wangming1993/share/grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	reg = flag.String("reg", "127.0.0.1:8500", "register address")
)

func main() {
	flag.Parse()

	sayHello()
	getMember()
}

func sayHello() {
	r := consul.NewResolver("HelloService")
	b := grpc.RoundRobin(r)

	conn, err := grpc.Dial(*reg, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	client := pb.NewHelloServiceClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Greeting: "world"})
	if err != nil {
		panic(err)
	}

	fmt.Printf(" Reply is %s\n", resp.Reply)
}

func getMember() {
	r := consul.NewResolver("MemberService")
	b := grpc.RoundRobin(r)

	conn, err := grpc.Dial(*reg, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	client := pb.NewMemberServiceClient(conn)
	resp, err := client.GetMember(context.Background(), &pb.MemberInfoRequest{Id: "mike"})
	if err != nil {
		panic(err)
	}

	fmt.Printf(" Reply is %+v \n", resp)
}
