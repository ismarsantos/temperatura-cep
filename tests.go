package main

import (
	"testing"
)

// Classe de teste para testar as principais funções do projeto
func TestGetCityFromCEP(t *testing.T) {
	city, err := getCityFromCEP("01001000")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if city == "" {
		t.Fatalf("expected a valid city name, got empty string")
	}
}

func TestGetTemperature(t *testing.T) {
	tempC, err := getTemperature("Sao Paulo")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if tempC == 0 {
		t.Fatalf("expected a valid temperature, got 0")
	}
}

func TestConvertTemperature(t *testing.T) {
	tempC := 25.0
	tempF, tempK := convertTemperature(tempC)
	if tempF != 77.0 {
		t.Fatalf("expected 77.0, got %v", tempF)
	}
	if tempK != 298.0 {
		t.Fatalf("expected 298.0, got %v", tempK)
	}
}
