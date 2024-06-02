package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"stakeholder/handler"
	"stakeholder/infrastructure"
	stakeholders "stakeholder/proto"
	"stakeholder/repo"
	"stakeholder/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	database := infrastructure.InitDB()
	if database == nil {
		log.Fatalln("[DB] - Failed to connect to database")
		return
	}

	log.Println("[DB] - Successfully connected to database")

	personRepository := &repo.PersonRepository{DatabaseConnection: database}
	personService := &service.PersonService{PersonRepository: personRepository}
	personHandler := &handler.PersonHandler{PersonService: personService}

	userRepository := &repo.UserRepository{DatabaseConnection: database}
	userService := &service.UserService{UserRepository: userRepository}
	userHandler := &handler.UserHandler{UserService: userService}

	//port := ":" + os.Getenv("ENV_HERE")
	port := ":" + os.Getenv("STAKEHOLDERS_PORT")
	if port == ":" {
		port = ":4119"
	}
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalln(err)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(lis)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	fmt.Println("Registered gRPC server")

	fmt.Println("Listening at:" + port)

	stakeholders.RegisterPersonServiceServer(grpcServer, personHandler)
	stakeholders.RegisterUserServiceServer(grpcServer, userHandler)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}()
	fmt.Println("Serving gRPC")
	grpcServer.Serve(lis)

	select {}
}
