package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key")

func generateToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = "john_doe"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time (24 hours)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Return the token to the client
	fmt.Fprintf(w, tokenString)
}

func verifyToken(w http.ResponseWriter, r *http.Request) {
	// Check for the "Authorization" header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	// Extract the token from the header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		http.Error(w, "Invalid token format", http.StatusUnauthorized)
		return
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	// Check for token parsing errors
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing token: %v", err), http.StatusUnauthorized)
		return
	}

	// Check if the token is valid
	if !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Token is valid
	fmt.Fprint(w, "Token is valid")
}

func main() {
	http.HandleFunc("/generate-token", generateToken)
	http.HandleFunc("/verify-token", verifyToken)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
