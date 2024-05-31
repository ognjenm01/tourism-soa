package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"module_blog.xws.com/handler"
	"module_blog.xws.com/infrastructure"
	"module_blog.xws.com/proto/blog"
	"module_blog.xws.com/repo"
	"module_blog.xws.com/service"
)

func startServer(blogCommentHandler *handler.BlogCommentHandler, ratingHandler *handler.BlogRatingHandler) {
	router := mux.NewRouter().StrictSlash(true)

	//postRouter := router.Methods(http.MethodPost).Subrouter()
	//postRouter.HandleFunc("/blogs", blogHandler.Create)

	router.HandleFunc("/api/blogcomments", blogCommentHandler.CreateComment).Methods("POST")
	router.HandleFunc("/api/blogcomments", blogCommentHandler.GetAllComments).Methods("GET")
	router.HandleFunc("/api/blogcomments/{id}", blogCommentHandler.UpdateComment).Methods("PUT")
	router.HandleFunc("/api/blogcomments/{id}", blogCommentHandler.GetCommentById).Methods("GET")
	router.HandleFunc("/api/blogcomments/{id}", blogCommentHandler.DeleteComment).Methods("DELETE")
	router.HandleFunc("/api/blogcomments/blog/{id}", blogCommentHandler.GetCommentsByBlogId).Methods("GET")

	router.HandleFunc("/ratings/{id}", ratingHandler.GetByBlog).Methods("GET")

	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	//println("Server starting")
	//log.Fatal(http.ListenAndServe("localhost:3333", router))
	//port := ":" + os.Getenv("BLOGS_APP_PORT")
	//port := ":8080"
	//log.Fatal(http.ListenAndServe(port, router))
}

func main() {

	database1, err := infrastructure.InitDb()
	database := infrastructure.InitDb1()

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err != nil {
		log.Fatalln("mongo ne valja")
	}

	if database == nil {
		log.Fatalln("Hit the road jack")
	}
	defer infrastructure.Disconnect(database1, timeoutContext)

	infrastructure.Ping(database1)

	//blogCommentRepository := &repo.BlogCommentRepository{DatabaseConnection: database1}
	//blogCommentService := &service.BlogCommentService{BlogCommentRepository: blogCommentRepository}
	//blogCommentHandler := &handler.BlogCommentHandler{BlogCommentService: blogCommentService}

	ratingRepo := &repo.BlogRatingRepository{DatabaseConnection: database}
	ratingService := &service.BlogRatingService{BlogRatingRepo: ratingRepo}
	//ratingHandler := &handler.BlogRatingHandler{BlogRatingService: ratingService}

	statusRepo := &repo.BlogStatusRepository{DatabaseConnection: database}
	statusService := &service.BlogStatusService{BlogStatusRepo: statusRepo}

	blogRepo := &repo.BlogRepository{DatabaseConnection: database1}
	blogService := &service.BlogService{
		BlogRepo:          blogRepo,
		BlogRatingService: ratingService,
		BlogStatusService: statusService,
	}
	blogHandler := &handler.BlogHandler{BlogService: blogService}

	//port := ":" + os.Getenv("BLOGS_APP_PORT")
	port := ":49155"
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

	blog.RegisterBlogServiceServer(grpcServer, blogHandler)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}()
	fmt.Println("Serving gRPC")
	grpcServer.Serve(lis)

	/*stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh*/

	//startServer(blogCommentHandler, ratingHandler)
	//grpcServer.Stop()
}
