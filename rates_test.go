package klikresi

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestRatesCalculateByID(t *testing.T) {
	var gotBody map[string]any
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/rates" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("content-type = %q", ct)
		}
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		writeFixture(t, w, "rates_by_id.json")
	})

	got, err := c.Rates.CalculateByID(context.Background(), "33.08.20", "32.09.31", 1, CourierJNE)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotBody["origin_id"] != "33.08.20" {
		t.Errorf("origin_id = %v", gotBody["origin_id"])
	}
	if gotBody["destination_id"] != "32.09.31" {
		t.Errorf("destination_id = %v", gotBody["destination_id"])
	}
	if gotBody["weight"] != float64(1) {
		t.Errorf("weight = %v", gotBody["weight"])
	}
	couriers, ok := gotBody["couriers"].([]any)
	if !ok || len(couriers) != 1 || couriers[0] != "jne" {
		t.Errorf("couriers = %v", gotBody["couriers"])
	}
	if got.Origin.ID != "33.08.20" {
		t.Errorf("origin id = %q", got.Origin.ID)
	}
	if got.Destination.Name != "Depok, Kabupaten Cirebon, Jawa Barat" {
		t.Errorf("destination name = %q", got.Destination.Name)
	}
	if len(got.Pricing) != 7 {
		t.Fatalf("pricing = %d, want 7", len(got.Pricing))
	}
	if got.Pricing[4].CourierCode != "jne" || got.Pricing[4].Price != 21000 || got.Pricing[4].Duration != "3-6 Hari" {
		t.Errorf("pricing[4] = %+v", got.Pricing[4])
	}
	if got.Pricing[6].Type != "cargo" || got.Pricing[6].Service != "JTR" {
		t.Errorf("pricing[6] = %+v", got.Pricing[6])
	}
}

func TestRatesCalculateByName(t *testing.T) {
	var gotBody map[string]any
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		writeFixture(t, w, "rates_by_name.json")
	})

	got, err := c.Rates.CalculateByName(context.Background(), "Secang, Kabupaten Magelang, Jawa Tengah", "Depok, Kabupaten Cirebon, Jawa Barat", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotBody["origin"] != "Secang, Kabupaten Magelang, Jawa Tengah" {
		t.Errorf("origin = %v", gotBody["origin"])
	}
	if gotBody["destination"] != "Depok, Kabupaten Cirebon, Jawa Barat" {
		t.Errorf("destination = %v", gotBody["destination"])
	}
	if len(got.Pricing) != 10 {
		t.Errorf("pricing = %d, want 10", len(got.Pricing))
	}
}

func TestRatesCalculateByPostalCode(t *testing.T) {
	var gotBody map[string]any
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		writeFixture(t, w, "rates_custom.json")
	})

	got, err := c.Rates.CalculateByPostalCode(context.Background(), 56195, 45155, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotBody["origin_postal_code"] != float64(56195) {
		t.Errorf("origin_postal_code = %v", gotBody["origin_postal_code"])
	}
	if gotBody["destination_postal_code"] != float64(45155) {
		t.Errorf("destination_postal_code = %v", gotBody["destination_postal_code"])
	}
	if len(got.Pricing) != 2 {
		t.Errorf("pricing = %d, want 2", len(got.Pricing))
	}
}
