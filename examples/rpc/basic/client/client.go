package main

import (
	"context"
	"fmt"

	"github.com/zhoushuguang/honor/examples/rpc/basic/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewHonorClient(conn)

	reply, err := c.GetHonor(context.Background(), &pb.GetHonorRequest{Id: 2233})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply.Title)
}
