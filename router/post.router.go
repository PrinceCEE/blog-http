package router

import (
	"net/http"
)

func (*Router) GetPostRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Fetch all the posts, allows filtering
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	// Create a new post
	mux.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {})

	// GET, DELETE or UPDATE the post
	mux.HandleFunc("/{postID}", func(w http.ResponseWriter, r *http.Request) {})

	// like a post
	mux.HandleFunc("/{postID}/like", func(w http.ResponseWriter, r *http.Request) {})

	// unlike a post
	mux.HandleFunc("/{postID}/unlike", func(w http.ResponseWriter, r *http.Request) {})

	return mux
}
