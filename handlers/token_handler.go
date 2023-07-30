package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/johnesleyer/jwt-with-go/utils"
)

var secretKey = []byte("secret-key")

func GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := utils.GenerateToken("ralph", secretKey, time.Hour*24)
	fmt.Fprintf(w, token)
}

func VerifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Token is valid")
}
