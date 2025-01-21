package main_test

import (
	"testing"

	"github.com/jtprogru/slicesvsmap/internal/mapped"
	"github.com/jtprogru/slicesvsmap/internal/sliced"
	uuids "github.com/jtprogru/slicesvsmap/internal/uuids"
)

func TestCacheSliced_Add(t *testing.T) {
	cache := sliced.New()
	cache.Add(uuids.GetUUID())
	if len(cache.IDs) != 1 {
		t.Errorf("expected 1 id, got %d", len(cache.IDs))
	}
}

func TestCacheMapped_Add(t *testing.T) {
	cache := mapped.New()
	cache.Add(uuids.GetUUID())
	if len(cache.IDs) != 1 {
		t.Errorf("expected 1 id, got %d", len(cache.IDs))
	}
}

func BenchmarkCacheSliced_Add(b *testing.B) {
	cache := sliced.New()
	for i := 0; i < b.N; i++ {
		cache.Add(uuids.GetUUID())
	}
}

func BenchmarkCacheMapped_Add(b *testing.B) {
	cache := mapped.New()
	for i := 0; i < b.N; i++ {
		cache.Add(uuids.GetUUID())
	}
}
