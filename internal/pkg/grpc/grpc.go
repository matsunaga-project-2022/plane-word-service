package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func NewGRPCServer() *grpc.Server {
	server := grpc.NewServer()
	return server
}
func StartGRPCServer(ctx context.Context, server *grpc.Server, port uint) (func(), error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// 3. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		err := server.Serve(listener)
		if err != nil {
			return
		}
	}()

	reflection.Register(server)

	stopServer := func() {
		log.Println("stopping gRPC server...")
		server.GracefulStop()
	}
	return stopServer, nil
}
