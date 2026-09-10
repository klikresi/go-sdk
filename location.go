package klikresi

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// Location provides access to the location API.
type Location struct {
	client *Client
}

type pageOptions struct {
	limit  int
	cursor string
}

// PageOption configures a paginated location request.
type PageOption func(*pageOptions)

// WithLimit sets the number of items per page. The API default is 50.
func WithLimit(limit int) PageOption {
	return func(o *pageOptions) {
		o.limit = limit
	}
}

// WithCursor sets the pagination cursor returned by the previous page.
func WithCursor(cursor string) PageOption {
	return func(o *pageOptions) {
		o.cursor = cursor
	}
}

func applyPageOptions(opts []PageOption) pageOptions {
	o := pageOptions{}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

func pageQuery(o pageOptions, extra url.Values) url.Values {
	query := url.Values{}
	for k, vs := range extra {
		query[k] = vs
	}
	if o.limit > 0 {
		query.Set("limit", strconv.Itoa(o.limit))
	}
	if o.cursor != "" {
		query.Set("cursor", o.cursor)
	}
	return query
}

// Search looks up locations matching the given keyword. Billed per request
// (Location API, Rp 1 per request).
func (l *Location) Search(ctx context.Context, keyword string, opts ...PageOption) (*LocationPage, error) {
	query := pageQuery(applyPageOptions(opts), url.Values{"keyword": {keyword}})
	var page LocationPage
	if err := l.client.do(ctx, http.MethodGet, "/api/locations", query, nil, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

// Provinces returns one page of provinces. Billed per request (Location
// API, Rp 1 per request).
func (l *Location) Provinces(ctx context.Context, opts ...PageOption) (*ProvincePage, error) {
	query := pageQuery(applyPageOptions(opts), nil)
	var page ProvincePage
	if err := l.client.do(ctx, http.MethodGet, "/api/provinces", query, nil, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

// Cities returns one page of cities within the given province. Billed per
// request (Location API, Rp 1 per request).
func (l *Location) Cities(ctx context.Context, provinceID string, opts ...PageOption) (*CityPage, error) {
	query := pageQuery(applyPageOptions(opts), url.Values{"province_id": {provinceID}})
	var page CityPage
	if err := l.client.do(ctx, http.MethodGet, "/api/cities", query, nil, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

// Districts returns one page of districts within the given city. Billed
// per request (Location API, Rp 1 per request).
func (l *Location) Districts(ctx context.Context, cityID string, opts ...PageOption) (*DistrictPage, error) {
	query := pageQuery(applyPageOptions(opts), url.Values{"city_id": {cityID}})
	var page DistrictPage
	if err := l.client.do(ctx, http.MethodGet, "/api/districts", query, nil, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

// AllLocations returns every location matching the keyword, following
// pagination cursors automatically.
func (l *Location) AllLocations(ctx context.Context, keyword string, opts ...PageOption) ([]LocationInfo, error) {
	var all []LocationInfo
	cursor := ""
	for {
		page, err := l.Search(ctx, keyword, append(opts, WithCursor(cursor))...)
		if err != nil {
			return nil, err
		}
		all = append(all, page.Data...)
		if page.NextCursor == "" || page.NextCursor == cursor {
			return all, nil
		}
		cursor = page.NextCursor
	}
}

// AllProvinces returns every province, following pagination cursors
// automatically.
func (l *Location) AllProvinces(ctx context.Context, opts ...PageOption) ([]Province, error) {
	var all []Province
	cursor := ""
	for {
		page, err := l.Provinces(ctx, append(opts, WithCursor(cursor))...)
		if err != nil {
			return nil, err
		}
		all = append(all, page.Data...)
		if page.NextCursor == "" || page.NextCursor == cursor {
			return all, nil
		}
		cursor = page.NextCursor
	}
}

// AllCities returns every city within the given province, following
// pagination cursors automatically.
func (l *Location) AllCities(ctx context.Context, provinceID string, opts ...PageOption) ([]City, error) {
	var all []City
	cursor := ""
	for {
		page, err := l.Cities(ctx, provinceID, append(opts, WithCursor(cursor))...)
		if err != nil {
			return nil, err
		}
		all = append(all, page.Data...)
		if page.NextCursor == "" || page.NextCursor == cursor {
			return all, nil
		}
		cursor = page.NextCursor
	}
}

// AllDistricts returns every district within the given city, following
// pagination cursors automatically.
func (l *Location) AllDistricts(ctx context.Context, cityID string, opts ...PageOption) ([]District, error) {
	var all []District
	cursor := ""
	for {
		page, err := l.Districts(ctx, cityID, append(opts, WithCursor(cursor))...)
		if err != nil {
			return nil, err
		}
		all = append(all, page.Data...)
		if page.NextCursor == "" || page.NextCursor == cursor {
			return all, nil
		}
		cursor = page.NextCursor
	}
}
