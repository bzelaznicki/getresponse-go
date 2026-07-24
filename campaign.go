package getresponse

import (
	"context"
	"net/http"
)

// Campaign represents a GetResponse campaign (list).
type Campaign struct {
	Description  string `json:"description,omitzero"`
	CampaignID   string `json:"campaignId"`
	Name         string `json:"name"`
	TechName     string `json:"techName,omitzero"`
	LanguageCode string `json:"languageCode,omitzero"`
	IsDefault    string `json:"isDefault,omitzero"`
	CreatedOn    string `json:"createdOn,omitzero"`
	Href         string `json:"href"`
	Postal       struct {
		AddPostalToMessages string `json:"addPostalToMessages"`
		City                string `json:"city"`
		CompanyName         string `json:"companyName"`
		Country             string `json:"country"`
		Design              string `json:"design"`
		State               string `json:"state"`
		Street              string `json:"street"`
		ZipCode             string `json:"zipCode"`
	} `json:"postal,omitzero"`
	Confirmation struct {
		FromField                         FromField `json:"fromField"`
		RedirectType                      string    `json:"redirectType"`
		MimeType                          string    `json:"mimeType"`
		RedirectURL                       string    `json:"redirectUrl"`
		ReplyTo                           FromField `json:"replyTo"`
		SubscriptionConfirmationBodyId    string    `json:"subscriptionConfirmationBodyId"`
		SubscriptionConfirmationSubjectId string    `json:"subscriptionConfirmationSubjectId"`
	} `json:"confirmation,omitzero"`

	OptInTypes struct {
		Email   string `json:"email"`
		Api     string `json:"api"`
		Import  string `json:"import"`
		Webform string `json:"webform"`
	} `json:"optInTypes,omitzero"`

	SubscriptionNotifications struct {
		Status     string      `json:"status"`
		Recipients []FromField `json:"recipients"`
	} `json:"subscriptionNotifications,omitzero"`

	Profile struct {
		Description   string `json:"description"`
		IndustryTagID string `json:"industryTagId"`
		Logo          string `json:"logo"`
		LogoLinkURL   string `json:"logoLinkUrl"`
		Title         string `json:"title"`
	} `json:"profile,omitzero"`
}

// GetCampaigns retrieves a single page of campaigns. Use the returned
// ResponseHeader for pagination, or AllCampaigns to iterate every page.
func (c *Client) GetCampaigns(ctx context.Context, opts ...QueryOption) ([]Campaign, ResponseHeader, error) {
	return do[[]Campaign](ctx, c, http.MethodGet, "campaigns", buildQuery(opts...), nil)
}
