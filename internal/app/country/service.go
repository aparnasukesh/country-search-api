package country

import (
	"context"

	"github.com/aparnasukesh/country-search-api/internal/cache"
	"github.com/aparnasukesh/country-search-api/internal/client"
)

type CountryService struct {
	cache  cache.Cache
	client client.RestCountriesClient
}

func NewCountryService(cache cache.Cache, client client.RestCountriesClient) *CountryService {
	return &CountryService{
		cache:  cache,
		client: client,
	}
}

func (s *CountryService) GetCountry(ctx context.Context, name string) (*Country, error) {
	if val, ok := s.cache.Get(name); ok {
		country := val.(Country)
		return &country, nil
	}

	cResp, err := s.client.GetCountryByName(ctx, name)
	if err != nil {
		return nil, err
	}

	var currencySymbol string
	for _, v := range cResp.Currencies {
		if currencyMap, ok := v.(map[string]interface{}); ok {
			if sym, ok := currencyMap["symbol"].(string); ok {
				currencySymbol = sym
				break
			}
		}
	}

	country := Country{
		Name:       cResp.Name.Common,
		Capital:    getFirstCapital(cResp.Capital),
		Currency:   currencySymbol,
		Population: cResp.Population,
	}

	s.cache.Set(name, country)
	return &country, nil
}

func getFirstCapital(caps []string) string {
	if len(caps) > 0 {
		return caps[0]
	}
	return ""
}
