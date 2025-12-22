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
	//TODO stop using temp database url implmeent .env
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatalln("DATABASE_URL is empty")
	}

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Pool connection failed: %v\n", err)
	}
	r := router.Setup(pool)
	defer pool.Close()

	fmt.Print("Listening on port 8000 at http://localhost:8000!")
	log.Fatalln(http.ListenAndServe(":8000", r))
}
