package controllers

import "net/http"

type PostController struct{}

func (*PostController) GetPosts(w http.ResponseWriter, r *http.Request) {}

func (*PostController) NewPost(w http.ResponseWriter, r *http.Request) {}

func (*PostController) HandlePost(w http.ResponseWriter, r *http.Request) {}

func (*PostController) LikePost(w http.ResponseWriter, r *http.Request) {}

func (*PostController) UnlikePost(w http.ResponseWriter, r *http.Request) {}
