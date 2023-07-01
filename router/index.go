package router

import (
	"net/http"
)

type Router struct{}

func GetRouters() *http.ServeMux {
	mux := http.NewServeMux()
	router := Router{}

	authRouter := router.GetAuthRouter()
	postRouter := router.GetPostRouter()
	commentRouter := router.GetCommentRouter()

	mux.Handle("/api/auth/", http.StripPrefix("/api/auth", authRouter))
	mux.Handle("/api/posts/", http.StripPrefix("/api/posts", postRouter))
	mux.Handle("/api/comments/", http.StripPrefix("/api/comments", commentRouter))

	return mux
}
