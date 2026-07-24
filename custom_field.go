package getresponse

import (
	"context"
	"fmt"
	"net/http"
)

// CustomField represents a GetResponse custom field definition.
type CustomField struct {
	CustomFieldID string   `json:"customFieldId,omitzero"`
	Href          string   `json:"href,omitzero"`
	Name          string   `json:"name"`
	Type          string   `json:"type"`
	ValueType     string   `json:"valueType"`
	Format        string   `json:"format"`
	FieldType     string   `json:"fieldType"`
	Hidden        string   `json:"hidden"`
	Values        []string `json:"values,omitzero"`
}

var validCustomFieldTypes = map[string]struct{}{
	"string": {}, "number": {}, "date": {}, "datetime": {}, "country": {},
	"currency": {}, "phone": {}, "gender": {}, "ip": {}, "url": {},
}

var validCustomFieldFormats = map[string]struct{}{
	"text": {}, "textarea": {}, "radio": {}, "checkbox": {},
	"single_select": {}, "multi_select": {},
}

// GetCustomFieldsList retrieves a single page of custom fields. Use
// AllCustomFields to iterate every page.
func (c *Client) GetCustomFieldsList(ctx context.Context, opts ...QueryOption) ([]CustomField, ResponseHeader, error) {
	return do[[]CustomField](ctx, c, http.MethodGet, "custom-fields", buildQuery(opts...), nil)
}

// CreateCustomField creates a new custom field. The name must be 2-64
// characters and the type and format must be among the values accepted by
// GetResponse.
func (c *Client) CreateCustomField(ctx context.Context, cf CustomField) (CustomField, error) {
	if len(cf.Name) < 2 || len(cf.Name) > 64 {
		return CustomField{}, fmt.Errorf("getresponse: custom field name must be between 2 and 64 characters")
	}
	if _, ok := validCustomFieldTypes[cf.Type]; !ok {
		return CustomField{}, fmt.Errorf("getresponse: invalid custom field type: %s", cf.Type)
	}
	if _, ok := validCustomFieldFormats[cf.Format]; !ok {
		return CustomField{}, fmt.Errorf("getresponse: invalid custom field format: %s", cf.Format)
	}

	created, _, err := do[CustomField](ctx, c, http.MethodPost, "custom-fields", nil, cf)
	return created, err
}
