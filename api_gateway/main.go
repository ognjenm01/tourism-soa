package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn1, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8095",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn1.Close()

	conn2, err := grpc.DialContext(
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

	gwServer := &http.Server{Addr: ":8099", Handler: gwmux}

	go func() {
		log.Println("Starting HTTP server on port 8099")
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)
	<-stopCh

	if err = gwServer.Close(); err != nil {
		log.Fatalln(err)
	}
}
