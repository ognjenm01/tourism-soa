package main

import (
	"fmt"
	"log"

	"module_blog.xws.com/handler"
	"module_blog.xws.com/infrastructure"
	"module_blog.xws.com/repo"
	"module_blog.xws.com/service"
)

func startServer(handler *handler.BlogHandler) {
	fmt.Println("ok")
}

func main() {
	database := infrastructure.InitDb()
	if database == nil {
		log.Fatalln("Odjebi")
	}

	blogRepo := &repo.BlogRepository{DatabaseConnection: database}
	blogService := &service.BlogService{BlogRepo: blogRepo}
	blogHandler := &handler.BlogHandler{BlogService: blogService}

	startServer(blogHandler)
}
