package controllers

import (
	"blog-http/db"
	"blog-http/helpers"
	"blog-http/models"
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthController struct{}

type RegisterDto struct {
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"firstName" validate:"required,min=2"`
	LastName  string `json:"lastName" validate:"required,min=2"`
	Email     string `json:"email" validate:"required,lowercase,email"`
}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var registerDto RegisterDto
	helpers.ReadJSON(r, &registerDto)

	err := helpers.ValidateBody(registerDto)
	if err != nil {
		validationErrors := map[string]string{}
		for _, v := range err.(validator.ValidationErrors) {
			validationErrors[v.StructField()] = v.Error()
		}

		errorResponse := helpers.GenerateErrorResponse(validationErrors, "error validating payload")

		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	var user models.User
	err = db.UserCollection.FindOne(context.Background(), bson.M{"email": registerDto.Email}).Decode(&user)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err, "user with email already exists")
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	// if not, hash the password, create the user and auth, and return access token
	hashedPwd, err := helpers.HashPassword(registerDto.Password)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err, err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	user = models.User{
		FirstName: registerDto.FirstName,
		LastName:  registerDto.LastName,
		Email:     registerDto.Email,
	}
	res, err := db.AuthCollection.InsertOne(context.Background(), user)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err, err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	auth := models.Auth{
		User:     res.InsertedID.(primitive.ObjectID),
		Password: hashedPwd,
	}
	_, err = db.AuthCollection.InsertOne(context.Background(), auth)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err, err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	// return a jwt response to the user

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
