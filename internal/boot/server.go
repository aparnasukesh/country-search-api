package boot

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aparnasukesh/country-search-api/config"
	"github.com/aparnasukesh/country-search-api/internal/di"
	"github.com/gin-gonic/gin"
)

func StartServer(cfg *config.Config, container *di.Container) {

	router := gin.Default()

	router.GET("/api/countries/search", container.CountryHandler.SearchCountry)

	server := &http.Server{
		Addr:    cfg.ServerPort,
		Handler: router,
	}

	go func() {
		fmt.Println("Server running at http://localhost" + cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server Error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\n Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Shutdown Error:", err)
	}

	fmt.Println("Server stopped gracefully.")
}
