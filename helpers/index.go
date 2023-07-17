package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Helpers struct{}

type ResponseData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
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

func ReadJSON(r *http.Request, data *any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}
