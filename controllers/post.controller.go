package controllers

import (
	"blog-http/helpers"
	"net/http"
)

type PostController struct{}

func (*PostController) GetPosts(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, helpers.ResponseData{Success: true, Message: "Not Implemented"}, http.StatusNotImplemented)
}

func (*PostController) NewPost(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, helpers.ResponseData{Success: true, Message: "Not Implemented"}, http.StatusNotImplemented)
}

func (*PostController) HandlePost(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, helpers.ResponseData{Success: true, Message: "Not Implemented"}, http.StatusNotImplemented)
}

func (*PostController) LikePost(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, helpers.ResponseData{Success: true, Message: "Not Implemented"}, http.StatusNotImplemented)
}

func (*PostController) UnlikePost(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, helpers.ResponseData{Success: true, Message: "Not Implemented"}, http.StatusNotImplemented)
}
