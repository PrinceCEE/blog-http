package router

import (
	"net/http"
)

func (*Router) GetAuthRouter() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
