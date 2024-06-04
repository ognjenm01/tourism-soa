package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"api_gateway.xws.com/proto/blog"
	"api_gateway.xws.com/proto/stakeholders"
	"api_gateway.xws.com/proto/tour"
	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			if key == "Authorization" {
				return key, true
			}
			return key, false
		}),
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			if authHeader := req.Header.Get("Authorization"); authHeader != "" {
				return metadata.Pairs("authorization", authHeader)
			}
			return nil
		}),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err1 := blog.RegisterBlogServiceHandlerFromEndpoint(ctx, mux, "blogs-module:49155", opts)
	err2 := tour.RegisterTourProgressServiceHandlerFromEndpoint(ctx, mux, "tours-module:7777", opts)
	err3 := tour.RegisterTouristPositionServiceHandlerFromEndpoint(ctx, mux, "tours-module:7777", opts)
	err4 := tour.RegisterKeypointServiceHandlerFromEndpoint(ctx, mux, "tours-module:7777", opts)
	err5 := stakeholders.RegisterAccessTokenServiceHandlerFromEndpoint(ctx, mux, "stakeholders-module:4119", opts)
	err6 := tour.RegisterTourServiceHandlerFromEndpoint(ctx, mux, "tours-module:7777", opts)
	err7 := tour.RegisterEquipmentServiceHandlerFromEndpoint(ctx, mux, "tours-module:7777", opts)
	err8 := tour.RegisterTourEquipmentServiceHandlerFromEndpoint(ctx, mux, "tours-module:7777", opts)
	err9 := tour.RegisterTourReviewServiceHandlerFromEndpoint(ctx, mux, "tours-module:7777", opts)

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
	if err6 != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err6)
	}
	if err7 != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err7)
	}
	if err8 != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err8)
	}
	if err9 != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err9)
	}

	authMux := authenticate(mux) // Create a new ServeMux with authentication middleware
	log.Println("HTTP gateway is running on port 5002")
	if err := http.ListenAndServe(":5002", authMux); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// authenticate middleware to validate access token
func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if r.Method == "POST" && r.URL.Path == "/api/login" {
			next.ServeHTTP(w, r)
			return
		}

		if r.Method == "POST" && r.URL.Path == "/api/register" {
			next.ServeHTTP(w, r)
			return
		}

		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		// Parse the token
		tokenString := strings.Split(authHeader, "Bearer ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("explorer_secret_key"), nil // Use the same key that was used to sign the token
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// If the token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
