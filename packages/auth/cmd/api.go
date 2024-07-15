package cmd

import (
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
	router.HandleFunc("GET /v1/login", loginHandler)
	router.HandleFunc("GET /v1/register", registerHandler)

	server := http.Server{
		Addr:    s.addr,
		Handler: LoggerMiddleware(router),
	}

	return server.ListenAndServe()
}

func LoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[LOG] %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
