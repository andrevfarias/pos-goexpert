package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAddress(t *testing.T) {
	tests := []struct {
		name         string
		zipCode      string
		street       string
		complement   string
		neighborhood string
		city         string
		state        string
		expectedErr  error
	}{
		{
			name:         "deve criar um endereço válido",
			zipCode:      "12345678",
			street:       "Rua Teste",
			complement:   "Apto 123",
			neighborhood: "Centro",
			city:         "São Paulo",
			state:        "SP",
			expectedErr:  nil,
		},
		{
			name:         "deve criar um endereço válido com CEP formatado",
			zipCode:      "12345-678",
			street:       "Rua Teste",
			complement:   "Apto 123",
			neighborhood: "Centro",
			city:         "São Paulo",
			state:        "SP",
			expectedErr:  nil,
		},
		{
			name:         "deve retornar erro para CEP inválido",
			zipCode:      "123456",
			street:       "Rua Teste",
			complement:   "Apto 123",
			neighborhood: "Centro",
			city:         "São Paulo",
			state:        "SP",
			expectedErr:  ErrInvalidZipCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAddress(tt.zipCode, tt.street, tt.complement, tt.neighborhood, tt.city, tt.state)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
				assert.Nil(t, got)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, got)
			assert.Equal(t, tt.zipCode, got.ZipCode)
			assert.Equal(t, tt.street, got.Street)
			assert.Equal(t, tt.complement, got.Complement)
			assert.Equal(t, tt.neighborhood, got.Neighborhood)
			assert.Equal(t, tt.city, got.City)
			assert.Equal(t, tt.state, got.State)
		})
	}
}

func TestAddress_IsValidZipCode(t *testing.T) {
	tests := []struct {
		name     string
		zipCode  string
		expected bool
	}{
		{
			name:     "deve validar CEP com 8 dígitos",
			zipCode:  "12345678",
			expected: true,
		},
		{
			name:     "deve validar CEP formatado",
			zipCode:  "12345-678",
			expected: true,
		},
		{
			name:     "deve invalidar CEP com menos dígitos",
			zipCode:  "1234567",
			expected: false,
		},
		{
			name:     "deve invalidar CEP com mais dígitos",
			zipCode:  "123456789",
			expected: false,
		},
		{
			name:     "deve invalidar CEP com letras",
			zipCode:  "1234567a",
			expected: false,
		},
		{
			name:     "deve invalidar CEP vazio",
			zipCode:  "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{ZipCode: tt.zipCode}
			got := a.IsValidZipCode()
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestAddress_GetCleanZipCode(t *testing.T) {
	tests := []struct {
		name     string
		zipCode  string
		expected string
	}{
		{
			name:     "deve limpar CEP formatado",
			zipCode:  "12345-678",
			expected: "12345678",
		},
		{
			name:     "deve manter CEP já limpo",
			zipCode:  "12345678",
			expected: "12345678",
		},
		{
			name:     "deve lidar com CEP vazio",
			zipCode:  "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{ZipCode: tt.zipCode}
			got := a.GetCleanZipCode()
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestAddress_FormatZipCode(t *testing.T) {
	tests := []struct {
		name     string
		zipCode  string
		expected string
	}{
		{
			name:     "deve formatar CEP limpo",
			zipCode:  "12345678",
			expected: "12345-678",
		},
		{
			name:     "deve manter CEP já formatado",
			zipCode:  "12345-678",
			expected: "12345-678",
		},
		{
			name:     "deve retornar CEP inválido sem formatação",
			zipCode:  "1234567",
			expected: "1234567",
		},
		{
			name:     "deve lidar com CEP vazio",
			zipCode:  "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{ZipCode: tt.zipCode}
			got := a.FormatZipCode()
			assert.Equal(t, tt.expected, got)
		})
	}
}
