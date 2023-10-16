package main

import (
	"fmt"
	"github.com/ricardojonathanromero/go-circuit-breaker/client/config"
	"github.com/ricardojonathanromero/go-circuit-breaker/client/internal/handler"
	"github.com/ricardojonathanromero/go-circuit-breaker/client/internal/request"
	"net/http"
)

func main() {
	circuitBreaker := config.CircuitBreakerConfig()
	req := request.NewHttpRequest(circuitBreaker, http.Client{}, "http://localhost:8082/api/v1/ping")
	hdl := handler.NewHandler(*req)

	http.HandleFunc("/api/v1/ping", hdl.Ping)
	fmt.Println("Client Application running .")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("error init server", err)
	}
}
