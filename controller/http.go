package controller

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/log"
	"github.com/ubiqueworks/joat/util"
)

func configureRouter(c *controller) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/", healthCheck())
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
		sendJSON(w, http.StatusOK, c.members(role))
	}
}

func sendHttpError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func sendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	body, err := util.MarshalJSON(data)
	if err != nil {
		log.Error().Err(err).Msg("error while marshalling JSON data")
		sendHttpError(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.Header().Add("Content-Length", string(len(body)))
	w.Write(body)
}
