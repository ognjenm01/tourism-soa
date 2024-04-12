package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"module_blog.xws.com/handler"
	"module_blog.xws.com/infrastructure"
	"module_blog.xws.com/repo"
	"module_blog.xws.com/service"
)

func startServer(blogHandler *handler.BlogHandler, blogCommentHandler *handler.BlogCommentHandler, ratingHandler *handler.BlogRatingHandler) {
	router := mux.NewRouter().StrictSlash(true)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/blogs", blogHandler.Create)

	//router.HandleFunc("/blogs", blogHandler.Create).Methods("POST")
	router.HandleFunc("/blogs", blogHandler.GetAll).Methods("GET")
	router.HandleFunc("/blogs/{id}", blogHandler.GetById).Methods("GET")
	router.HandleFunc("/blogs", blogHandler.Update).Methods("PUT")
	router.HandleFunc("/blogs/{id}", blogHandler.Delete).Methods("DELETE")
	router.HandleFunc("/blogs/rate", blogHandler.AddRating).Methods("POST")

	router.HandleFunc("/api/blogcomments", blogCommentHandler.CreateComment).Methods("POST")
	router.HandleFunc("/api/blogcomments", blogCommentHandler.GetAllComments).Methods("GET")
	router.HandleFunc("/api/blogcomments/{id}", blogCommentHandler.UpdateComment).Methods("PUT")
	router.HandleFunc("/api/blogcomments/{id}", blogCommentHandler.GetCommentById).Methods("GET")
	router.HandleFunc("/api/blogcomments/{id}", blogCommentHandler.DeleteComment).Methods("DELETE")
	router.HandleFunc("/api/blogcomments/blog/{id}", blogCommentHandler.GetCommentsByBlogId).Methods("GET")

	router.HandleFunc("/ratings/{id}", ratingHandler.GetByBlog).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	//log.Fatal(http.ListenAndServe("localhost:3333", router))
	//port := ":" + os.Getenv("BLOGS_APP_PORT")
	port := ":8080"
	log.Fatal(http.ListenAndServe(port, router))
}

func main() {

	database1, err := infrastructure.InitDb()
	database := infrastructure.InitDb1()

	if database1 == nil {
		log.Fatalln("tuki")
	}

	if err != nil {
		log.Fatalln("Hit the road jack")
	}

	blogCommentRepository := &repo.BlogCommentRepository{DatabaseConnection: database}
	blogCommentService := &service.BlogCommentService{BlogCommentRepository: blogCommentRepository}
	blogCommentHandler := &handler.BlogCommentHandler{BlogCommentService: blogCommentService}

	ratingRepo := &repo.BlogRatingRepository{DatabaseConnection: database}
	ratingService := &service.BlogRatingService{BlogRatingRepo: ratingRepo}
	ratingHandler := &handler.BlogRatingHandler{BlogRatingService: ratingService}

	statusRepo := &repo.BlogStatusRepository{DatabaseConnection: database}
	statusService := &service.BlogStatusService{BlogStatusRepo: statusRepo}

	blogRepo := &repo.BlogRepository{DatabaseConnection: database1}
	blogService := &service.BlogService{
		BlogRepo:          blogRepo,
		BlogRatingService: ratingService,
		BlogStatusService: statusService,
	}
	blogHandler := &handler.BlogHandler{BlogService: blogService}

	startServer(blogHandler, blogCommentHandler, ratingHandler)
}
