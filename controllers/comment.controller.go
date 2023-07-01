package controllers

import (
	"fmt"
	"net/http"
)

type CommentController struct{}

func (*CommentController) GetComments(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not yet implemented")
}

func (*CommentController) NewComment(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not yet implemented")
}

func (*CommentController) HandleComment(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not yet implemented")
}

func (*CommentController) LikeComment(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not yet implemented")
}

func (*CommentController) UnlikeComment(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not yet implemented")
}
