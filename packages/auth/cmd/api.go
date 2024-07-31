package cmd

import (
	"fmt"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	InitValidator()
	router := http.NewServeMux()

	router.HandleFunc("GET /", defaultHandler)
	router.HandleFunc("GET /v1/validate", validateTokenHandler)
	router.HandleFunc("POST /v1/login", loginHandler)
	router.HandleFunc("POST /v1/register", registerHandler)

	server := http.Server{
		Addr:    s.addr,
		Handler: LoggerMiddleware(GatewayAuth(router)),
	}

	return server.ListenAndServe()
}

func LoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[LOG] %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func GatewayAuth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("API-Authorization")

		_, err := DecodeGatewayToken(token)
		if err != nil {
			fmt.Printf("%v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
