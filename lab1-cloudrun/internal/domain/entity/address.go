package entity

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidZipCode = errors.New("CEP inválido")
)

// Address representa o endereço obtido a partir de um CEP
type Address struct {
	ZipCode      string `json:"zipcode"`
	Street       string `json:"street"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

// NewAddress cria uma nova instância de Address
func NewAddress(zipCode, street, complement, neighborhood, city, state string) (*Address, error) {
	a := &Address{
		ZipCode:      zipCode,
		Street:       street,
		Complement:   complement,
		Neighborhood: neighborhood,
		City:         city,
		State:        state,
	}

	if err := a.Validate(); err != nil {
		return nil, err
	}

	return a, nil
}

// Validate valida o endereço
func (a *Address) Validate() error {
	if !a.IsValidZipCode() {
		return ErrInvalidZipCode
	}
	return nil
}

// IsValidZipCode verifica se o CEP é válido
func (a *Address) IsValidZipCode() bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(a.GetCleanZipCode())
}

// GetCleanZipCode retorna o CEP sem formatação
func (a *Address) GetCleanZipCode() string {
	return strings.ReplaceAll(a.ZipCode, "-", "")
}

// FormatZipCode retorna o CEP formatado (00000-000)
func (a *Address) FormatZipCode() string {
	clean := a.GetCleanZipCode()
	if len(clean) != 8 {
		return clean
	}
	return clean[:5] + "-" + clean[5:]
}
