package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

// Estrutura para armazenar as informações de temperatura
type TemperatureResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

// Função principal
func main() {
	// Carregar as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Erro ao carregar o arquivo .env: %v\n", err)
		os.Exit(1)
	}

	http.HandleFunc("/weather", weatherHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server is listening on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
}

// Manipulador HTTP para o endpoint /weather
func weatherHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cep := r.URL.Query().Get("cep")
	if len(cep) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	city, err := getCityFromCEP(cep)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	tempC, err := getTemperature(city)
	if err != nil {
		http.Error(w, "could not fetch weather data", http.StatusInternalServerError)
		return
	}

	tempF := tempC*1.8 + 32
	tempK := tempC + 273

	response := TemperatureResponse{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Função para obter o nome da cidade a partir do CEP
func getCityFromCEP(cep string) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unable to fetch data for the given CEP")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data struct {
		Localidade string `json:"localidade"`
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	if data.Localidade == "" {
		return "", fmt.Errorf("city not found for the given CEP")
	}

	return data.Localidade, nil
}

// Função para obter a temperatura em graus Celsius da cidade
func getTemperature(city string) (float64, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("Weather API key not set")
		return 0, fmt.Errorf("Weather API key not set")
	}

	// Fazer o encoding do nome da cidade para incluir caracteres especiais corretamente
	encodedCity := url.QueryEscape(city)

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&lang=pt", apiKey, encodedCity)
	fmt.Printf("Requesting weather data from: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to make request: %v\n", err)
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Failed response body: %s\n", string(body))
		return 0, fmt.Errorf("unable to fetch weather data, status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		return 0, err
	}

	var data struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Failed to unmarshal response: %v\n", err)
		return 0, err
	}

	return data.Current.TempC, nil
}

// Função para converter a temperatura para Fahrenheit e Kelvin
func convertTemperature(tempC float64) (float64, float64) {
	tempF := tempC*1.8 + 32
	tempK := tempC + 273
	return tempF, tempK
}
