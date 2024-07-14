package cmd

import (
	"encoding/json"
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
	router := http.NewServeMux()

	router.HandleFunc("GET /", defaultHandler)
	router.HandleFunc("GET /login", loginHandler)

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

func WriteJson(w http.ResponseWriter, value interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(value)
}

func ReadJson(w http.ResponseWriter, r *http.Request, value *any) error {
	if err := json.NewDecoder(r.Body).Decode(&value); err != nil {
		http.Error(w, "Coundn't parse request body", http.StatusBadRequest)
		return err
	}

	return nil
}
