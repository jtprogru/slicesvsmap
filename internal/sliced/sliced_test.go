package sliced_test

import (
	"testing"

	"github.com/jtprogru/slicesvsmap/internal/sliced"
	"github.com/jtprogru/slicesvsmap/internal/uuids"
)

func TestCacheSliced_Add(t *testing.T) {
	cache := sliced.New()
	key := uuids.GetUUID()
	cache.Add(key)
	if len(cache.IDs) != 1 {
		t.Errorf("expected 1 id, got %d", len(cache.IDs))
	}
}

func TestCacheSliced_AddDuplicate(t *testing.T) {
	cache := sliced.New()
	key := uuids.GetUUID()
	cache.Add(key)
	cache.Add(key)
	if len(cache.IDs) != 1 {
		t.Errorf("expected 1 id after adding duplicate, got %d", len(cache.IDs))
	}
}

func TestCacheSliced_Contains(t *testing.T) {
	cache := sliced.New()
	key := uuids.GetUUID()
	cache.Add(key)
	if !cache.Contains(key) {
		t.Errorf("expected cache to contain key %s", key)
	}
}

func BenchmarkCacheSliced_Add(b *testing.B) {
	cache := sliced.New()
	for i := 0; i < b.N; i++ {
		cache.Add(uuids.GetUUID())
	}
}
