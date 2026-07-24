package getresponse

import (
	"net/http"
	"strconv"
)

// ResponseHeader holds metadata parsed from a GetResponse API response,
// including pagination counters and rate-limit information.
//
// Numeric fields are best-effort: endpoints that do not return a given header
// (for example single-resource or POST endpoints omit pagination counters)
// leave the corresponding field at its zero value rather than erroring.
type ResponseHeader struct {
	ContentEncoding       string
	ContentSecurityPolicy string
	ContentType           string
	Date                  string
	TotalCount            int
	CurrentPage           int
	TotalPages            int
	RateLimit             int
	RateLimitRemaining    int
	UniqueID              string
}

// parseResponseHeaders extracts known headers from an API response. Missing or
// non-numeric counters are treated as zero so every endpoint can share one
// response path.
func parseResponseHeaders(h http.Header) ResponseHeader {
	atoi := func(key string) int {
		n, _ := strconv.Atoi(h.Get(key))
		return n
	}
	return ResponseHeader{
		ContentEncoding:       h.Get("content-encoding"),
		ContentSecurityPolicy: h.Get("content-security-policy"),
		ContentType:           h.Get("content-type"),
		Date:                  h.Get("date"),
		TotalCount:            atoi("totalcount"),
		CurrentPage:           atoi("currentpage"),
		TotalPages:            atoi("totalpages"),
		RateLimit:             atoi("x-ratelimit-limit"),
		RateLimitRemaining:    atoi("x-ratelimit-remaining"),
		UniqueID:              h.Get("x-unique-id"),
	}
}
