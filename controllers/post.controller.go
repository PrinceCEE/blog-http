package controllers

import (
	"blog-http/db"
	"blog-http/helpers"
	"blog-http/models"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostController struct{}

type NewPostDto struct {
	Content string `json:"content" validate:"required,min=10"`
}

func (*PostController) GetPosts(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, helpers.ResponseData{Success: true, Message: "Not Implemented"}, http.StatusNotImplemented)
}

func (*PostController) NewPost(w http.ResponseWriter, r *http.Request) {
	// get the access token
	// decode it
	// fetch the user with it
	// create a new post
	ctx := context.Background()
	if r.Method != http.MethodPost {
		errorResponse := helpers.GenerateErrorResponse("not found")
		helpers.WriteJSON(w, errorResponse, http.StatusNotFound)
		return
	}

	accessToken, _ := helpers.GetAccessTokenFromRequest(r)
	token, err := helpers.VerifyJwtPayload(accessToken)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusNotFound)
		return
	}

	var newPostDto NewPostDto
	helpers.ReadJSON(r, &newPostDto)
	err = helpers.ValidateBody(newPostDto)
	if err != nil {
		validationErrors := ""
		for _, v := range err.(validator.ValidationErrors) {
			validationErrors += fmt.Sprintf("validation failed for field: %s,", strings.ToLower(v.Field()))
		}
		errorResponse := helpers.GenerateErrorResponse(validationErrors)
		helpers.WriteJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	userId, _ := primitive.ObjectIDFromHex(token.Subject)
	var user models.User
	err = db.UserCollection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusNotFound)
		return
	}

	post := models.Post{
		User:    user.ID,
		Content: newPostDto.Content,
	}
	res, err := db.PostCollection.InsertOne(ctx, post)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusNotFound)
		return
	}
	err = db.PostCollection.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&post)
	if err != nil {
		errorResponse := helpers.GenerateErrorResponse(err.Error())
		helpers.WriteJSON(w, errorResponse, http.StatusNotFound)
		return
	}

	helpers.WriteJSON(
		w,
		helpers.ResponseData{
			Success: true,
			Message: "Post created successfully",
			Data:    post,
		},
		http.StatusOK,
	)
}

func (*PostController) HandlePost(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, helpers.ResponseData{Success: true, Message: "Not Implemented"}, http.StatusNotImplemented)
}

func (*PostController) LikePost(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, helpers.ResponseData{Success: true, Message: "Not Implemented"}, http.StatusNotImplemented)
}

func (*PostController) UnlikePost(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, helpers.ResponseData{Success: true, Message: "Not Implemented"}, http.StatusNotImplemented)
}
