// Package klikresi provides a client for the Klik Resi API.
//
// It supports account lookup, shipment tracking, shipping rate
// calculation, and location lookup (search, provinces, cities, districts)
// across Indonesian couriers.
//
// Create a client with NewClient and use the resources grouped under
// Tracking, Rates, Location, and Me:
//
//	client := klikresi.NewClient("your-api-key")
//	tracking, err := client.Tracking.Get(ctx, "YOUR-AWB", klikresi.CourierJNE)
package klikresi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	// DefaultBaseURL is the production API endpoint.
	DefaultBaseURL = "https://klikresi.com"

	// BaseURLEnv overrides DefaultBaseURL when set. It is mainly useful
	// for tests and proxies.
	BaseURLEnv = "KLIKRESI_BASE_URL"

	// DefaultTimeout is applied to every request when the client is
	// created with NewClient.
	DefaultTimeout = 30 * time.Second
)

// Client is a Klik Resi API client. It is safe for concurrent use.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client

	// Tracking provides access to the tracking API.
	Tracking *Tracking
	// Rates provides access to the rates API.
	Rates *Rates
	// Location provides access to the location API.
	Location *Location
	// Me provides access to the account profile API.
	Me *Me
}

// NewClient returns a Client for the given API key. The base URL defaults
// to https://klikresi.com and can be overridden with the KLIKRESI_BASE_URL
// environment variable.
func NewClient(apiKey string) *Client {
	baseURL := DefaultBaseURL
	if v := os.Getenv(BaseURLEnv); v != "" {
		baseURL = v
	}
	c := &Client{
		apiKey:     apiKey,
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: DefaultTimeout},
	}
	c.Tracking = &Tracking{client: c}
	c.Rates = &Rates{client: c}
	c.Location = &Location{client: c}
	c.Me = &Me{client: c}
	return c
}

// do performs an HTTP request against the API and decodes the JSON
// response body into out.
func (c *Client) do(ctx context.Context, method, path string, query url.Values, body any, out any) error {
	u := c.baseURL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}

	var reader io.Reader
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("klikresi: marshal request body: %w", err)
		}
		reader = bytes.NewReader(buf)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, reader)
	if err != nil {
		return fmt.Errorf("klikresi: create request: %w", err)
	}
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("klikresi: %s %s: %w", method, path, err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("klikresi: read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return newAPIError(resp.StatusCode, data)
	}

	if out == nil || len(data) == 0 {
		return nil
	}
	if err := json.Unmarshal(data, out); err != nil {
		return fmt.Errorf("klikresi: decode response: %w", err)
	}
	return nil
}

// dataEnvelope wraps the { "data": ... } envelope used by most endpoints.
type dataEnvelope[T any] struct {
	Data T `json:"data"`
}

// APIError is returned for any non-2xx API response.
type APIError struct {
	// StatusCode is the HTTP status code returned by the API.
	StatusCode int
	// Message is the error message returned by the API, when available.
	Message string
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("klikresi: %s (HTTP %d)", e.Message, e.StatusCode)
	}
	return fmt.Sprintf("klikresi: request failed with HTTP status %d", e.StatusCode)
}

// newAPIError builds an APIError from an HTTP status code and response
// body, extracting the message when the body contains one.
func newAPIError(statusCode int, body []byte) *APIError {
	err := &APIError{StatusCode: statusCode}
	var payload struct {
		Message string `json:"message"`
	}
	if json.Unmarshal(body, &payload) == nil && payload.Message != "" {
		err.Message = payload.Message
	}
	return err
}
