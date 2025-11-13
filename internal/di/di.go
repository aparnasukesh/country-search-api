package di

import (
	"github.com/aparnasukesh/country-search-api/internal/app/country"
	"github.com/aparnasukesh/country-search-api/internal/cache"
	"github.com/aparnasukesh/country-search-api/internal/client"
)

type Container struct {
	CountryHandler *country.CountryHandler
}

func InitializeContainer() *Container {
	cache := cache.NewCache()
	client := client.NewRestCountriesClient()
	service := country.NewCountryService(cache, client)
	handler := country.NewCountryHandler(service)

	return &Container{
		CountryHandler: handler,
	}
}
