package getresponse

import (
	"context"
	"fmt"
	"net/http"
)

// SearchContacts describes the criteria for a contact search or a saved segment.
//
// Name is required when creating a segment with NewSearchContacts but optional
// when searching ad hoc with GetContactsFromSearchContactsConditions.
type SearchContacts struct {
	SubscribersType      []string                `json:"subscribersType"`
	SectionLogicOperator string                  `json:"sectionLogicOperator"`
	Section              []SearchContactsSection `json:"section"`
	SearchContactId      string                  `json:"searchContactId,omitempty"`
	Name                 string                  `json:"name,omitempty"`
	CreatedOn            string                  `json:"createdOn,omitempty"`
	Href                 string                  `json:"href,omitempty"`
}

// SearchContactsSection is one section of search criteria combined via the
// parent's SectionLogicOperator.
type SearchContactsSection struct {
	CampaignIdsList  []string `json:"campaignIdsList"`
	LogicOperator    string   `json:"logicOperator"`
	SubscriberCycle  []string `json:"subscriberCycle"`
	SubscriptionDate string   `json:"subscriptionDate"`
	Conditions       []any    `json:"conditions"`
}

// GetSearchContactsByID retrieves a saved search (segment) by its ID.
func (c *Client) GetSearchContactsByID(ctx context.Context, searchContactsID string, opts ...QueryOption) (SearchContacts, ResponseHeader, error) {
	path := fmt.Sprintf("search-contacts/%s", searchContactsID)
	return do[SearchContacts](ctx, c, http.MethodGet, path, buildQuery(opts...), nil)
}

// NewSearchContacts creates (saves) a new search-contacts segment.
func (c *Client) NewSearchContacts(ctx context.Context, conditions SearchContacts) (SearchContacts, error) {
	created, _, err := do[SearchContacts](ctx, c, http.MethodPost, "search-contacts", nil, conditions)
	return created, err
}

// GetSearchContactsList retrieves a single page of saved searches (segments).
// Use AllSearchContacts to iterate every page.
func (c *Client) GetSearchContactsList(ctx context.Context, opts ...QueryOption) ([]SearchContacts, ResponseHeader, error) {
	return do[[]SearchContacts](ctx, c, http.MethodGet, "search-contacts", buildQuery(opts...), nil)
}

// GetContactsByIDSearchContacts retrieves a single page of contacts matching a
// saved search (segment). Use AllContactsByIDSearchContacts to iterate every
// page.
func (c *Client) GetContactsByIDSearchContacts(ctx context.Context, searchContactsID string, opts ...QueryOption) ([]Contact, ResponseHeader, error) {
	path := fmt.Sprintf("search-contacts/%s/contacts", searchContactsID)
	return do[[]Contact](ctx, c, http.MethodGet, path, buildQuery(opts...), nil)
}

// GetContactsFromSearchContactsConditions retrieves a single page of contacts
// matching the supplied ad-hoc conditions. Use
// AllContactsFromSearchContactsConditions to iterate every page.
func (c *Client) GetContactsFromSearchContactsConditions(ctx context.Context, conditions SearchContacts, opts ...QueryOption) ([]Contact, ResponseHeader, error) {
	return do[[]Contact](ctx, c, http.MethodPost, "search-contacts/contacts", buildQuery(opts...), conditions)
}
