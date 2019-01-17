package main

import (
	"context"
	"io/ioutil"
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

	files := []*pb.File{}
	for _, arg := range os.Args[1:] {
		var b []byte
		b, err = ioutil.ReadFile(arg)
		if err != nil {
			panic(err)
		}
		files = append(files, &pb.File{
			Filename: arg,
			Data:     b,
		})
	}

	req := &pb.SendRequest{Files: files}
	_, err = c.Send(ctx, req)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
