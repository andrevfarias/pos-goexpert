package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerMiddleware(t *testing.T) {
	// Configurar captura de logs
	var buf bytes.Buffer

	// Criar handler de teste
	handler := Logger(&buf)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))

	// Criar request de teste
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	// Executar o middleware
	handler.ServeHTTP(w, req)

	// Verificar se o log foi registrado
	logOutput := buf.String()
	assert.Contains(t, logOutput, "GET")
	assert.Contains(t, logOutput, "/test")
	assert.Contains(t, logOutput, "200")

	// Verificar a resposta
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", strings.TrimSpace(w.Body.String()))
}

func TestRecovererMiddleware(t *testing.T) {
	// Criar handler que causa pânico
	handler := Recoverer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("teste de pânico")
	}))

	// Criar request de teste
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	// Executar o middleware
	handler.ServeHTTP(w, req)

	// Verificar se o pânico foi recuperado e retornou 500
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "Internal Server Error", strings.TrimSpace(w.Body.String()))
}

func TestCombinedMiddlewares(t *testing.T) {
	// Configurar captura de logs
	var buf bytes.Buffer

	// Criar handler com ambos os middlewares
	handler := Logger(&buf)(
		Recoverer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				panic("teste de pânico com log")
			}),
		),
	)

	// Criar request de teste
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	// Executar os middlewares
	handler.ServeHTTP(w, req)

	// Verificar se o log foi registrado e o pânico foi recuperado
	logOutput := buf.String()
	assert.Contains(t, logOutput, "GET")
	assert.Contains(t, logOutput, "/test")
	assert.Contains(t, logOutput, "500")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "Internal Server Error", strings.TrimSpace(w.Body.String()))
}
