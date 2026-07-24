package getresponse

import (
	"net/url"
	"strconv"
)

// QueryOption customizes the query string of a list request.
type QueryOption func(url.Values)

// WithPage sets the "page" query parameter.
func WithPage(page int) QueryOption {
	return func(v url.Values) { v.Set("page", strconv.Itoa(page)) }
}

// WithPerPage sets the "perPage" query parameter.
func WithPerPage(perPage int) QueryOption {
	return func(v url.Values) { v.Set("perPage", strconv.Itoa(perPage)) }
}

// WithName sets the "query[name]" filter parameter.
func WithName(name string) QueryOption {
	return func(v url.Values) { v.Set("query[name]", name) }
}

// WithFields sets the "fields" parameter, limiting the fields returned by the
// API (e.g. "name,email").
func WithFields(fields string) QueryOption {
	return func(v url.Values) { v.Set("fields", fields) }
}

// buildQuery applies the supplied options and returns the encoded values.
func buildQuery(opts ...QueryOption) url.Values {
	values := url.Values{}
	for _, opt := range opts {
		opt(values)
	}
	return values
}
