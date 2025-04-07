package handler

import (
	"encoding/json"
	"net/http"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/entity"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/usecase"
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
	rawZipCode := r.PathValue("cep")
	zipCode, err := entity.NewZipCode(rawZipCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	weather, err := h.getTemperatureUseCase.Execute(r.Context(), zipCode)
	if err != nil {
		switch err {
		case usecase.ErrInvalidZipCode:
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		case usecase.ErrZipCodeNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Retornar o resultado como JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(weather); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
