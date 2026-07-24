# getresponse-go

A Go client for the [GetResponse v3 API](https://apidocs.getresponse.com/v3).

Supports the GetResponse **SMB** environment as well as GetResponse **MAX**
(formerly GetResponse 360) **EU** and **US** environments.

```bash
go get github.com/bzelaznicki/getresponse-go
```

Requires Go 1.24+.

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	getresponse "github.com/bzelaznicki/getresponse-go"
)

func main() {
	client, err := getresponse.New("YOUR_API_KEY", getresponse.EnvMaxUS,
		getresponse.WithMailingDomain("example.com"),
		getresponse.WithTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	account, err := client.GetAccount(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Logged in as", account.Email)
}
```

## Environments

| Constant             | Code  | Host                             |
| -------------------- | ----- | -------------------------------- |
| `getresponse.EnvSMB` | `SMB` | `api.getresponse.com`            |
| `getresponse.EnvMaxEU` | `PL`  | `api3.getresponse360.pl`         |
| `getresponse.EnvMaxUS` | `US`  | `api3.getresponse360.com`        |

`ParseEnvironment("US")` converts a stored short code into an `Environment`
(returns an error on unknown codes — no silent fallback).

## Pagination

Each list endpoint comes in two flavours:

- `Get*List` / `Get*` — returns a single page plus a `ResponseHeader`
  (`TotalPages`, `TotalCount`, `RateLimit`, ...) for manual paging.
- `All*` — a Go 1.23 [range-over-func](https://go.dev/ref/spec#For_range)
  iterator that transparently walks every page.

```go
for campaign, err := range client.AllCampaigns(ctx) {
	if err != nil {
		log.Fatal(err) // loop stops after yielding the error
	}
	fmt.Println(campaign.Name)
}
```

Break out of the loop at any time and no further pages are fetched.

## Errors

Any non-2xx response is returned as an `APIError` carrying the structured
payload from GetResponse:

```go
_, err := client.ScheduleNewContactImport(ctx, imp)
var apiErr getresponse.APIError
if errors.As(err, &apiErr) {
	fmt.Println(apiErr.Code, apiErr.Message)
}
```

## Options

| Option                          | Purpose                                            |
| ------------------------------- | -------------------------------------------------- |
| `WithMailingDomain(domain)`     | Set `X-Domain` (required for MAX/360 accounts)     |
| `WithTimeout(d)`                | Timeout for the default HTTP client                |
| `WithHTTPClient(hc)`            | Supply a custom `*http.Client`                     |
| `WithUserAgent(ua)`             | Override the `User-Agent` header                   |
| `WithEndpoint(url)`             | Override the base URL (testing / mock servers)     |

## License

MIT — see [LICENSE](LICENSE).
