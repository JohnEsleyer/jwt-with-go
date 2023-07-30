package main

import (
	"fmt"
	"net/http"

	"github.com/johnesleyer/jwt-with-go/handlers"
	"github.com/johnesleyer/jwt-with-go/middleware"
)

func main() {
	http.HandleFunc("/generate-token", handlers.GenerateTokenHandler)
	http.HandleFunc("/verify-token", middleware.JWTMiddleware(handlers.VerifyTokenHandler))

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
