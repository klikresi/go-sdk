package klikresi

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"
)

func TestTrackingGet(t *testing.T) {
	var gotPath, gotKey string
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotKey = r.Header.Get("x-api-key")
		writeFixture(t, w, "tracking_success.json")
	})

	got, err := c.Tracking.Get(context.Background(), "1234567890", CourierJNE)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotPath != "/api/trackings/1234567890/couriers/jne" {
		t.Errorf("path = %q", gotPath)
	}
	if gotKey != "test-api-key" {
		t.Errorf("x-api-key = %q", gotKey)
	}
	if got.Status != StatusDelivered {
		t.Errorf("status = %q", got.Status)
	}
	if got.Origin.ContactName != "DAMAS AMIRUL KARIM" {
		t.Errorf("origin contact = %q", got.Origin.ContactName)
	}
	if got.Destination.Address != "JALAN PAKUAN TIMUR RT1RW1 DESA" {
		t.Errorf("destination address = %q", got.Destination.Address)
	}
	if len(got.Histories) != 11 {
		t.Fatalf("histories = %d, want 11", len(got.Histories))
	}
	want := time.Date(2025, 4, 23, 12, 24, 0, 0, time.FixedZone("", 7*3600))
	if !got.Histories[0].Date.Equal(want) {
		t.Errorf("history[0] date = %v, want %v", got.Histories[0].Date, want)
	}
	if got.Histories[0].Status != StatusDelivered {
		t.Errorf("history[0] status = %q", got.Histories[0].Status)
	}
	if got.Histories[1].Status != StatusInTransit {
		t.Errorf("history[1] status = %q", got.Histories[1].Status)
	}
}

func TestTrackingGetWithNumber(t *testing.T) {
	var gotQuery string
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.RawQuery
		writeFixture(t, w, "tracking_success.json")
	})

	_, err := c.Tracking.Get(context.Background(), "1234567890", CourierIDExpress, WithNumber("08123456789"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotQuery != "number=08123456789" {
		t.Errorf("query = %q", gotQuery)
	}
}

func TestTrackingGetFailed(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(loadFixture(t, "tracking_failed.json"))
	})

	_, err := c.Tracking.Get(context.Background(), "1234567890", CourierSPX)
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("err = %v, want *APIError", err)
	}
	if apiErr.StatusCode != http.StatusBadRequest {
		t.Errorf("status code = %d, want 400", apiErr.StatusCode)
	}
	want := "Failed to get tracking information. It's either invalid or expired. Please check again"
	if apiErr.Message != want {
		t.Errorf("message = %q, want %q", apiErr.Message, want)
	}
}
