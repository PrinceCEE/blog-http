package controllers

import "net/http"

type CommentController struct{}

func (*CommentController) GetComments(w http.ResponseWriter, r *http.Request) {}

func (*CommentController) NewComment(w http.ResponseWriter, r *http.Request) {}

func (*CommentController) HandleComment(w http.ResponseWriter, r *http.Request) {}

func (*CommentController) LikeComment(w http.ResponseWriter, r *http.Request) {}

func (*CommentController) UnlikeComment(w http.ResponseWriter, r *http.Request) {}
