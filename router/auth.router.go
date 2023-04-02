package router

import (
	Ctrls "blog-http/controllers"
	"net/http"
)

func (*Router) GetAuthRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", Ctrls.AuthCtrl.Register)
	mux.HandleFunc("/login", Ctrls.AuthCtrl.Login)
	mux.HandleFunc("/forgot-password", Ctrls.AuthCtrl.ForgotPassword)
	mux.HandleFunc("/change-password", Ctrls.AuthCtrl.ChangePassword)

	return mux
}
