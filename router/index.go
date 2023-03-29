package router

import (
	"net/http"
)

type Router struct{}

func GetRouters() *http.ServeMux {
	mux := http.NewServeMux()
	router := Router{}

	mux.Handle("/api/auth/", http.StripPrefix("/api/auth", router.GetAuthRouter()))
	mux.Handle("/api/posts/", http.StripPrefix("/api/posts", router.GetPostRouter()))
	mux.Handle("/api/comments", http.StripPrefix("/api/comments", router.GetCommentRouter()))

	return mux
}
