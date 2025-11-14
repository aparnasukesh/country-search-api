package country

import (
	"context"
	"errors"
	"testing"

	"github.com/aparnasukesh/country-search-api/internal/client"
)

type mockCache struct {
	store map[string]interface{}
}

func newMockCache() *mockCache {
	return &mockCache{store: make(map[string]interface{})}
}

func (m *mockCache) Get(key string) (interface{}, bool) {
	v, ok := m.store[key]
	return v, ok
}

func (m *mockCache) Set(key string, value interface{}) {
	m.store[key] = value
}

type mockRestClient struct {
	response *client.CountryResponse
	err      error
	called   bool
}

func (m *mockRestClient) GetCountryByName(ctx context.Context, name string) (*client.CountryResponse, error) {
	m.called = true
	return m.response, m.err
}
func TestGetCountry_ReturnsFromCache(t *testing.T) {
	mockCache := newMockCache()
	mockClient := &mockRestClient{}

	expected := Country{
		Name:       "India",
		Capital:    "Delhi",
		Currency:   "₹",
		Population: 1400000000,
	}

	mockCache.Set("India", expected)

	service := NewCountryService(mockCache, mockClient)

	result, err := service.GetCountry(context.Background(), "India")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Name != expected.Name {
		t.Errorf("expected name %s, got %s", expected.Name, result.Name)
	}

	if mockClient.called {
		t.Errorf("client should NOT have been called when data exists in cache")
	}
}
func TestGetCountry_FetchesFromClientAndCaches(t *testing.T) {
	mockCache := newMockCache()

	mockResponse := &client.CountryResponse{
		Population: 1400000000,
		Capital:    []string{"Delhi"},
		Currencies: map[string]interface{}{
			"INR": map[string]interface{}{
				"symbol": "₹",
			},
		},
	}
	mockResponse.Name.Common = "India"

	mockClient := &mockRestClient{response: mockResponse}

	service := NewCountryService(mockCache, mockClient)

	result, err := service.GetCountry(context.Background(), "India")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Currency != "₹" {
		t.Errorf("expected currency ₹, got %s", result.Currency)
	}

	if result.Capital != "Delhi" {
		t.Errorf("expected capital Delhi, got %s", result.Capital)
	}

	if _, ok := mockCache.Get("India"); !ok {
		t.Errorf("expected data to be stored in cache")
	}

	if !mockClient.called {
		t.Errorf("expected client call, got none")
	}
}
func TestGetCountry_ClientError(t *testing.T) {
	mockCache := newMockCache()
	mockClient := &mockRestClient{
		err: errors.New("API failure"),
	}

	service := NewCountryService(mockCache, mockClient)

	_, err := service.GetCountry(context.Background(), "India")
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
func TestGetCountry_NoCapital(t *testing.T) {
	mockCache := newMockCache()

	mockResponse := &client.CountryResponse{
		Population: 1000,
		Capital:    []string{},
		Currencies: map[string]interface{}{},
	}
	mockResponse.Name.Common = "TestLand"

	mockClient := &mockRestClient{response: mockResponse}

	service := NewCountryService(mockCache, mockClient)

	result, err := service.GetCountry(context.Background(), "TestLand")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Capital != "" {
		t.Errorf("expected empty capital, got %s", result.Capital)
	}
}
func TestGetCountry_NoCurrencySymbol(t *testing.T) {
	mockCache := newMockCache()

	mockResponse := &client.CountryResponse{
		Population: 5000,
		Capital:    []string{"XCity"},
		Currencies: map[string]interface{}{
			"AAA": map[string]interface{}{
				"name": "TestCurrency",
			},
		},
	}
	mockResponse.Name.Common = "XLand"

	mockClient := &mockRestClient{response: mockResponse}

	service := NewCountryService(mockCache, mockClient)

	result, err := service.GetCountry(context.Background(), "XLand")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Currency != "" {
		t.Errorf("expected empty currency symbol, got %s", result.Currency)
	}
}
