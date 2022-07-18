package middleware

import "net/http"

type CrosMiddleware struct {
}

func NewCrosMiddleware() *CrosMiddleware {
	return &CrosMiddleware{}
}

func (m *CrosMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,X-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		next(w, r)
	}
}
