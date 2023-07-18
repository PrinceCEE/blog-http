package controllers

import (
	"blog-http/helpers"
	"net/http"
)

type AuthController struct{}

type RegisterDto struct {
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var registerDto RegisterDto
	helpers.ReadJSON(r, &registerDto)

	// perform validation on the data
	// check if the user exists
	// if not, hash the password, create the user and auth, and return access token
	helpers.WriteJSON(w, helpers.ResponseData{Success: true, Message: "Not Implemented"}, http.StatusNotImplemented)
}

func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, helpers.ResponseData{Success: true, Message: "Not Implemented"}, http.StatusNotImplemented)
}

func (*AuthController) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, helpers.ResponseData{Success: true, Message: "Not Implemented"}, http.StatusNotImplemented)
}

func (*AuthController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, helpers.ResponseData{Success: true, Message: "Not Implemented"}, http.StatusNotImplemented)
}
