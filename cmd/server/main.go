package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"context"

	"github.com/CVWO/sample-go-app/internal/router"
	"github.com/jackc/pgx/v5"
)

func main() {
	r := router.Setup()
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	
	fmt.Print("Listening on port 8000 at http://localhost:8000!")

	log.Fatalln(http.ListenAndServe(":8000", r))
}
