package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"api_gateway.xws.com/proto/blog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	blog.UnimplementedBlogServiceServer
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	log.Println("aplication started")

	// Create a gRPC server object
	s := grpc.NewServer()

	service, err := NewServer()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	blog.RegisterBlogServiceServer(s, service)
	log.Println("Serving gRPC on 0.0.0.0:5001")

	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:5001",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln(err)
	}
	//defer conn.Close()

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = blog.RegisterBlogServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":5002",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:5002")
	log.Fatalln(gwServer.ListenAndServe())

	/*conn2, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8096",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn2.Close()

	gwmux := runtime.NewServeMux()

	client1 := blog.NewBlogServiceClient(conn1)
	err = blog.RegisterBlogServiceHandlerClient(context.Background(), gwmux, client1)

	if err != nil {
		log.Fatalln(err)
	}

	gwServer := &http.Server{Addr: ":5001", Handler: gwmux}

	go func() {
		log.Println("Starting HTTP server on port 5001")
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)
	<-stopCh

	if err = gwServer.Close(); err != nil {
		log.Fatalln(err)
	}*/
}
