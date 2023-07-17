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

	helper := helpers.Helpers{}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler: helper.Logger(router.GetRouters()),
	}

	// Connect to the DB and shut down if there's an error
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server connected to the DB")
	fmt.Println("Server listening on PORT: 3001")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
