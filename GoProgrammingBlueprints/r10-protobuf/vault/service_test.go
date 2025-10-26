package vault

import (
	"context"
	"testing"
)

func TestHasherService(t *testing.T) {
	srv := NewService()
	ctx := context.Background()
	h, err := srv.Hash(ctx, "pass")
	if err != nil {
		t.Errorf("Hash: %s", err)
	}
	ok, err := srv.Validate(ctx, "pass", h)
	if err != nil {
		t.Errorf("Valid: %s", err)
	}
	if !ok {
		t.Error("Method Validate should have returned true!")
	}
	ok, err = srv.Validate(ctx, "wrong_pass", h)
	if err == nil {
		t.Errorf("Validation of hash agains different password should return error: %s", err)
	}
	if ok {
		t.Error("Method Validate should have returned false!")
	}
}
