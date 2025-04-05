package entity

// Weather representa as temperaturas em diferentes unidades
type Weather struct {
	TempC float64 `json:"temp_c"` // Temperatura em Celsius
	TempF float64 `json:"temp_f"` // Temperatura em Fahrenheit
	TempK float64 `json:"temp_k"` // Temperatura em Kelvin
}
