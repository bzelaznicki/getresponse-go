package getresponse

import (
	"context"
	"iter"
)

// defaultPageSize is the per-page size used by the All* iterators.
const defaultPageSize = 1000

// pageFetcher fetches a single page of T given query options.
type pageFetcher[T any] func(ctx context.Context, opts ...QueryOption) ([]T, ResponseHeader, error)

// paginate turns a page-fetching function into a range-over-func iterator that
// walks every page. The iterator yields (zero, err) once and stops on the first
// error. Caller-supplied options are preserved; the iterator sets page/perPage.
func paginate[T any](ctx context.Context, fetch pageFetcher[T], opts []QueryOption) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for page := 1; ; page++ {
			pageOpts := append(append([]QueryOption{}, opts...), WithPage(page), WithPerPage(defaultPageSize))
			items, header, err := fetch(ctx, pageOpts...)
			if err != nil {
				var zero T
				yield(zero, err)
				return
			}
			for _, item := range items {
				if !yield(item, nil) {
					return
				}
			}
			// TotalPages is 0 when the endpoint returned no pagination header;
			// treat that as a single page to avoid looping forever.
			if header.TotalPages <= page {
				return
			}
		}
	}
}

// AllCampaigns iterates every campaign across all pages.
func (c *Client) AllCampaigns(ctx context.Context, opts ...QueryOption) iter.Seq2[Campaign, error] {
	return paginate(ctx, c.GetCampaigns, opts)
}

// AllTags iterates every tag across all pages.
func (c *Client) AllTags(ctx context.Context, opts ...QueryOption) iter.Seq2[Tag, error] {
	return paginate(ctx, c.GetTagsList, opts)
}

// AllCustomFields iterates every custom field across all pages.
func (c *Client) AllCustomFields(ctx context.Context, opts ...QueryOption) iter.Seq2[CustomField, error] {
	return paginate(ctx, c.GetCustomFieldsList, opts)
}

// AllSearchContacts iterates every saved search (segment) across all pages.
func (c *Client) AllSearchContacts(ctx context.Context, opts ...QueryOption) iter.Seq2[SearchContacts, error] {
	return paginate(ctx, c.GetSearchContactsList, opts)
}

// AllContactsFromCampaign iterates every contact in a campaign across all pages.
func (c *Client) AllContactsFromCampaign(ctx context.Context, campaignID string, opts ...QueryOption) iter.Seq2[Contact, error] {
	fetch := func(ctx context.Context, o ...QueryOption) ([]Contact, ResponseHeader, error) {
		return c.GetContactsFromCampaign(ctx, campaignID, o...)
	}
	return paginate(ctx, fetch, opts)
}

// AllContactsByIDSearchContacts iterates every contact in a saved search
// (segment) across all pages.
func (c *Client) AllContactsByIDSearchContacts(ctx context.Context, searchContactsID string, opts ...QueryOption) iter.Seq2[Contact, error] {
	fetch := func(ctx context.Context, o ...QueryOption) ([]Contact, ResponseHeader, error) {
		return c.GetContactsByIDSearchContacts(ctx, searchContactsID, o...)
	}
	return paginate(ctx, fetch, opts)
}

// AllContactsFromSearchContactsConditions iterates every contact matching the
// supplied ad-hoc conditions across all pages.
func (c *Client) AllContactsFromSearchContactsConditions(ctx context.Context, conditions SearchContacts, opts ...QueryOption) iter.Seq2[Contact, error] {
	fetch := func(ctx context.Context, o ...QueryOption) ([]Contact, ResponseHeader, error) {
		return c.GetContactsFromSearchContactsConditions(ctx, conditions, o...)
	}
	return paginate(ctx, fetch, opts)
}
