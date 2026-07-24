package getresponse

import (
	"context"
	"net/http"
)

// Account represents a GetResponse account and its profile details.
type Account struct {
	AccountID   string `json:"accountId"`
	Email       string `json:"email"`
	CountryCode struct {
		CountryCodeID string `json:"countryCodeId"`
		CountryCode   string `json:"countryCode"`
	} `json:"countryCode"`
	IndustryTag struct {
		IndustryTagID string `json:"industryTagId"`
	} `json:"industryTag"`
	TimeZone struct {
		Name   string `json:"name"`
		Offset string `json:"offset"`
	} `json:"timeZone"`
	Href              string `json:"href"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	CompanyName       string `json:"companyName"`
	Phone             string `json:"phone"`
	State             string `json:"state"`
	City              string `json:"city"`
	Street            string `json:"street"`
	ZipCode           string `json:"zipCode"`
	NumberOfEmployees string `json:"numberOfEmployees"`
	TimeFormat        string `json:"timeFormat"`
}

// GetAccount retrieves the details of the authenticated account.
func (c *Client) GetAccount(ctx context.Context) (Account, error) {
	account, _, err := do[Account](ctx, c, http.MethodGet, "accounts", nil, nil)
	return account, err
}
