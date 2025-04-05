package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWeather(t *testing.T) {
	tests := []struct {
		name     string
		tempC    float64
		expected *Weather
	}{
		{
			name:  "deve criar uma nova instância com temperatura positiva",
			tempC: 25.5,
			expected: &Weather{
				TempC: 25.5,
			},
		},
		{
			name:  "deve criar uma nova instância com temperatura negativa",
			tempC: -10.0,
			expected: &Weather{
				TempC: -10.0,
			},
		},
		{
			name:  "deve criar uma nova instância com temperatura zero",
			tempC: 0.0,
			expected: &Weather{
				TempC: 0.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewWeather(tt.tempC)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestWeather_GetTempF(t *testing.T) {
	tests := []struct {
		name     string
		tempC    float64
		expected float64
	}{
		{
			name:     "deve converter 0°C para 32°F",
			tempC:    0.0,
			expected: 32.0,
		},
		{
			name:     "deve converter 100°C para 212°F",
			tempC:    100.0,
			expected: 212.0,
		},
		{
			name:     "deve converter -40°C para -40°F",
			tempC:    -40.0,
			expected: -40.0,
		},
		{
			name:     "deve converter temperatura decimal",
			tempC:    25.5,
			expected: 77.9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWeather(tt.tempC)
			got := w.GetTempF()
			assert.InDelta(t, tt.expected, got, 0.1)
		})
	}
}

func TestWeather_GetTempK(t *testing.T) {
	tests := []struct {
		name     string
		tempC    float64
		expected float64
	}{
		{
			name:     "deve converter 0°C para 273.15K",
			tempC:    0.0,
			expected: 273.15,
		},
		{
			name:     "deve converter 100°C para 373.15K",
			tempC:    100.0,
			expected: 373.15,
		},
		{
			name:     "deve converter -273.15°C para 0K",
			tempC:    -273.15,
			expected: 0.0,
		},
		{
			name:     "deve converter temperatura decimal",
			tempC:    25.5,
			expected: 298.65,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWeather(tt.tempC)
			got := w.GetTempK()
			assert.InDelta(t, tt.expected, got, 0.01)
		})
	}
}

func TestWeather_ToJSON(t *testing.T) {
	tests := []struct {
		name     string
		tempC    float64
		expected WeatherJSON
	}{
		{
			name:  "deve serializar temperatura positiva",
			tempC: 25.5,
			expected: WeatherJSON{
				TempC: 25.5,
				TempF: 77.9,
				TempK: 298.65,
			},
		},
		{
			name:  "deve serializar temperatura negativa",
			tempC: -10.0,
			expected: WeatherJSON{
				TempC: -10.0,
				TempF: 14.0,
				TempK: 263.15,
			},
		},
		{
			name:  "deve serializar temperatura zero",
			tempC: 0.0,
			expected: WeatherJSON{
				TempC: 0.0,
				TempF: 32.0,
				TempK: 273.15,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWeather(tt.tempC)
			got := w.ToJSON()

			// Verifica cada campo individualmente com margem de erro
			assert.InDelta(t, tt.expected.TempC, got.TempC, 0.01)
			assert.InDelta(t, tt.expected.TempF, got.TempF, 0.1)
			assert.InDelta(t, tt.expected.TempK, got.TempK, 0.01)

			// Testa a serialização JSON
			jsonGot, err := json.Marshal(got)
			assert.NoError(t, err)
			assert.NotEmpty(t, jsonGot)

			var unmarshalled WeatherJSON
			err = json.Unmarshal(jsonGot, &unmarshalled)
			assert.NoError(t, err)
			assert.InDelta(t, tt.expected.TempC, unmarshalled.TempC, 0.01)
			assert.InDelta(t, tt.expected.TempF, unmarshalled.TempF, 0.1)
			assert.InDelta(t, tt.expected.TempK, unmarshalled.TempK, 0.01)
		})
	}
}
