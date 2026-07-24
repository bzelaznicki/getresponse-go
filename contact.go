package getresponse

import (
	"context"
	"fmt"
	"net/http"
)

// Contact represents a GetResponse contact (subscriber).
type Contact struct {
	ContactID       string   `json:"contactId"`
	Name            string   `json:"name"`
	Origin          string   `json:"origin"`
	TimeZone        string   `json:"timeZone"`
	Activities      string   `json:"activities"`
	ChangedOn       string   `json:"changedOn"`
	CreatedOn       string   `json:"createdOn"`
	Campaign        Campaign `json:"campaign"`
	Email           string   `json:"email"`
	DayOfCycle      string   `json:"dayOfCycle"`
	Scoring         int      `json:"scoring"`
	EngagementScore int      `json:"engagementScore"`
	Href            string   `json:"href"`
	IpAddress       string   `json:"ipAddress"`
	Geolocation     struct {
		Latitude      string `json:"latitude"`
		Longitude     string `json:"longitude"`
		ContinentCode string `json:"continentCode"`
		CountryCode   string `json:"countryCode"`
		Region        string `json:"region"`
		PostalCode    string `json:"postalCode"`
		DmaCode       string `json:"dmaCode"`
		City          string `json:"city"`
	} `json:"geolocation,omitzero"`

	Tags              []Tag              `json:"tags,omitzero"`
	CustomFieldValues []CustomFieldValue `json:"customFieldValues,omitzero"`
}

// GetContactsFromCampaign retrieves a single page of contacts belonging to the
// given campaign. Use AllContactsFromCampaign to iterate every page.
func (c *Client) GetContactsFromCampaign(ctx context.Context, campaignID string, opts ...QueryOption) ([]Contact, ResponseHeader, error) {
	path := fmt.Sprintf("campaigns/%s/contacts", campaignID)
	return do[[]Contact](ctx, c, http.MethodGet, path, buildQuery(opts...), nil)
}
