package country

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	service *CountryService
}

func NewCountryHandler(service *CountryService) *CountryHandler {
	return &CountryHandler{service: service}
}

func (h *CountryHandler) SearchCountry(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(400, gin.H{"error": "missing 'name' parameter"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	country, err := h.service.GetCountry(ctx, name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, country)
}
