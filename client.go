package getresponse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const defaultTimeout = 30 * time.Second

// Client is a GetResponse API client. Create one with New. A Client is safe for
// concurrent use by multiple goroutines.
type Client struct {
	apiKey        string
	endpoint      string
	mailingDomain string
	userAgent     string
	httpClient    *http.Client
}

// Option configures a Client in New.
type Option func(*Client)

// WithMailingDomain sets the X-Domain header, required for GetResponse MAX
// (360) accounts.
func WithMailingDomain(domain string) Option {
	return func(c *Client) { c.mailingDomain = domain }
}

// WithTimeout sets the timeout on the default HTTP client. It has no effect if
// a custom client is supplied via WithHTTPClient.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		if c.httpClient == nil {
			c.httpClient = &http.Client{}
		}
		c.httpClient.Timeout = d
	}
}

// WithHTTPClient supplies a custom *http.Client (for custom transports,
// proxies, or instrumentation).
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) { c.httpClient = hc }
}

// WithUserAgent overrides the default User-Agent header.
func WithUserAgent(ua string) Option {
	return func(c *Client) { c.userAgent = ua }
}

// WithEndpoint overrides the base API URL derived from the environment. The URL
// must end with a trailing slash. Primarily useful for testing against a mock
// server.
func WithEndpoint(raw string) Option {
	return func(c *Client) {
		if !strings.HasSuffix(raw, "/") {
			raw += "/"
		}
		c.endpoint = raw
	}
}

// New creates a Client for the given API key and environment. The environment
// determines the base API URL; supply additional configuration via options.
func New(apiKey string, env Environment, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("getresponse: API key is missing")
	}
	if !env.valid() {
		return nil, fmt.Errorf("getresponse: invalid environment %q", env)
	}

	c := &Client{
		apiKey:    apiKey,
		endpoint:  env.endpoint(),
		userAgent: UserAgent,
	}
	for _, opt := range opts {
		opt(c)
	}
	if c.httpClient == nil {
		c.httpClient = &http.Client{Timeout: defaultTimeout}
	}
	return c, nil
}

// do performs an API request and decodes the JSON response into T. Any non-2xx
// response is returned as an APIError. body, when non-nil, is JSON-encoded as
// the request payload.
func do[T any](ctx context.Context, c *Client, method, path string, query url.Values, body any) (T, ResponseHeader, error) {
	var zero T

	u := c.endpoint + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}

	var reader io.Reader
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return zero, ResponseHeader{}, fmt.Errorf("getresponse: marshal request body: %w", err)
		}
		reader = bytes.NewReader(buf)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, reader)
	if err != nil {
		return zero, ResponseHeader{}, fmt.Errorf("getresponse: create request: %w", err)
	}
	req.Header.Set("X-Auth-Token", "api-key "+c.apiKey)
	req.Header.Set("X-Domain", c.mailingDomain)
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return zero, ResponseHeader{}, fmt.Errorf("getresponse: request failed: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return zero, ResponseHeader{}, fmt.Errorf("getresponse: read response: %w", err)
	}

	headers := parseResponseHeaders(resp.Header)

	if resp.StatusCode >= 400 {
		return zero, headers, parseAPIError(resp.StatusCode, raw)
	}

	var out T
	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &out); err != nil {
			return zero, headers, fmt.Errorf("getresponse: decode response: %w", err)
		}
	}
	return out, headers, nil
}

// parseAPIError decodes a GetResponse error payload, falling back to the raw
// body when the response is not a structured APIError.
func parseAPIError(status int, body []byte) error {
	var apiErr APIError
	if err := json.Unmarshal(body, &apiErr); err == nil && (apiErr.Code != 0 || apiErr.Message != "" || apiErr.HttpStatus != 0) {
		if apiErr.HttpStatus == 0 {
			apiErr.HttpStatus = status
		}
		return apiErr
	}
	return fmt.Errorf("getresponse: API error (HTTP %d): %s", status, strings.TrimSpace(string(body)))
}
