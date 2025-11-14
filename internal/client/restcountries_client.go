package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type RestCountriesClient interface {
	GetCountryByName(ctx context.Context, name string) (*CountryResponse, error)
}

type restCountriesClient struct {
	httpClient *http.Client
	baseURL    string
}

type CountryResponse struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Capital    []string               `json:"capital"`
	Currencies map[string]interface{} `json:"currencies"`
	Population int64                  `json:"population"`
}

func NewRestCountriesClient() RestCountriesClient {
	return &restCountriesClient{
		httpClient: &http.Client{Timeout: 5 * time.Second},
		baseURL:    "https://restcountries.com/v3.1",
	}
}

func (c *restCountriesClient) GetCountryByName(ctx context.Context, name string) (*CountryResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/name/%s", c.baseURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status %d from API", resp.StatusCode)
	}

	var data []CountryResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("no countries found")
	}

	for _, country := range data {
		if strings.EqualFold(country.Name.Common, name) {
			return &country, nil
		}
	}

	return &data[0], nil
}
