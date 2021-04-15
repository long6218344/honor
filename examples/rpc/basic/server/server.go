package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zhoushuguang/honor/examples/rpc/basic/pb"
	"github.com/zhoushuguang/honor/net/rpc"

	"google.golang.org/grpc"
)

type honor struct {
	pb.UnimplementedHonorServer

	Title string
}

func (h *honor) GetHonor(ctx context.Context, req *pb.GetHonorRequest) (*pb.GetHonorReply, error) {
	return &pb.GetHonorReply{Title: h.Title}, nil
}

func main() {
	server := rpc.New()

	server.Use(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("interceptor 1 ...")
		return handler(ctx, req)
	}, func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("interceptor 2 ...")
		return handler(ctx, req)
	}, func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("interceptor 3 ...")
		return handler(ctx, req)
	}, func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("interceptors 4 ...")
		return handler(ctx, req)
	})

	pb.RegisterHonorServer(server.Server(), &honor{Title: "congratulations"})

	log.Fatal(server.Run(":9090"))
}
