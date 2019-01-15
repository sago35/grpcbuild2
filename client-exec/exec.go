package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/sago35/grpcbuild2/umedago"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcbuildClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.ExecRequest{
		Cmd: &pb.Cmd{
			Path: os.Args[1],
			Args: os.Args[2:],
		},
	}
	r, err := c.Exec(ctx, req)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Print(string(r.GetStdout()))
	fmt.Print(string(r.GetStderr()))
}
