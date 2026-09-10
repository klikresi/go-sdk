package klikresi

import (
	"context"
	"net/http"
	"testing"
)

func TestMeGet(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/api/me" {
			t.Errorf("path = %q, want /api/me", r.URL.Path)
		}
		writeFixture(t, w, "me.json")
	})

	me, err := c.Me.Get(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if me.ID != "1" || me.Name != "PT Contoh" {
		t.Errorf("me = %+v", me)
	}
	if me.Email != "user@example.com" {
		t.Errorf("email = %q", me.Email)
	}
	if me.Balance != 250000 {
		t.Errorf("balance = %v, want 250000", me.Balance)
	}
}
