package klikresi

import (
	"context"
	"net/http"
)

// Rates provides access to the rates API.
type Rates struct {
	client *Client
}

type rateRequest struct {
	OriginID              string   `json:"origin_id,omitempty"`
	DestinationID         string   `json:"destination_id,omitempty"`
	Origin                string   `json:"origin,omitempty"`
	Destination           string   `json:"destination,omitempty"`
	OriginPostalCode      int      `json:"origin_postal_code,omitempty"`
	DestinationPostalCode int      `json:"destination_postal_code,omitempty"`
	Weight                float64  `json:"weight"`
	Couriers              []string `json:"couriers,omitempty"`
}

func (r *Rates) calculate(ctx context.Context, req rateRequest) (*RateResult, error) {
	var env dataEnvelope[RateResult]
	if err := r.client.do(ctx, http.MethodPost, "/api/rates", nil, req, &env); err != nil {
		return nil, err
	}
	return &env.Data, nil
}

// CalculateByID calculates shipping rates using district IDs as the origin
// and destination. Couriers optionally filters the result to the given
// courier codes.
func (r *Rates) CalculateByID(ctx context.Context, originID, destinationID string, weight float64, couriers ...string) (*RateResult, error) {
	return r.calculate(ctx, rateRequest{
		OriginID:      originID,
		DestinationID: destinationID,
		Weight:        weight,
		Couriers:      couriers,
	})
}

// CalculateByName calculates shipping rates using location names as the
// origin and destination.
func (r *Rates) CalculateByName(ctx context.Context, origin, destination string, weight float64) (*RateResult, error) {
	return r.calculate(ctx, rateRequest{
		Origin:      origin,
		Destination: destination,
		Weight:      weight,
	})
}

// CalculateByPostalCode calculates shipping rates using postal codes as
// the origin and destination.
func (r *Rates) CalculateByPostalCode(ctx context.Context, originPostalCode, destinationPostalCode int, weight float64) (*RateResult, error) {
	return r.calculate(ctx, rateRequest{
		OriginPostalCode:      originPostalCode,
		DestinationPostalCode: destinationPostalCode,
		Weight:                weight,
	})
}
