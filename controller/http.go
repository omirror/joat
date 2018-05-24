package controller

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	chimw "github.com/go-chi/chi/middleware"
	"github.com/ubiqueworks/joat/internal/util/httpio"
)

func configureRouter(c *controller) chi.Router {
	router := chi.NewRouter()

	router.Use(chimw.RequestID)
	router.Use(chimw.RealIP)
	router.Use(chimw.Logger)
	router.Use(chimw.Recoverer)
	router.Use(chimw.Timeout(60 * time.Second))

	//router.Use(middleware.StaticAsset())
	router.Get("/health", healthCheck())
	router.Get("/api/members", clusterMembers(c))
	return router
}

func healthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func clusterMembers(c *controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := chi.URLParam(r, "role")
		httpio.SendJSON(w, http.StatusOK, c.members(role))
	}
}
