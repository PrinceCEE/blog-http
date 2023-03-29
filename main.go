package main

import (
	"blog-http/router"
	"log"
	"net/http"
)

func main() {
	srv := http.Server{
		Addr:    ":3000",
		Handler: router.GetRouters(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
