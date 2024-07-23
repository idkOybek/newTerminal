package utils

import (
	"net/http"
	"time"

	"github.com/idkOybek/internal/logger"
)

// LoggerMiddleware логирует все входящие HTTP запросы и ответы
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := &responseWriter{w, http.StatusOK}
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				logger.ErrorLogger.Printf(
					"[PANIC RECOVER] %s %s %s %d %s %v",
					r.Method,
					r.RequestURI,
					r.RemoteAddr,
					http.StatusInternalServerError,
					time.Since(start),
					err,
				)
			}
		}()

		next.ServeHTTP(ww, r)
		logger.InfoLogger.Printf(
			"%s %s %s %d %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			ww.status,
			time.Since(start),
		)
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}
