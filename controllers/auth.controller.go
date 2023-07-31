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
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

type RegisterDto struct {
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"firstName" validate:"required,min=2"`
	LastName  string `json:"lastName" validate:"required,min=2"`
	Email     string `json:"email" validate:"required,lowercase,email"`
}

type LoginDto struct {
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,lowercase,email"`
}

type ForgotPasswordDto struct {
	Email string `json:"email" validate:"required,email,lowercase"`
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
	ctx := context.Background()

	if r.Method != http.MethodPost {
		errorResponse := helpers.GenerateErrorResponse("not found")
		helpers.WriteJSON(w, errorResponse, http.StatusNotFound)
		return
	}

	var loginDto LoginDto
	helpers.ReadJSON(r, &loginDto)

	err := helpers.ValidateBody(loginDto)
	if err != nil {
		validationErrors := ""
		for _, v := range err.(validator.ValidationErrors) {
			validationErrors += fmt.Sprintf("validation failed for field: %s,", strings.ToLower(v.Field()))
		}
		errorResponse := helpers.GenerateErrorResponse(validationErrors)
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	var user models.User
	var auth models.Auth
	err = db.UserCollection.FindOne(ctx, bson.M{"email": loginDto.Email}).Decode(&user)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	err = db.AuthCollection.FindOne(ctx, bson.M{"user": user.ID}).Decode(&auth)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(loginDto.Password))
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse("password or email incorrect")
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	token, err := helpers.SignJwtPayload(user.ID.Hex())
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusUnauthorized)
		return
	}
	helpers.WriteJSON(w, helpers.ResponseData{
		Success:     true,
		Message:     "Login successful",
		AccessToken: token,
	}, http.StatusOK)
}

func (*AuthController) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	if r.Method != http.MethodPost {
		errorResponse := helpers.GenerateErrorResponse("not found")
		helpers.WriteJSON(w, errorResponse, http.StatusNotFound)
	}

	var forgotPwdDto ForgotPasswordDto
	helpers.ReadJSON(r, &forgotPwdDto)

	err := helpers.ValidateBody(forgotPwdDto)
	if err != nil {
		validationErrors := ""
		for _, v := range err.(validator.ValidationErrors) {
			validationErrors += fmt.Sprintf("validation failed for field: %s,", strings.ToLower(v.Field()))
		}
		errorResponse := helpers.GenerateErrorResponse(validationErrors)
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	var user models.User
	err = db.UserCollection.FindOne(ctx, bson.M{"email": forgotPwdDto.Email}).Decode(&user)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusUnauthorized)
		return
	}

	resetCode := helpers.GenerateRandomCode()
	d := helpers.EmailData{
		To:      forgotPwdDto.Email,
		Subject: "Password reset code",
		Body:    fmt.Sprintf("Use the code: %s to reset your password", resetCode),
	}
	err = helpers.SendEmail(d)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusInternalServerError)
		return
	}

	code := models.Code{
		User:   user.ID,
		IsUsed: false,
		Code:   resetCode,
	}
	_, err = db.CodeCollection.InsertOne(ctx, code)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusInternalServerError)
		return
	}

	helpers.WriteJSON(w, helpers.ResponseData{
		Success: true,
		Message: "Reset code sent to your email",
	}, http.StatusOK)
}

func (*AuthController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	// fetch the code sent to the user
	helpers.WriteJSON(w, helpers.ResponseData{Success: true, Message: "Not Implemented"}, http.StatusNotImplemented)
}
