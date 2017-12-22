package grpc

import (
	"log"
	"net"
	"os"
	"testing"

	pb "github.com/wangming1993/share/grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type memberServer struct{}

func (*memberServer) GetMember(ctx context.Context, req *pb.MemberInfoRequest) (*pb.MemberInfoResponse, error) {
	return &pb.MemberInfoResponse{
		Name:  "mike",
		Age:   25,
		Likes: []string{"sports", "reading", "sleep"},
	}, nil
}

func Server() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMemberServiceServer(s, &memberServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func TestMain(m *testing.M) {
	go Server()
	os.Exit(m.Run())
}

func TestMessages(t *testing.T) {

	// Set up a connection to the Server.
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMemberServiceClient(conn)

	// Test GetMember
	t.Run("GetMember", func(t *testing.T) {
		r, err := c.GetMember(context.Background(), &pb.MemberInfoRequest{})
		if err != nil {
			t.Fatalf("could not greet: %v", err)
		}
		t.Logf("Greeting: %s", r)

	})
}
