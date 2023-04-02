package router

import (
	Ctrls "blog-http/controllers"
	"net/http"
)

func (*Router) GetCommentRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Fetch all the comments in a post, allows filtering
	mux.HandleFunc("/{postID}", Ctrls.CommentCtrl.GetComments)

	// Create a new comment
	mux.HandleFunc("/new", Ctrls.CommentCtrl.NewComment)

	// GET, DELETE or UPDATE the post
	mux.HandleFunc("/{commentID}", Ctrls.CommentCtrl.HandleComment)

	// like a post
	mux.HandleFunc("/{commentID}/like", Ctrls.CommentCtrl.LikeComment)

	// unlike a post
	mux.HandleFunc("/{commentID}/unlike", Ctrls.CommentCtrl.UnlikeComment)

	return mux
}
