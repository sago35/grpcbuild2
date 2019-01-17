package main

import (
	"context"
	"fmt"
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

	reply, err := c.Recv(ctx, &pb.RecvRequest{Files: os.Args[1:]})
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, r := range reply.GetFiles() {
		fmt.Printf("%s\n", r.GetFilename())
		err := ioutil.WriteFile(r.GetFilename(), r.GetData(), 0644)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}
}
