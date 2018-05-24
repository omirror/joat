package middleware

import (
	"mime"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/ubiqueworks/joat/controller/webui"
)

func StaticAsset(timeout time.Duration) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" && r.Method != "HEAD" {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Allow", "GET, HEAD")
				w.Header().Set("Content-Length", "0")
				return
			}

			requestURI := r.RequestURI
			if strings.HasPrefix(requestURI, "/api") {
				next.ServeHTTP(w, r)
				return
			}

			staticPath := strings.TrimLeft(r.RequestURI, "/")
			if staticPath == "" {
				staticPath = "index.html"
			}

			log.Debug().Msgf("staticPath = %s", staticPath)

			if data, err := webui.Asset(staticPath); err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				//i, _ := webui.AssetInfo()
				//i.ModTime()
				mimeType := mime.TypeByExtension(filepath.Ext(staticPath))
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", mimeType)
				w.Header().Set("Content-Length", string(len(data)))
				w.Write(data)
			}
		}
		return http.HandlerFunc(fn)
	}
}
