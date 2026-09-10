package klikresi

import (
	"context"
	"net/http"
	"testing"
)

func TestLocationSearch(t *testing.T) {
	var gotKeyword, gotLimit, gotCursor string
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		gotKeyword = q.Get("keyword")
		gotLimit = q.Get("limit")
		gotCursor = q.Get("cursor")
		writeFixture(t, w, "locations.json")
	})

	page, err := c.Location.Search(context.Background(), "depok", WithLimit(10), WithCursor(""))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotKeyword != "depok" {
		t.Errorf("keyword = %q", gotKeyword)
	}
	if gotLimit != "10" {
		t.Errorf("limit = %q", gotLimit)
	}
	if gotCursor != "" {
		t.Errorf("cursor = %q", gotCursor)
	}
	if len(page.Data) != 10 {
		t.Fatalf("data = %d, want 10", len(page.Data))
	}
	if page.Data[0].ID != "32.09.31" || page.Data[0].District != "Depok" || page.Data[0].Province != "Jawa Barat" {
		t.Errorf("data[0] = %+v", page.Data[0])
	}
	if page.NextCursor != "32.76.09" {
		t.Errorf("next_cursor = %q", page.NextCursor)
	}
}

func TestLocationProvinces(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		writeFixture(t, w, "provinces.json")
	})

	page, err := c.Location.Provinces(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(page.Data) != 10 {
		t.Fatalf("data = %d, want 10", len(page.Data))
	}
	if page.Data[0].ID != "11" || page.Data[0].Name != "Aceh" {
		t.Errorf("data[0] = %+v", page.Data[0])
	}
	if page.NextCursor != "21" {
		t.Errorf("next_cursor = %q", page.NextCursor)
	}
}

func TestLocationCities(t *testing.T) {
	var gotProvinceID string
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotProvinceID = r.URL.Query().Get("province_id")
		writeFixture(t, w, "cities.json")
	})

	page, err := c.Location.Cities(context.Background(), "33")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotProvinceID != "33" {
		t.Errorf("province_id = %q", gotProvinceID)
	}
	if len(page.Data) != 10 || page.Data[7].Name != "Kabupaten Magelang" {
		t.Errorf("data = %+v", page.Data)
	}
}

func TestLocationDistricts(t *testing.T) {
	var gotCityID string
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotCityID = r.URL.Query().Get("city_id")
		writeFixture(t, w, "districts.json")
	})

	page, err := c.Location.Districts(context.Background(), "33.08")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotCityID != "33.08" {
		t.Errorf("city_id = %q", gotCityID)
	}
	if len(page.Data) != 10 || page.Data[9].Name != "Mertoyudan" {
		t.Errorf("data = %+v", page.Data)
	}
}

func TestLocationAllLocations(t *testing.T) {
	calls := 0
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.URL.Query().Get("cursor") == "" {
			writeFixture(t, w, "locations_page1.json")
		} else {
			writeFixture(t, w, "locations_page2.json")
		}
	})

	all, err := c.Location.AllLocations(context.Background(), "depok")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if calls != 2 {
		t.Errorf("requests = %d, want 2", calls)
	}
	if len(all) != 3 {
		t.Fatalf("results = %d, want 3", len(all))
	}
	if all[0].ID != "32.09.31" || all[2].ID != "32.76.09" {
		t.Errorf("results = %+v", all)
	}
}
