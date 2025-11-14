package country

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockCountryService struct {
	result *Country
	err    error
}

func (m *mockCountryService) GetCountry(ctx context.Context, name string) (*Country, error) {
	return m.result, m.err
}
func TestSearchCountry_MissingName(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &mockCountryService{}
	handler := NewCountryHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/countries/search", nil)
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler.SearchCountry(c)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}
func TestSearchCountry_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &mockCountryService{
		result: &Country{
			Name:       "India",
			Capital:    "New Delhi",
			Currency:   "₹",
			Population: 1380004385,
		},
		err: nil,
	}

	handler := NewCountryHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/countries/search?name=India", nil)
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler.SearchCountry(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	expectedBody := `{"name":"India","capital":"New Delhi","currency":"₹","population":1380004385}`
	if w.Body.String() != expectedBody {
		t.Errorf("expected %s, got %s", expectedBody, w.Body.String())
	}
}
func TestSearchCountry_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &mockCountryService{
		result: nil,
		err:    errors.New("service failure"),
	}

	handler := NewCountryHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/countries/search?name=India", nil)
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler.SearchCountry(c)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", w.Code)
	}
}
