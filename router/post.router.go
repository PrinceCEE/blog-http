package router

import (
	"net/http"
)

func (*Router) GetPostRouter() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
