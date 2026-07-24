package getresponse

import (
	"context"
	"fmt"
	"net/http"
)

// Tag represents a GetResponse tag.
type Tag struct {
	TagID     string `json:"tagId,omitzero"`
	CreatedAt string `json:"createdAt,omitzero"`
	Name      string `json:"name"`
	Color     string `json:"color,omitzero"`
}

// GetTagsList retrieves a single page of tags. Use AllTags to iterate every
// page.
func (c *Client) GetTagsList(ctx context.Context, opts ...QueryOption) ([]Tag, ResponseHeader, error) {
	return do[[]Tag](ctx, c, http.MethodGet, "tags", buildQuery(opts...), nil)
}

// CreateTag creates a new tag. The name must be 2-64 characters.
func (c *Client) CreateTag(ctx context.Context, t Tag) (Tag, error) {
	if len(t.Name) < 2 || len(t.Name) > 64 {
		return Tag{}, fmt.Errorf("getresponse: tag name must be between 2 and 64 characters")
	}
	created, _, err := do[Tag](ctx, c, http.MethodPost, "tags", nil, t)
	return created, err
}
