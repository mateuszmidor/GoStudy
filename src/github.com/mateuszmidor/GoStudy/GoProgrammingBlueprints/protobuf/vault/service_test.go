package vault

import (
	context "context"
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
		t.Error("Method Valid should have returned true!")
	}
	ok, err = src.Validate(ctx, "wrong_pass", h)
	if err != nil {
		t.Errorf("Valid: %s", err)
	}
	if ok {
		t.Error("Method Valid should have returned false!")
	}
}
