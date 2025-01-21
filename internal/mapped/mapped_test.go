package mapped_test

import (
	"testing"

	"github.com/jtprogru/slicesvsmap/internal/mapped"
	"github.com/jtprogru/slicesvsmap/internal/uuids"
)

func TestCacheMapped_Add(t *testing.T) {
	cache := mapped.New()
	key := uuids.GetUUID()
	cache.Add(key)
	if len(cache.IDs) != 1 {
		t.Errorf("expected 1 id, got %d", len(cache.IDs))
	}
}

func TestCacheMapped_AddDuplicate(t *testing.T) {
	cache := mapped.New()
	key := uuids.GetUUID()
	cache.Add(key)
	cache.Add(key)
	if len(cache.IDs) != 1 {
		t.Errorf("expected 1 id after adding duplicate, got %d", len(cache.IDs))
	}
}

func TestCacheMapped_Contains(t *testing.T) {
	cache := mapped.New()
	key := uuids.GetUUID()
	cache.Add(key)
	if !cache.Contains(key) {
		t.Errorf("expected cache to contain key %s", key)
	}
}

func BenchmarkCacheMapped_Add(b *testing.B) {
	cache := mapped.New()
	for i := 0; i < b.N; i++ {
		cache.Add(uuids.GetUUID())
	}
}
