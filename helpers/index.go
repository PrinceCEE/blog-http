package helpers

import (
	"fmt"
	"net/http"
	"time"
)

type Helpers struct{}

func (h *Helpers) Logger(f http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Printf("%s - ", r.Method)
		f.ServeHTTP(w, r)
		fmt.Printf("%d ms\n", time.Since(now).Milliseconds())
	}
}
