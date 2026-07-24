// Package getresponse is a Go client for the GetResponse v3 API.
//
// It supports the GetResponse SMB environment as well as the GetResponse MAX
// (formerly GetResponse 360) EU and US environments.
//
// # Getting started
//
// Create a client with an API key and a target environment:
//
//	client, err := getresponse.New(apiKey, getresponse.EnvMaxUS,
//		getresponse.WithMailingDomain("example.com"),
//		getresponse.WithTimeout(30*time.Second),
//	)
//	if err != nil {
//		// handle error
//	}
//
//	account, err := client.GetAccount(ctx)
//
// # Pagination
//
// List endpoints expose two shapes. The *List / Get* methods return a single
// page plus the parsed ResponseHeader (TotalPages, RateLimit, ...) so callers
// can page manually. The All* methods return a Go 1.23 range-over-func iterator
// that transparently walks every page:
//
//	for campaign, err := range client.AllCampaigns(ctx) {
//		if err != nil {
//			// handle error; the loop stops after yielding it
//			break
//		}
//		// use campaign
//	}
//
// # Errors
//
// Any non-2xx response is returned as an APIError carrying the structured
// error payload from GetResponse (code, message, context, ...). Use errors.As
// to inspect it.
package getresponse
