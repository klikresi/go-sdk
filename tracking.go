package klikresi

import (
	"context"
	"net/http"
	"net/url"
)

// Tracking provides access to the tracking API.
type Tracking struct {
	client *Client
}

type trackingOptions struct {
	number string
}

// TrackingOption configures a tracking request.
type TrackingOption func(*trackingOptions)

// WithNumber adds the `number` query parameter. It is required by
// ID Express (courier "ide") and is passed through for every courier
// whenever it is provided.
func WithNumber(number string) TrackingOption {
	return func(o *trackingOptions) {
		o.number = number
	}
}

// Get returns the tracking information for the given tracking number and
// courier code. Tracking is charged only for successful requests.
func (t *Tracking) Get(ctx context.Context, trackingNumber, courierCode string, opts ...TrackingOption) (*TrackingInfo, error) {
	o := trackingOptions{}
	for _, opt := range opts {
		opt(&o)
	}

	query := url.Values{}
	if o.number != "" {
		query.Set("number", o.number)
	}

	path := "/api/trackings/" + url.PathEscape(trackingNumber) + "/couriers/" + url.PathEscape(courierCode)
	var env dataEnvelope[TrackingInfo]
	if err := t.client.do(ctx, http.MethodGet, path, query, nil, &env); err != nil {
		return nil, err
	}
	return &env.Data, nil
}
