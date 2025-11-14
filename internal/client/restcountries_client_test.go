package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type mockCountry struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Capital    []string               `json:"capital"`
	Currencies map[string]interface{} `json:"currencies"`
	Population int64                  `json:"population"`
}

func TestGetCountryByName_Success(t *testing.T) {

	mockResponse := []mockCountry{
		{
			Capital:    []string{"Delhi"},
			Currencies: map[string]interface{}{"INR": map[string]string{"name": "Indian Rupee"}},
			Population: 1400000000,
		},
	}
	mockResponse[0].Name.Common = "India"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	client := &restCountriesClient{
		httpClient: &http.Client{},
		baseURL:    server.URL,
	}

	ctx := context.Background()

	country, err := client.GetCountryByName(ctx, "India")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if country.Name.Common != "India" {
		t.Errorf("expected India, got %s", country.Name.Common)
	}
	if country.Capital[0] != "Delhi" {
		t.Errorf("expected capital Delhi, got %v", country.Capital)
	}
	if country.Population != 1400000000 {
		t.Errorf("expected population 1400000000, got %d", country.Population)
	}
}

func TestGetCountryByName_NotFound(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]mockCountry{})
	}))
	defer server.Close()

	client := &restCountriesClient{
		httpClient: &http.Client{},
		baseURL:    server.URL,
	}

	_, err := client.GetCountryByName(context.Background(), "UnknownCountry")
	if err == nil {
		t.Fatal("expected error for no countries, got nil")
	}
}

func TestGetCountryByName_Non200Status(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := &restCountriesClient{
		httpClient: &http.Client{},
		baseURL:    server.URL,
	}

	_, err := client.GetCountryByName(context.Background(), "India")
	if err == nil {
		t.Fatal("expected error for API status != 200")
	}
}

func TestGetCountryByName_ContextTimeout(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
	}))
	defer server.Close()

	client := &restCountriesClient{
		httpClient: &http.Client{},
		baseURL:    server.URL,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := client.GetCountryByName(ctx, "India")
	if err == nil {
		t.Fatal("expected timeout error, got nil")
	}
}
