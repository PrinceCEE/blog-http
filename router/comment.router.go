package router

import (
	"net/http"
)

func (*Router) GetCommentRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Fetch all the comments in a post, allows filtering
	mux.HandleFunc("/{postID}", func(w http.ResponseWriter, r *http.Request) {})

	// Create a new comment
	mux.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {})

	// GET, DELETE or UPDATE the post
	mux.HandleFunc("/{commentID}", func(w http.ResponseWriter, r *http.Request) {})

	// like a post
	mux.HandleFunc("/{commentID}/like", func(w http.ResponseWriter, r *http.Request) {})

	// unlike a post
	mux.HandleFunc("/{commentID}/unlike", func(w http.ResponseWriter, r *http.Request) {})

	return mux
}
