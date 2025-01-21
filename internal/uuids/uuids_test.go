package uuids_test

import (
	"testing"

	"github.com/jtprogru/slicesvsmap/internal/uuids"
)

func TestGetUUID(t *testing.T) {
	uuid1 := uuids.GetUUID()
	uuid2 := uuids.GetUUID()

	if uuid1 == "" {
		t.Errorf("expected non-empty UUID, got empty string")
	}

	if uuid2 == "" {
		t.Errorf("expected non-empty UUID, got empty string")
	}

	if uuid1 == uuid2 {
		t.Errorf("expected different UUIDs, got the same: %s", uuid1)
	}
}
