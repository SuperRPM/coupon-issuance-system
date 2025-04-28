package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	greetv1 "github.com/SuperRPM/coupon-issuance-system/gen/greet/v1"
	"github.com/SuperRPM/coupon-issuance-system/gen/greet/v1/greetv1connect"
	"github.com/rs/cors"
)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Greet called with:", req.Msg)

	response := &greetv1.GreetResponse{
		Greeting: "Hello, " + req.Msg.Name,
	}
	return connect.NewResponse(response), nil
}

func main() {
	greetServer := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greetServer)
	mux.Handle(path, handler)

	// CORS 설정
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedHeaders: []string{"*"},
	}).Handler(mux)

	log.Println("서버가 8080 포트에서 시작됩니다...")
	if err := http.ListenAndServe(":8080", corsHandler); err != nil {
		log.Fatal(err)
	}
}
