package cache

import (
	"testing"
)

func TestCacheSetAndGet(t *testing.T) {
	c := NewCache()
	c.Set("India", "Delhi")

	if val, ok := c.Get("India"); !ok || val != "Delhi" {
		t.Errorf("Expected Delhi, got %v", val)
	}
}
