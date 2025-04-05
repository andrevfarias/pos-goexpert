package entity

// Weather representa as temperaturas em diferentes unidades
type Weather struct {
	TempC float64 `json:"temp_c"` // Temperatura em Celsius
}

// NewWeather cria uma nova instância de Weather
func NewWeather(tempC float64) *Weather {
	return &Weather{
		TempC: tempC,
	}
}

// GetTempF retorna a temperatura em Fahrenheit
func (w *Weather) GetTempF() float64 {
	return w.TempC*9.0/5.0 + 32.0
}

// GetTempK retorna a temperatura em Kelvin
func (w *Weather) GetTempK() float64 {
	return w.TempC + 273.15
}

// ToJSON retorna uma estrutura com todas as temperaturas para serialização
type WeatherJSON struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

// ToJSON converte Weather para uma estrutura JSON
func (w *Weather) ToJSON() WeatherJSON {
	return WeatherJSON{
		TempC: w.TempC,
		TempF: w.GetTempF(),
		TempK: w.GetTempK(),
	}
}
