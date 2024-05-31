package main

import (
	"context"
	"log"
	"net/http"

	"api_gateway.xws.com/proto/blog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	blog.UnimplementedBlogServiceServer
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := blog.RegisterBlogServiceHandlerFromEndpoint(ctx, mux, "blogs-module:49155", opts)
	if err != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err)
	}

	log.Println("HTTP gateway is running on port 5002")
	if err := http.ListenAndServe(":5002", mux); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
