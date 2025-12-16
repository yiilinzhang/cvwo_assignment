package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	

	//"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5"
	"github.com/go-chi/cors"
	"github.com/yiilinzhang/cvwo_assignment/internal/router"
)


func main() {

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	databaseURL := os.Getenv("DATABASE_URL")
	fmt.Println("Database URL:", databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	r := router.Setup(conn)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))
	fmt.Print("Listening on port 8000 at http://localhost:8000!")
	log.Fatalln(http.ListenAndServe(":8000", r))
}
