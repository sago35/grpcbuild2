package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"os/exec"

	"github.com/sago35/grpcbuild2/umedago"
	pb "github.com/sago35/grpcbuild2/umedago"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) Exec(ctx context.Context, in *umedago.ExecRequest) (*umedago.ExecReply, error) {
	c := in.GetCmd()

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	cmd := exec.Command(c.GetPath(), c.GetArgs()...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("%s : %s", err.Error(), stderr.String())
	}

	log.Printf("Path: %s, Args: %v", c.GetPath(), c.GetArgs())

	return &pb.ExecReply{
		Stdout: stdout.Bytes(),
		Stderr: stderr.Bytes(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGrpcbuildServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
