package entity

// Address representa o endereço obtido a partir de um CEP
type Address struct {
	ZipCode      string `json:"zipcode"`
	Street       string `json:"street"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}
