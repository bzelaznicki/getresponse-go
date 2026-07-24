package getresponse

import (
	"context"
	"fmt"
	"net/http"
	"slices"
)

// Import represents the state and statistics of an import operation.
type Import struct {
	ImportID   string   `json:"importId"`
	Campaign   Campaign `json:"campaign"`
	Status     string   `json:"status"`
	Statistics struct {
		Uploaded    int `json:"uploaded"`
		Invalid     int `json:"invalid"`
		Updated     int `json:"updated"`
		AddedToList int `json:"addedToList"`
	} `json:"statistics"`
	ErrorStatistics struct {
		SyntaxErrors       int `json:"syntaxErrors"`
		AleradyInQueue     int `json:"alreadyInQueue"`
		InvalidDomains     int `json:"invalidDomains"`
		Blacklist          int `json:"blacklist"`
		PolicyFailures     int `json:"policyFailures"`
		MismatchedCriteria int `json:"mismatchedCriteria"`
	} `json:"errorStatistics"`

	CreatedOn  string `json:"createdOn"`
	FinishedOn string `json:"finishedOn"`
	Href       string `json:"href"`
}

// ScheduleImport describes a batch of contacts to import into a campaign.
//
// FieldMapping names the column for each position in every Contacts row and
// must include "email". See
// https://www.getresponse.com/help/how-do-i-prepare-a-file-for-import.html
type ScheduleImport struct {
	CampaignID   string     `json:"campaignId"`
	FieldMapping []string   `json:"fieldMapping"`
	Contacts     [][]string `json:"contacts"`
}

// ScheduleNewContactImport validates and submits a contact import. Each contact
// row must have exactly as many columns as FieldMapping, which must include an
// "email" field.
func (c *Client) ScheduleNewContactImport(ctx context.Context, imp ScheduleImport) (Import, error) {
	if len(imp.Contacts) == 0 {
		return Import{}, fmt.Errorf("getresponse: contacts cannot be empty")
	}
	if len(imp.FieldMapping) == 0 {
		return Import{}, fmt.Errorf("getresponse: field mapping cannot be empty")
	}
	if !slices.Contains(imp.FieldMapping, "email") {
		return Import{}, fmt.Errorf("getresponse: field mapping must contain an email field")
	}
	for _, row := range imp.Contacts {
		if len(row) != len(imp.FieldMapping) {
			return Import{}, fmt.Errorf("getresponse: each contact row must match the number of mapped fields")
		}
	}

	created, _, err := do[Import](ctx, c, http.MethodPost, "imports", nil, imp)
	return created, err
}

// GetImport retrieves the status of a previously scheduled import.
func (c *Client) GetImport(ctx context.Context, importID string) (Import, error) {
	path := fmt.Sprintf("imports/%s", importID)
	imp, _, err := do[Import](ctx, c, http.MethodGet, path, nil, nil)
	return imp, err
}
