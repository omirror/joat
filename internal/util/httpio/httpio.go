package httpio

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/ubiqueworks/joat/internal/util/jsonx"
)

func SendHttpError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func SendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	body, err := jsonx.MarshalJSON(data)
	if err != nil {
		log.Error().Err(err).Msg("error while marshalling JSON data")
		SendHttpError(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.Header().Add("Content-Length", string(len(body)))
	w.Write(body)
}
