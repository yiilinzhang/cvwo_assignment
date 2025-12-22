package auth

import (
	"os"
	"log"
	"fmt"

	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
)

//From jwt auth github adjust as needed later
var TokenAuth *jwtauth.JWTAuth

func init() {
	err := godotenv.Load()
	if err != nil {
    log.Fatal("Error loading .env file")
  }
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
  log.Fatal("JWT_SECRET not set")
}
	TokenAuth = jwtauth.New("HS256", []byte(secret), nil)
}