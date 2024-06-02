package main

import (
	"context"
	"log"
	"net/http"

	"api_gateway.xws.com/proto/blog"
	"api_gateway.xws.com/proto/stakeholders"
	"api_gateway.xws.com/proto/tour"
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

	err1 := blog.RegisterBlogServiceHandlerFromEndpoint(ctx, mux, "blogs-module:49155", opts)
	err2 := tour.RegisterTourProgressServiceHandlerFromEndpoint(ctx, mux, "tours-module:7777", opts)
	err3 := tour.RegisterTouristPositionServiceHandlerFromEndpoint(ctx, mux, "tours-module:7777", opts)
	err4 := tour.RegisterKeypointServiceHandlerFromEndpoint(ctx, mux, "tours-module:7777", opts)
	err5 := stakeholders.RegisterAccessTokenServiceHandlerFromEndpoint(ctx, mux, "stakeholders-module:4119", opts)

	if err1 != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err1)
	}
	if err2 != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err2)
	}

	if err3 != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err3)
	}

	if err4 != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err4)
	}
	if err5 != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err5)
	}

	if err := http.ListenAndServe(":5002", mux); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
