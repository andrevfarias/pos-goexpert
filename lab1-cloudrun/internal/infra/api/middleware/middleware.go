package middleware

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Logger é um middleware que registra informações sobre a requisição
func Logger(out io.Writer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Criar um ResponseWriter personalizado para capturar o status
			rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}

			// Executar o próximo handler
			next.ServeHTTP(rw, r)

			// Registrar informações da requisição
			fmt.Fprintf(out, "%s %s %d %v\n",
				r.Method,
				r.URL.Path,
				rw.status,
				time.Since(start),
			)
		})
	}
}

// Recoverer é um middleware que recupera de pânicos
func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Internal Server Error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// responseWriter é um wrapper para http.ResponseWriter que captura o status code
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if rw.status == 0 {
		rw.status = http.StatusOK
	}
	return rw.ResponseWriter.Write(b)
}
