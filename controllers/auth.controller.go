package controllers

import (
	"blog-http/db"
	"blog-http/helpers"
	"blog-http/models"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct{}

type RegisterDto struct {
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"firstName" validate:"required,min=2"`
	LastName  string `json:"lastName" validate:"required,min=2"`
	Email     string `json:"email" validate:"required,lowercase,email"`
}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse := helpers.GenerateErrorResponse("not found")
		helpers.WriteJSON(w, errorResponse, http.StatusNotFound)
		return
	}

	var registerDto RegisterDto
	helpers.ReadJSON(r, &registerDto)

	err := helpers.ValidateBody(registerDto)
	if err != nil {
		validationErrors := ""
		for _, v := range err.(validator.ValidationErrors) {
			validationErrors += fmt.Sprintf("validation failed for field: %s,", strings.ToLower(v.Field()))
		}
		errorResponse := helpers.GenerateErrorResponse(validationErrors)
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	var user *models.User
	err = db.UserCollection.FindOne(context.Background(), bson.M{"email": registerDto.Email}).Decode(user)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		errorResponse := helpers.GenerateErrorResponse("user with email already exists")
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	hashedPwd, err := helpers.HashPassword(registerDto.Password)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	user = &models.User{
		FirstName: registerDto.FirstName,
		LastName:  registerDto.LastName,
		Email:     registerDto.Email,
	}

	res, err := db.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	userId := res.InsertedID.(primitive.ObjectID)
	user.ID = userId
	auth := models.Auth{
		User:     res.InsertedID.(primitive.ObjectID),
		Password: hashedPwd,
	}

	_, err = db.AuthCollection.InsertOne(context.Background(), auth)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	token, err := helpers.SignJwtPayload(userId.Hex())
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	helpers.WriteJSON(
		w,
		helpers.ResponseData{
			Success:     true,
			Message:     "User registered successfully",
			AccessToken: token,
			Data:        user,
		},
		http.StatusOK,
	)
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
