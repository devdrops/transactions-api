package app

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	// Basic router
	r := chi.NewRouter()

	// go-chi useful middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(commonHeadersMiddleware)

	// Routes: healthcheck
	r.Get("/health", HealthCheck)
	// Routes: Accounts
	r.Route("/accounts", func(r chi.Router) {
		r.Post("/", CreateAccount)
		r.Get("/{accountId}", GetAccount)
	})
	// Routes: Transactions
	r.Post("/transactions", CreateTransaction)

	return r
}

// getURLParam is used to read a value from the URL, as a string.
func getURLParam(r *http.Request, v string) string {
	return chi.URLParam(r, v)
}

// commonHeadersMiddleware adds a group of headers for every HTTP response.
func commonHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Server", "transactions-api")
		next.ServeHTTP(w, r)
	})
}
