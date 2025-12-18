package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	

	//"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/router"
)


func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		fmt.Fprintf(os.Stderr, "DATABASE_URL is empty")
		os.Exit(1)
	}
	fmt.Println("Database URL:", databaseURL)

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Pool connection failed: %v\n", err)
		os.Exit(1)
	}
	r := router.Setup(pool)
	defer pool.Close()

	fmt.Print("Listening on port 8000 at http://localhost:8000!")
	log.Fatalln(http.ListenAndServe(":8000", r))
}
