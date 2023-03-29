package router

import (
	"net/http"
)

func (*Router) GetCommentRouter() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
