package main

import (
	"github.com/aparnasukesh/country-search-api/config"
	"github.com/aparnasukesh/country-search-api/internal/boot"
	"github.com/aparnasukesh/country-search-api/internal/di"
)

func main() {
	cfg := config.LoadConfig()
	container := di.InitializeContainer()
	boot.StartServer(cfg, container)
}
