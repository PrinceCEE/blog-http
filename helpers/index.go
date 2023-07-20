package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type Helpers struct{}

type ResponseData struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	Data        any    `json:"data,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
}

func (h *Helpers) Logger(f http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Printf("%s %s - ", r.Method, r.URL.Path)
		f.ServeHTTP(w, r)
		fmt.Printf("%d ms\n", time.Since(now).Milliseconds())
	}
}

func WriteJSON(w http.ResponseWriter, data ResponseData, code int) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(code)
	w.Write(jsonData)

}

func ReadJSON(r *http.Request, data any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}

func ValidateBody(data any) error {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		return err
	}

	return nil
}

func GenerateErrorResponse(err any, msg string) ResponseData {
	return ResponseData{
		Success: false,
		Message: msg,
		Data:    err,
	}
}

func HashPassword(pwd string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPwd), nil
}
