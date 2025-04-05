package handler

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/application/usecase"
)

// TemperatureHandler lida com as requisições HTTP para obter temperatura por CEP
type TemperatureHandler struct {
	getTemperatureUseCase *usecase.GetTemperatureByZipCode
}

// NewTemperatureHandler cria uma nova instância do handler de temperatura
func NewTemperatureHandler(getTemperatureUseCase *usecase.GetTemperatureByZipCode) *TemperatureHandler {
	return &TemperatureHandler{
		getTemperatureUseCase: getTemperatureUseCase,
	}
}

// GetTemperatureByZipCode busca a temperatura a partir de um CEP
func (h *TemperatureHandler) GetTemperatureByZipCode(w http.ResponseWriter, r *http.Request) {
	// Obter o CEP da URL
	zipCode := r.URL.Query().Get("zipcode")
	if zipCode == "" {
		http.Error(w, "zipcode query parameter is required", http.StatusBadRequest)
		return
	}

	// Validar o formato do CEP
	isValid, err := validateZipCode(zipCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !isValid {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	// Executar o caso de uso
	weather, err := h.getTemperatureUseCase.Execute(r.Context(), zipCode)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	// Retornar o resultado como JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(weather); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// validateZipCode valida se o CEP tem o formato correto (8 dígitos)
func validateZipCode(zipCode string) (bool, error) {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(zipCode), nil
}
