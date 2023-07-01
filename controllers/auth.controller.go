package controllers

import (
	"fmt"
	"net/http"
)

type AuthController struct{}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not yet implemented")
}

func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not yet implemented")
}

func (*AuthController) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not yet implemented")
}

func (*AuthController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not yet implemented")
}
