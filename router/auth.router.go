package router

import (
	"net/http"
)

func (*Router) GetAuthRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/forgot-password", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/change-password", func(w http.ResponseWriter, r *http.Request) {})

	return mux
}
