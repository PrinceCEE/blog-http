package main

import (
	"blog-http/db"
	"blog-http/router"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	srv := &http.Server{
		Addr:    ":3000",
		Handler: router.GetRouters(),
	}

	// Load the environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Connect to the DB and shut down if there's an error
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server connected to the DB")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server listening on PORT: 3000")
}
