package getresponse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// newTestClient returns a client pointed at the given test server.
func newTestClient(t *testing.T, url string) *Client {
	t.Helper()
	c, err := New("test-key", EnvMaxUS,
		WithMailingDomain("example.com"),
		WithEndpoint(url),
	)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return c
}

func TestScheduleNewContactImport_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		_ = json.NewEncoder(w).Encode(APIError{
			HttpStatus:      429,
			Code:            1015,
			CodeDescription: "Too many requests",
			Message:         "You have reached your requests limit",
			MoreInfo:        "https://apidocs.getresponse.com/en/v3/errors/1015",
			Context:         map[string]any{"timeToReset": "100 seconds"},
			Uuid:            "510c6726-7f65-46b7-a798-ca403133924f",
		})
	}))
	defer server.Close()

	c := newTestClient(t, server.URL)
	_, err := c.ScheduleNewContactImport(context.Background(), ScheduleImport{
		CampaignID:   "test_campaign",
		FieldMapping: []string{"email"},
		Contacts:     [][]string{{"test@example.com"}},
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	var apiErr APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected APIError, got %T: %v", err, err)
	}
	if apiErr.HttpStatus != 429 || apiErr.Code != 1015 {
		t.Errorf("unexpected APIError: %+v", apiErr)
	}
}

func TestScheduleNewContactImport_Validation(t *testing.T) {
	c := newTestClient(t, "http://unused.invalid/")
	cases := []ScheduleImport{
		{FieldMapping: []string{"email"}},                                     // no contacts
		{Contacts: [][]string{{"a@b.com"}}},                                   // no mapping
		{FieldMapping: []string{"name"}, Contacts: [][]string{{"John"}}},      // no email field
		{FieldMapping: []string{"email", "name"}, Contacts: [][]string{{""}}}, // column mismatch
	}
	for i, in := range cases {
		if _, err := c.ScheduleNewContactImport(context.Background(), in); err == nil {
			t.Errorf("case %d: expected validation error, got nil", i)
		}
	}
}

// TestAllCampaigns_Pagination verifies the iterator walks every page using the
// TotalPages header.
func TestAllCampaigns_Pagination(t *testing.T) {
	const totalPages = 3
	var requestedPages []string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := r.URL.Query().Get("page")
		requestedPages = append(requestedPages, page)
		w.Header().Set("TotalPages", fmt.Sprintf("%d", totalPages))
		w.Header().Set("TotalCount", fmt.Sprintf("%d", totalPages))
		w.WriteHeader(http.StatusOK)
		// one campaign per page, id == page number
		_ = json.NewEncoder(w).Encode([]Campaign{{CampaignID: page, Name: "camp-" + page}})
	}))
	defer server.Close()

	c := newTestClient(t, server.URL)

	var got []string
	for camp, err := range c.AllCampaigns(context.Background()) {
		if err != nil {
			t.Fatalf("iterator error: %v", err)
		}
		got = append(got, camp.CampaignID)
	}

	if len(got) != totalPages {
		t.Fatalf("expected %d campaigns, got %d (%v)", totalPages, len(got), got)
	}
	if len(requestedPages) != totalPages {
		t.Errorf("expected %d page requests, got %v", totalPages, requestedPages)
	}
}

// TestAllCampaigns_EarlyBreak ensures breaking out of the loop stops fetching.
func TestAllCampaigns_EarlyBreak(t *testing.T) {
	var requests int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.Header().Set("TotalPages", "10")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode([]Campaign{{CampaignID: "a"}, {CampaignID: "b"}})
	}))
	defer server.Close()

	c := newTestClient(t, server.URL)
	count := 0
	for _, err := range c.AllCampaigns(context.Background()) {
		if err != nil {
			t.Fatalf("iterator error: %v", err)
		}
		count++
		break
	}
	if count != 1 {
		t.Errorf("expected to consume 1 item, got %d", count)
	}
	if requests != 1 {
		t.Errorf("expected 1 request after early break, got %d", requests)
	}
}

func TestGetAccount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("X-Auth-Token"); got != "api-key test-key" {
			t.Errorf("X-Auth-Token = %q", got)
		}
		if got := r.Header.Get("X-Domain"); got != "example.com" {
			t.Errorf("X-Domain = %q", got)
		}
		_ = json.NewEncoder(w).Encode(Account{AccountID: "acc1", Email: "a@b.com"})
	}))
	defer server.Close()

	c := newTestClient(t, server.URL)
	acc, err := c.GetAccount(context.Background())
	if err != nil {
		t.Fatalf("GetAccount: %v", err)
	}
	if acc.AccountID != "acc1" {
		t.Errorf("AccountID = %q, want acc1", acc.AccountID)
	}
}
