package main

import (
	"blog-http/db"
	"blog-http/helpers"
	"blog-http/router"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load the environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("APP_PORT")
	helper := helpers.Helpers{}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: helper.Logger(router.GetRouters()),
	}

	// Connect to the DB and shut down if there's an error
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server connected to the DB")
	fmt.Printf("Server listening on PORT %s\n", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
