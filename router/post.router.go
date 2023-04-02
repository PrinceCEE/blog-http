package router

import (
	Ctrls "blog-http/controllers"
	"net/http"
)

func (*Router) GetPostRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Fetch all the posts, allows filtering
	mux.HandleFunc("/", Ctrls.PostCtrl.GetPosts)

	// Create a new post
	mux.HandleFunc("/new", Ctrls.PostCtrl.NewPost)

	// GET, DELETE or UPDATE the post
	mux.HandleFunc("/{postID}", Ctrls.PostCtrl.HandlePost)

	// like a post
	mux.HandleFunc("/{postID}/like", Ctrls.PostCtrl.LikePost)

	// unlike a post
	mux.HandleFunc("/{postID}/unlike", Ctrls.PostCtrl.UnlikePost)

	return mux
}
