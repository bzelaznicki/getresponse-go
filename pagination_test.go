package getresponse

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// newPagedCampaignServer serves total campaigns the way the real campaigns
// endpoint does: a perPage above maxPageSize is silently clamped, while the
// TotalPages header is still calculated from the perPage that was requested.
// Requesting 1000 therefore answers "TotalPages: 1" next to a clamped page.
func newPagedCampaignServer(t *testing.T, total, maxPageSize int, sendCounters bool, requests *int) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if requests != nil {
			*requests++
		}

		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page < 1 {
			page = 1
		}
		perPage, err := strconv.Atoi(r.URL.Query().Get("perPage"))
		if err != nil || perPage < 1 {
			perPage = maxPageSize
		}
		pageSize := min(perPage, maxPageSize)

		campaigns := make([]Campaign, 0, pageSize)
		for i := (page - 1) * pageSize; i < page*pageSize && i < total; i++ {
			campaigns = append(campaigns, Campaign{
				CampaignID: fmt.Sprintf("campaign-%d", i),
				Name:       fmt.Sprintf("List %d", i),
				TechName:   fmt.Sprintf("tech_%d", i),
			})
		}

		w.Header().Set("Content-Type", "application/json")
		if sendCounters {
			w.Header().Set("TotalCount", strconv.Itoa(total))
			w.Header().Set("CurrentPage", strconv.Itoa(page))
			w.Header().Set("TotalPages", strconv.Itoa((total+perPage-1)/perPage))
		}
		_ = json.NewEncoder(w).Encode(campaigns)
	}))
}

func collectCampaignIDs(t *testing.T, c *Client) []string {
	t.Helper()

	ids := []string{}
	for campaign, err := range c.AllCampaigns(context.Background()) {
		if err != nil {
			t.Fatalf("AllCampaigns: %v", err)
		}
		ids = append(ids, campaign.CampaignID)
	}
	return ids
}

// The endpoint caps pages at 100 while reporting TotalPages for the requested
// 1000, so trusting TotalPages alone stopped the walk after the first page.
func TestPaginateWalksPastClampedPageSize(t *testing.T) {
	requests := 0
	server := newPagedCampaignServer(t, 250, 100, true, &requests)
	defer server.Close()

	ids := collectCampaignIDs(t, newTestClient(t, server.URL))

	if len(ids) != 250 {
		t.Fatalf("expected 250 campaigns, got %d", len(ids))
	}
	if ids[0] != "campaign-0" || ids[249] != "campaign-249" {
		t.Fatalf("unexpected first/last campaign: %s / %s", ids[0], ids[249])
	}
	if requests != 3 {
		t.Fatalf("expected 3 requests, got %d", requests)
	}
}

func TestPaginateHonouredPageSizeUsesSingleRequest(t *testing.T) {
	requests := 0
	server := newPagedCampaignServer(t, 250, defaultPageSize, true, &requests)
	defer server.Close()

	if ids := collectCampaignIDs(t, newTestClient(t, server.URL)); len(ids) != 250 {
		t.Fatalf("expected 250 campaigns, got %d", len(ids))
	}
	if requests != 1 {
		t.Fatalf("expected a single request, got %d", requests)
	}
}

func TestPaginateWithoutCountersReadsOnePage(t *testing.T) {
	requests := 0
	server := newPagedCampaignServer(t, 250, 100, false, &requests)
	defer server.Close()

	if ids := collectCampaignIDs(t, newTestClient(t, server.URL)); len(ids) != 100 {
		t.Fatalf("expected the single page of 100, got %d", len(ids))
	}
	if requests != 1 {
		t.Fatalf("expected a single request, got %d", requests)
	}
}

func TestPaginateStopsWhenCallerBreaks(t *testing.T) {
	requests := 0
	server := newPagedCampaignServer(t, 250, 100, true, &requests)
	defer server.Close()

	seen := 0
	for range newTestClient(t, server.URL).AllCampaigns(context.Background()) {
		seen++
		if seen == 5 {
			break
		}
	}

	if seen != 5 {
		t.Fatalf("expected to stop after 5 campaigns, got %d", seen)
	}
	if requests != 1 {
		t.Fatalf("expected no further pages after the break, got %d requests", requests)
	}
}

func TestPaginateYieldsErrorOnce(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	errs := 0
	for _, err := range newTestClient(t, server.URL).AllCampaigns(context.Background()) {
		if err != nil {
			errs++
			continue
		}
		t.Fatal("expected no campaigns to be yielded")
	}

	if errs != 1 {
		t.Fatalf("expected exactly one error, got %d", errs)
	}
}
