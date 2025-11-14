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
func TestCacheGetMissing(t *testing.T) {
	c := NewCache()

	if _, ok := c.Get("Unknown"); ok {
		t.Errorf("Expected ok=false for missing key")
	}
}
func TestCacheConcurrentAccess(t *testing.T) {
	c := NewCache()

	for i := 0; i < 100; i++ {
		go c.Set("key", i)
		go c.Get("key")
	}
}
