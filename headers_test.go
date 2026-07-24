package getresponse

import (
	"net/http"
	"testing"
)

func TestParseResponseHeaders(t *testing.T) {
	h := http.Header{}
	h.Set("TotalCount", "42")
	h.Set("CurrentPage", "2")
	h.Set("TotalPages", "5")
	h.Set("X-RateLimit-Limit", "30000")
	h.Set("X-RateLimit-Remaining", "29999")
	h.Set("X-Unique-Id", "abc123")

	got := parseResponseHeaders(h)
	if got.TotalCount != 42 || got.CurrentPage != 2 || got.TotalPages != 5 {
		t.Errorf("pagination parsed wrong: %+v", got)
	}
	if got.RateLimit != 30000 || got.RateLimitRemaining != 29999 {
		t.Errorf("rate limit parsed wrong: %+v", got)
	}
	if got.UniqueID != "abc123" {
		t.Errorf("UniqueID = %q, want abc123", got.UniqueID)
	}
}

func TestParseResponseHeadersLenient(t *testing.T) {
	// Missing/non-numeric counters must default to zero, not error, so
	// single-resource and POST endpoints can share the response path.
	h := http.Header{}
	h.Set("Content-Type", "application/json")

	got := parseResponseHeaders(h)
	if got.TotalCount != 0 || got.CurrentPage != 0 || got.TotalPages != 0 {
		t.Errorf("missing counters should be zero: %+v", got)
	}
	if got.ContentType != "application/json" {
		t.Errorf("ContentType = %q, want application/json", got.ContentType)
	}
}
