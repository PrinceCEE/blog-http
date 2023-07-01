package controllers

import (
	"fmt"
	"net/http"
)

type PostController struct{}

func (*PostController) GetPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not yet implemented")
}

func (*PostController) NewPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not yet implemented")
}

func (*PostController) HandlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not yet implemented")
}

func (*PostController) LikePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not yet implemented")
}

func (*PostController) UnlikePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not yet implemented")
}
