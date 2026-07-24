package getresponse

// FromField represents a "from" (sender) address configured on an account.
type FromField struct {
	FromFieldID    string `json:"fromFieldId"`
	Href           string `json:"href"`
	Email          string `json:"email,omitzero"`
	RewrittenEmail string `json:"rewrittenEmail,omitzero"`
	Name           string `json:"name,omitzero"`
	IsActive       string `json:"isActive,omitzero"`
	IsDefault      string `json:"isDefault,omitzero"`
	CreatedOn      string `json:"createdOn,omitzero"`
	Domain         struct {
		Status      string `json:"status"`
		DKIMWarning string `json:"DKIMWarning"`
	} `json:"domain,omitzero"`
}
