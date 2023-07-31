package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type Helpers struct{}

type ResponseData struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	Data        any    `json:"data,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
}

type EmailData struct {
	To      string
	Subject string
	Body    string
	From    string
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
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Fatal("\n", err)
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

func GenerateErrorResponse(msg string) ResponseData {
	return ResponseData{
		Success: false,
		Message: msg,
	}
}

func HashPassword(pwd string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPwd), nil
}

func SignJwtPayload(userId string) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "blog-http",
		Subject:   userId,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 7 * 24)),
		NotBefore: jwt.NewNumericDate(time.Now()),
	})
	return t.SignedString(key)
}

func VerifyJwtPayload(s string) (jwt.RegisteredClaims, error) {
	t, err := jwt.Parse(s, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	return t.Claims.(jwt.RegisteredClaims), err
}

func SendEmail(e EmailData) error {
	e.From = os.Getenv("EMAIL_FROM")

	m := gomail.NewMessage()
	m.SetHeader("To", e.To)
	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
	m.SetHeader("Subject", e.Subject)
	m.SetBody("text/plain", e.Body)

	p, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	host := os.Getenv("GMAIL_SMTP_HOST")
	email := os.Getenv("EMAIL")
	emailPwd := os.Getenv("EMAIL_PASSWORD")
	if err != nil {
		return err
	}

	d := gomail.NewDialer(host, p, email, emailPwd)
	return d.DialAndSend(m)
}

func GenerateRandomCode() string {
	min := 100000
	max := 999999
	return strconv.Itoa(rand.Intn(max-min+1) + min)
}
