package controllers

import "net/http"

type AuthController struct{}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {}

func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {}

func (*AuthController) ForgotPassword(w http.ResponseWriter, r *http.Request) {}

func (*AuthController) ChangePassword(w http.ResponseWriter, r *http.Request) {}
