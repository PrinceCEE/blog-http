package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Helpers struct{}

type ResponseData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (h *Helpers) Logger(f http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Printf("%s - ", r.Method)
		f.ServeHTTP(w, r)
		fmt.Printf("%d ms\n", time.Since(now).Milliseconds())
	}
}

func (h *Helpers) writeJSON(w http.ResponseWriter, data ResponseData, code int) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.WriteHeader(code)
	w.Write(jsonData)

	return nil
}

func (h *Helpers) readJSON(r *http.Request, data *any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}
