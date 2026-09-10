package klikresi

import (
	"context"
	"net/http"
)

// Me provides access to the account profile API.
type Me struct {
	client *Client
}

// Get returns the profile of the account that owns the API key.
func (m *Me) Get(ctx context.Context) (*AccountProfile, error) {
	var env dataEnvelope[AccountProfile]
	if err := m.client.do(ctx, http.MethodGet, "/api/me", nil, nil, &env); err != nil {
		return nil, err
	}
	return &env.Data, nil
}
