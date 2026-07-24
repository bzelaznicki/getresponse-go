package getresponse

type SearchContactsCondition interface {
	GetType() string
}

// CrmCondition represents a condition used for CRM conditions in GetResponse. This feature is Deprecated in GetResponse.
// It specifies the type of condition, the pipeline it applies to, and the stage within the pipeline.
//
// Fields:
// - ConditionType: Specifies the type of condition. For this struct, it is always set to "crm".
// - PipelineScope: Identifies the pipeline to which the condition applies.
// - StageScope: Specifies the stage within the pipeline. It can be "all" to include all stages or a specific stage identifier.
type CrmCondition struct {
	ConditionType string `json:"conditionType"`
	PipelineScope string `json:"pipelineScope"`
	StageScope    string `json:"stageScope"`
}

func (c *CrmCondition) GetType() string {
	return c.ConditionType
}

// NameCondition represents a condition used for filtering or searching contacts
// based on their name field in GetResponse.
//
// Fields:
//   - ConditionType: Specifies the type of condition. This should always be set to "name".
//   - OperatorType: Specifies the type of operator. This should always be set to "string_operator".
//   - Operator: Defines the comparison operation to be performed. Available options include:
//     "is", "is_not", "contains", "not_contains", "starts", "ends", "not_starts", "not_ends".
//   - Value: The value to be used in the comparison operation.
type NameCondition struct {
	ConditionType string `json:"conditionType"`
	OperatorType  string `json:"operatorType"`
	Operator      string `json:"operator"`
	Value         string `json:"value"`
}

func (n *NameCondition) GetType() string {
	return n.ConditionType
}

// EmailCondition represents a condition used for filtering or searching contacts
// based on their email field in GetResponse.
//
// Fields:
//   - ConditionType: Specifies the type of condition. This should always be set to "email".
//   - OperatorType: Specifies the type of operator. This should always be set to "string_operator".
//   - Operator: Defines the comparison operation to be performed. Available options include:
//     "is", "is_not", "contains", "not_contains", "starts", "ends", "not_starts", "not_ends".
//   - Value: The value to be used in the comparison operation.
type EmailCondition struct {
	ConditionType string `json:"conditionType"`
	OperatorType  string `json:"operatorType"`
	Operator      string `json:"operator"`
	Value         string `json:"value"`
}

func (e *EmailCondition) GetType() string {
	return e.ConditionType
}

// CustomCondition represents a condition used for filtering or searching contacts
// with custom field values.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must be set to "custom".
//   - OperatorType: Defines the type of operator to be used. Available options:
//     "string_operator_list", "string_operator", "numeric_operator", "date_operator".
//   - Operator: Specifies the operator based on the selected OperatorType:
//   - For "string_operator" and "string_operator_list": "is", "is_not", "contains",
//     "not_contains", "starts", "ends", "not_starts", "not_ends", "assigned", "not_assigned".
//   - For "numeric_operator": "numeric_lt", "numeric_gt", "numeric_eq", "numeric_not_eq",
//     "numeric_lt_eq", "numeric_gt_eq", "assigned", "not_assigned".
//   - For "date_operator": "date_to", "date_from", "custom", "specific_date", "assigned",
//     "not_assigned".
//   - Value: Specifies the value based on the selected Operator and OperatorType:
//   - Any string for "string_operator" and "string_operator_list".
//   - Stringified number for "numeric_operator".
//   - For "date_operator" with "specific_date": "today", "yesterday", "last_7_days",
//     "last_30_days", "last_n_days", "this_week", "last_week", "this_month", "last_month",
//     "last_2_months".
//   - Date string formatted as yyyy-mm-dd for "date_operator" with "date_to" or "date_from".
//   - ISO 8601 compliant date-time string for "date_operator" with "date_to" or "date_from".
//   - ISO 8601 compliant date interval string (<start_date>/<end_date>) for "date_operator"
//     with "custom".
//   - Scope: Specifies the scope of the condition, typically a customFieldId.
//   - NumberOfDays: Required only when "value" is set to "last_n_days". Represents the number
//     of days for the condition.
//   - IncludeCurrentPeriod: Determines if the current day is included in the chosen range.
//     This flag is applicable when "value" is set to "last_n_days".
type CustomCondition struct {
	ConditionType        string `json:"conditionType"`
	OperatorType         string `json:"operatorType"`
	Operator             string `json:"operator"`
	Value                string `json:"value"`
	Scope                string `json:"scope"`
	NumberOfDays         int    `json:"numberOfDays,omitempty"`
	IncludeCurrentPeriod bool   `json:"includeCurrentPeriod,omitempty"`
}

func (c *CustomCondition) GetType() string {
	return c.ConditionType
}

// SubscriptionDateCondition represents a condition used to filter contacts based on their subscription date.
//
// Fields:
// - ConditionType: Specifies the type of condition. For this struct, it should always be set to "subscription_date".
// - OperatorType: Specifies the type of operator. For this struct, it should always be set to "date_operator".
// - Operator: Defines the comparison operator to use. Available options include:
//   - "date_to": Matches dates up to a specific date.
//   - "date_from": Matches dates starting from a specific date.
//   - "specific_date": Matches a specific date.
//   - "custom": Allows custom date ranges.
//
// - Value: Specifies the value to compare against. Supported formats include:
//   - A date string in the format yyyy-mm-dd.
//   - A datetime string formatted in ISO 8601.
//   - A date interval string compliant with ISO 8601 in the format <start_date>/<end_date>.
//   - Predefined date ranges such as:
//   - "today"
//   - "yesterday"
//   - "last_7_days"
//   - "last_30_days"
//   - "last_n_days"
//   - "this_week"
//   - "last_week"
//   - "this_month"
//   - "last_month"
//   - "last_2_months"
type SubscriptionDateCondition struct {
	ConditionType string `json:"conditionType"`
	OperatorType  string `json:"operatorType"`
	Operator      string `json:"operator"`
	Value         string `json:"value"`
}

func (s *SubscriptionDateCondition) GetType() string {
	return s.ConditionType
}

// SubscriptionMethodCondition represents the conditions for filtering contacts
// based on their subscription method in the GetResponse system.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Should always be set to "subscription_method".
//   - Method: Indicates the subscription method. Available options include:
//     "webform", "import", "landing_page", "api", "email", "panel", "mobile",
//     "survey", "sales", "copy", "leads", "webinar", "chat", "website_builder_elegant",
//     "course", "premium_newsletter".
//   - WebformType: Specifies the type of webform when the method is "webform" or "webformsv2".
//     Available options are "all", "webforms", "webformsv2", "popups". This field is optional
//     and only required for webform-related methods.
//   - Value: Provides additional context or identifier for the subscription method. This field
//     is required for certain methods such as "import", "landing_page", "webinar", and
//     "website_builder_elegant". For example:
//   - If the method is "webform", the value corresponds to the webform, webformv2, or popup identifiers.
//   - If the method is "import", the value corresponds to the import identifier or "all" for all imports.
//   - If the method is "landing_page", the value corresponds to the landing page identifier or "all" for all landing pages.
//   - If the method is "webinar", the value corresponds to the webinar identifier or "all" for all webinars.
//   - If the method is "website_builder_elegant", the value corresponds to the website UUID or "all" for all websites.
type SubscriptionMethodCondition struct {
	ConditionType string `json:"conditionType"`
	Method        string `json:"method"`
	WebformType   string `json:"webformType,omitempty"`
	Value         string `json:"value"`
}

func (s *SubscriptionMethodCondition) GetType() string {
	return s.ConditionType
}

//Opened

// OpenedCondition represents a condition used to filter contacts based on
// whether they have opened a specific type of message. It includes the type
// of condition, the operator type, the operator, and the ID of the selected
// resource.
//
// Fields:
//   - ConditionType: Specifies the type of condition. For this struct, it should
//     always be set to "opened".
//   - OperatorType: Specifies the type of operator. For this struct, it should
//     always be set to "message_operator".
//   - Operator: Defines the type of message to filter by. Available options are
//     "autoresponder", "newsletter", "splittest", and "automation".
//   - Value: The ID of the selected resource (e.g., the specific message or
//     campaign being referenced).
type OpenedCondition struct {
	ConditionType string `json:"conditionType"`
	OperatorType  string `json:"operatorType"`
	Operator      string `json:"operator"`
	Value         string `json:"value"`
}

func (o *OpenedCondition) GetType() string {
	return o.ConditionType
}

// NotOpenedCondition represents a condition used to filter contacts who have not opened specific types of messages
// within a given time frame. This struct is used to define the parameters for such a condition.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must be set to "not_opened".
//   - Operator: Defines the type of messages to filter. Available options are:
//     "all", "autoresponder", "newsletter", "splittest", "automation".
//   - Scope: Specifies the scope of the condition. This field is optional and not required when "all" is selected.
//   - DateOperator: Determines the date range for the condition. Available options include:
//     "date_from", "never", "today", "yesterday", "last_7_days", "last_30_days", "last_n_days", "this_week", "this_month".
//   - Value: Provides additional information for certain DateOperator values. Required for "date_from" (ISO8601 date format)
//     and "last_n_days" (stringified number).
//   - IncludeCurrentPeriod: A flag indicating whether the current day should be included in the selected date range.
//     This flag is applicable for "last_7_days", "last_30_days", and "last_n_days".
type NotOpenedCondition struct {
	ConditionType        string `json:"conditionType"`
	Operator             string `json:"operator"`
	Scope                string `json:"scope,omitempty"`
	DateOperator         string `json:"dateOperator"`
	Value                string `json:"value,omitempty"`
	IncludeCurrentPeriod bool   `json:"includeCurrentPeriod,omitempty"`
}

func (n *NotOpenedCondition) GetType() string {
	return n.ConditionType
}

// PhaseCondition represents a condition used for filtering or querying based on a specific autoresponder day.
// It includes the type of condition, the operator type, the operator itself, and an optional value.
//
// Fields:
//   - ConditionType: Specifies the type of condition. For this struct, it should be set to "phase".
//   - OperatorType: Specifies the type of operator. For this struct, it should be set to "numeric_operator".
//   - Operator: Defines the comparison operator to use. Available options are:
//     "numeric_lt", "numeric_gt", "numeric_eq", "numeric_not_eq", "numeric_lt_eq", "numeric_gt_eq",
//     "assigned", and "not_assigned".
//   - Value: Represents the value to compare against. This field is optional and should not be included
//     when using the "assigned" or "not_assigned" operators.
type PhaseCondition struct {
	ConditionType string `json:"conditionType"`
	OperatorType  string `json:"operatorType"`
	Operator      string `json:"operator"`
	Value         int    `json:"value,omitempty"`
}

func (p *PhaseCondition) GetType() string {
	return p.ConditionType
}

// LastSendDateCondition represents a condition used to filter contacts based on the last send date.
//
// Fields:
// - ConditionType: Specifies the type of condition. Must be set to "last_send_date".
// - Operator: Defines the operator to use. Possible values are:
//   - "date_to": Filter contacts with a last send date up to a specific date.
//   - "date_from": Filter contacts with a last send date starting from a specific date.
//   - "specific_date": Filter contacts based on predefined date ranges such as "today", "yesterday", etc.
//   - "custom": Filter contacts using a custom time interval in ISO8601 format.
//   - OperatorType: Specifies the type of operator. Must be set to "date_operator".
//   - Value: The value associated with the operator. For "date_to" or "date_from", this should be a stringified date in ISO8601 format.
//     For "custom", this should be a stringified time interval in ISO8601 format. For "specific_date", this can be one of the following:
//     "today", "yesterday", "last_7_days", "last_30_days", "last_n_days", "this_week", "last_week", "this_month", "last_month", "last_2_months".
type LastSendDateCondition struct {
	ConditionType string `json:"conditionType"`
	Operator      string `json:"operator"`
	OperatorType  string `json:"operatorType"`
	Value         string `json:"value"`
}

func (l *LastSendDateCondition) GetType() string {
	return l.ConditionType
}

// LastClickDateCondition represents a condition used to filter contacts based on the last send date.
//
// Fields:
// - ConditionType: Specifies the type of condition. Must be set to "last_click_date".
// - Operator: Defines the operator to use. Possible values are:
//   - "date_to": Filter contacts with a last send date up to a specific date.
//   - "date_from": Filter contacts with a last send date starting from a specific date.
//   - "specific_date": Filter contacts based on predefined date ranges such as "today", "yesterday", etc.
//   - "custom": Filter contacts using a custom time interval in ISO8601 format.
//   - OperatorType: Specifies the type of operator. Must be set to "date_operator".
//   - Value: The value associated with the operator. For "date_to" or "date_from", this should be a stringified date in ISO8601 format.
//     For "custom", this should be a stringified time interval in ISO8601 format. For "specific_date", this can be one of the following:
//     "today", "yesterday", "last_7_days", "last_30_days", "last_n_days", "this_week", "last_week", "this_month", "last_month", "last_2_months".
type LastClickDateCondition struct {
	ConditionType string `json:"conditionType"`
	Operator      string `json:"operator"`
	OperatorType  string `json:"operatorType"`
	Value         int    `json:"value"`
}

func (l *LastClickDateCondition) GetType() string {
	return l.ConditionType
}

// LastOpenDateCondition represents a condition used to filter contacts based on the last send date.
//
// Fields:
// - ConditionType: Specifies the type of condition. Must be set to "last_open_date".
// - Operator: Defines the operator to use. Possible values are:
//   - "date_to": Filter contacts with a last send date up to a specific date.
//   - "date_from": Filter contacts with a last send date starting from a specific date.
//   - "specific_date": Filter contacts based on predefined date ranges such as "today", "yesterday", etc.
//   - "custom": Filter contacts using a custom time interval in ISO8601 format.
//   - OperatorType: Specifies the type of operator. Must be set to "date_operator".
//   - Value: The value associated with the operator. For "date_to" or "date_from", this should be a stringified date in ISO8601 format.
//     For "custom", this should be a stringified time interval in ISO8601 format. For "specific_date", this can be one of the following:
//     "today", "yesterday", "last_7_days", "last_30_days", "last_n_days", "this_week", "last_week", "this_month", "last_month", "last_2_months".
type LastOpenDateCondition struct {
	ConditionType string `json:"conditionType"`
	Operator      string `json:"operator"`
	OperatorType  string `json:"operatorType"`
	Value         int    `json:"value"`
}

func (l *LastOpenDateCondition) GetType() string {
	return l.ConditionType
}

//Webinar

// WebinarCondition represents the conditions for filtering contacts based on their
// participation in a webinar. It includes the type of condition, the scope of the
// webinar, the type of contact, and the specific webinar condition.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must be set to "webinar".
//   - Scope: Represents the Webinar ID to which the condition applies.
//   - ContactType: Specifies the type of contact. Possible values are "host", "listener",
//     "presenter", "registrant", or "all".
//   - WebinarCondition: Defines the participation condition for the webinar. Possible values
//     are "participated" or "not_participated".
type WebinarCondition struct {
	ConditionType    string `json:"conditionType"`
	Scope            string `json:"scope"`
	ContactType      string `json:"contactType"`
	WebinarCondition string `json:"webinarCondition"`
}

func (w *WebinarCondition) GetType() string {
	return w.ConditionType
}

// ClickedCondition represents the conditions for filtering contacts based on
// whether they clicked a specific link in a message. This structure is used
// to define the criteria for such a condition.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must be set to "clicked".
//   - Scope: The ID of the message to which the condition applies.
//   - Operator: Specifies the type of message. Can be one of "autoresponder",
//     "newsletter", "splittest", or "automation".
//   - OperatorType: Specifies the operator type. Must be set to "message_operator".
//   - ClickTrackId: The ID of the link click track or "all" to apply to all links.
type ClickedCondition struct {
	ConditionType string `json:"conditionType"`
	Scope         string `json:"scope"`
	Operator      string `json:"operator"`
	OperatorType  string `json:"operatorType"`
	ClickTrackId  string `json:"clickTrackId"`
}

func (c *ClickedCondition) GetType() string {
	return c.ConditionType
}

// NotClickedCondition represents the conditions for filtering contacts who have not clicked
// on a specific message or link within a given time frame. This struct is used to define
// the parameters for such a condition in the GetResponse API.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must be set to "not_clicked".
//   - Scope: The ID of the message. Must be left empty when using the "all" operator.
//   - Operator: Defines the type of message to filter by. Can be one of the following:
//     "all", "autoresponder", "newsletter", "splittest", or "automation".
//   - OperatorType: Specifies the operator type. Must be set to "complex_message_operator".
//   - ClickTrackId: Indicates the link to track. Can be "all" or the clicktrack ID of the message link.
//   - Value: Required only for specific date-based operators. For "date_from", it should be a stringified
//     date in ISO8601 format. For "last_n_days", it should be a stringified number.
//   - DateOperator: Specifies the date-based operator. Can be one of the following:
//     "date_from", "never", "today", "yesterday", "last_7_days", "last_30_days", "last_n_days",
//     "this_week", or "this_month".
//   - IncludeCurrentPeriod: A flag that determines whether the current day is included in the chosen
//     date range. This flag is applicable for "last_7_days", "last_30_days", and "last_n_days".
type NotClickedCondition struct {
	ConditionType        string `json:"conditionType"`
	Scope                string `json:"scope,omitempty"`
	Operator             string `json:"operator"`
	OperatorType         string `json:"operatorType"`
	ClickTrackId         string `json:"clickTrackId"`
	Value                string `json:"value,omitempty"`
	DateOperator         string `json:"dateOperator"`
	IncludeCurrentPeriod bool   `json:"includeCurrentPeriod"`
}

func (n *NotClickedCondition) GetType() string {
	return n.ConditionType
}

// SentCondition represents a condition used to filter contacts based on messages sent to them.
// It includes the type of condition, the message ID, the operator specifying the type of message,
// and the operator type which must be set to "message_operator".
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must be set to "sent".
//   - Value: The ID of the message used in the condition.
//   - Operator: Specifies the type of message. Can be one of "autoresponder", "newsletter",
//     "splittest", or "automation".
//   - OperatorType: Specifies the operator type. Must be set to "message_operator".
type SentCondition struct {
	ConditionType string `json:"conditionType"`
	Value         string `json:"value"`
	Operator      string `json:"operator"`
	OperatorType  string `json:"operatorType"`
}

func (s *SentCondition) GetType() string {
	return s.ConditionType
}

// Not sent
// NotSentCondition represents a condition used to filter contacts who have not been sent a specific message.
// It is used in search queries to identify contacts based on message sending criteria.
//
// Fields:
// - ConditionType: Specifies the type of condition. Must be set to "not_sent".
// - Value: The ID of the message that has not been sent.
// - Operator: Defines the type of message. Can be one of "autoresponder", "newsletter", "splittest", or "automation".
// - OperatorType: Specifies the operator type. Must be set to "message_operator".
type NotSentCondition struct {
	ConditionType string `json:"conditionType"`
	Value         string `json:"value"`
	Operator      string `json:"operator"`
	OperatorType  string `json:"operatorType"`
}

func (n *NotSentCondition) GetType() string {
	return n.ConditionType
}

// GeoCondition represents a condition used for geographic-based searches.
// It specifies the type of condition, the scope of the geographic attribute,
// the operator to apply, the type of operator, and the value to search for.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must be set to "geo".
//   - Scope: Defines the geographic attribute to search. Possible values include:
//     "country", "country_code", "region", "city", "longitude", "latitude",
//     "postal_code", "dma_code".
//   - Operator: Specifies the comparison operator to use. Possible values include:
//     "is", "is_not", "contains", "not_contains", "starts", "ends",
//     "not_starts", "not_ends".
//   - OperatorType: Specifies the type of operator. Must be set to "string_operator".
//   - Value: The value to search for within the specified scope.
type GeoCondition struct {
	ConditionType string `json:"conditionType"`
	Scope         string `json:"scope"`
	Operator      string `json:"operator"`
	OperatorType  string `json:"operatorType"`
	Value         string `json:"value"`
}

func (g *GeoCondition) GetType() string {
	return g.ConditionType
}

// GDPRCondition represents the conditions for marketing consent,
// indicating whether an individual has provided consent or not.
// It includes the type of condition, the consent status, and optional
// details about the consent date. This structure is used to define
// and query consent-related information in compliance with GDPR regulations.
// GDPRCondition represents the conditions related to marketing consent,
// specifically whether consent has been given or not, and optionally the
// date associated with the consent status.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must always be set to "gdpr".
//   - ConsentStatus: Indicates the consent status. Can be one of the following:
//     "given" or "not_given".
//   - ConsentDate: An optional nested structure that defines the date-related
//     conditions for the consent status. If omitted, the default operator is "anytime".
//   - Operator: Specifies the date operator. Can be one of "anytime", "date_from",
//     or "date_to".
//   - Date: The specific date in "Y-m-d" format. This field is required when the
//     operator is "date_from" or "date_to", and prohibited when the operator is "anytime".
type GDPRCondition struct {
	ConditionType string `json:"conditionType"`
	ConsentStatus string `json:"consentStatus"`
	ConsentDate   struct {
		Operator string `json:"operator"`
		Date     string `json:"date,omitempty"`
	} `json:"consentDate,omitempty"`
}

func (g *GDPRCondition) GetType() string {
	return g.ConditionType
}

// ScoreCondition represents a condition used for filtering or searching contacts
// based on their score. It includes the type of condition, the operator type,
// the specific operator (if applicable), and the score value.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must always be set to "score".
//   - OperatorType: Defines the type of operator to be used. Can be one of:
//     "numeric_operator" or "not_exists".
//   - Operator: Specifies the numeric comparison operator. This field is optional
//     and only required when OperatorType is "numeric_operator". Possible values
//     include:
//   - "numeric_lt": Less than
//   - "numeric_gt": Greater than
//   - "numeric_eq": Equal to
//   - "numeric_not_eq": Not equal to
//   - "numeric_lt_eq": Less than or equal to
//   - "numeric_gt_eq": Greater than or equal to
//   - Value: The score value to compare against
type ScoreCondition struct {
	ConditionType string `json:"conditionType"`
	OperatorType  string `json:"operatorType"`
	Operator      string `json:"operator,omitempty"`
	Value         int    `json:"value"`
}

func (s *ScoreCondition) GetType() string {
	return s.ConditionType
}

// Engagement score
// EngagementScoreCondition represents a condition used to filter contacts based on their engagement score.
// The engagement score is a numeric value between 1 and 5, and the condition specifies how to compare this value.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must always be set to "engagement_score".
//   - Operator: Defines the comparison operator to use. Possible values are:
//     "numeric_lt" (less than), "numeric_gt" (greater than), "numeric_eq" (equal to),
//     "numeric_not_eq" (not equal to), "numeric_lt_eq" (less than or equal to),
//     "numeric_gt_eq" (greater than or equal to).
//   - Value: The numeric value (1 to 5) to compare against, represented as an integer.
type EngagementScoreCondition struct {
	ConditionType string `json:"conditionType"`
	Operator      string `json:"operator"`
	Value         int    `json:"value"`
}

func (e *EngagementScoreCondition) GetType() string {
	return e.ConditionType
}

// Tag
// TagCondition represents a condition used to filter contacts based on tags in the GetResponse API.
// It specifies the type of condition, the operator to apply, and the value (tag ID) to match.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must be set to "tag".
//   - OperatorType: Specifies the type of operator. Must be set to "exists".
//   - Operator: Defines the operation to perform. Available values are "exists" (to check if the tag exists)
//     and "not_exists" (to check if the tag does not exist).
//   - Value: The ID of the tag to be used in the condition.
type TagCondition struct {
	ConditionType string `json:"conditionType"` // set to "tag"
	OperatorType  string `json:"operatorType"`  // set to "exists"
	Operator      string `json:"operator"`      // available values: "exists", "not_exists"
	Value         string `json:"value"`         // tagId
}

func (t *TagCondition) GetType() string {
	return t.ConditionType
}

// GoalCondition represents a condition for Goals in GetResponse.
// It defines the type of condition, the operator to be applied, the value for comparison,
// and the scope (goal ID) to which the condition applies.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must be set to "goal".
//   - OperatorType: Specifies the type of operator. Must be set to "numeric_operator".
//   - Operator: Defines the comparison operator. Can be one of the following:
//     "numeric_lt", "numeric_gt", "numeric_eq", "numeric_not_eq", "numeric_lt_eq",
//     "numeric_gt_eq", "assigned", "not_assigned".
//   - Value: The value to be compared for the specified relation. This field is optional
//     and will be omitted if not set.
//   - Scope: Specifies the goal ID to which this condition applies.
type GoalCondition struct {
	ConditionType string `json:"conditionType"`
	OperatorType  string `json:"operatorType"`
	Operator      string `json:"operator"`
	Value         int    `json:"value,omitempty"`
	Scope         string `json:"scope"`
}

func (g *GoalCondition) GetType() string {
	return g.ConditionType
}

// EcommerceNumberOfPurchasesCondition represents a condition used to filter contacts
// based on the number of purchases in an e-commerce context.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must be set to "ecommerce_number_of_purchases".
//   - OperatorType: Specifies the type of operator. Must be set to "numeric_operator".
//   - Operator: Specifies the comparison operator. Can be one of the following:
//     "numeric_lt", "numeric_gt", "numeric_eq", "numeric_not_eq", "numeric_lt_eq", "numeric_gt_eq".
//   - Value: Specifies the number of purchases to compare against.
//   - Scope: Specifies the scope of the condition. Can be "all" or a specific shop ID.
//   - Date: Specifies the date range for the condition. Includes the following fields:
//   - Operator: Specifies the date operator. Can be one of the following:
//     "anytime", "today", "yesterday", "last_n_days", "this_week", "last_week",
//     "this_month", "last_month", "last_2_months", "custom".
//   - NumberOfDays: Required for "last_n_days", specifies the number of days.
//   - IncludeCurrentPeriod: Determines if the current day is included in the range
//     (required for the "last_n_days" operator).
//   - Range: Specifies a custom date range in ISO8601 interval format (<start_date>/<end_date>),
//     allowed only for the "custom" operator.
type EcommerceNumberOfPurchasesCondition struct {
	ConditionType string `json:"conditionType"`
	OperatorType  string `json:"operatorType"`
	Operator      string `json:"operator"`
	Value         int    `json:"value"`
	Scope         string `json:"scope"`
	Date          struct {
		Operator             string `json:"operator"`
		NumberOfDays         int    `json:"numberOfDays,omitempty"`
		IncludeCurrentPeriod bool   `json:"includeCurrentPeriod"`
		Range                string `json:"range"`
	} `json:"date"`
}

func (e *EcommerceNumberOfPurchasesCondition) GetType() string {
	return e.ConditionType
}

// E-commerce total spent

// EcommerceTotalSpentCondition represents a condition used to filter contacts based on their total
// spending in an e-commerce context. This structure includes details about the condition type,
// operator, scope, value, currency, and an optional date range.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must be set to "ecommerce_total_spent".
//   - OperatorType: Specifies the type of operator. Must be set to "numeric_operator".
//   - Operator: Defines the comparison operator. Can be one of the following:
//     "numeric_lt" (less than), "numeric_gt" (greater than), "numeric_eq" (equal to),
//     "numeric_not_eq" (not equal to), "numeric_lt_eq" (less than or equal to),
//     "numeric_gt_eq" (greater than or equal to).
//   - Scope: Specifies the scope of the condition. Can be "all" or a specific Shop ID.
//   - Value: The total spending value to compare against.
//   - Currency: The currency code for the spending value, following the ISO4217 standard
//     (e.g., USD, GBP, EUR, PLN).
//   - Date: A nested structure that defines the date range for the condition. Includes:
//   - Operator: Specifies the date operator. Can be one of the following:
//     "anytime", "today", "yesterday", "last_n_days", "this_week", "last_week",
//     "this_month", "last_month", "last_2_months", or "custom".
//   - NumberOfDays: Required for the "last_n_days" operator. Specifies the number of days.
//   - IncludeCurrentPeriod: Determines whether the current day is included in the range.
//     Required for the "last_n_days" operator.
//   - Range: Specifies a custom date range in ISO8601 format (<start_date>/<end_date>).
//     Allowed only for the "custom" operator.
type EcommerceTotalSpentCondition struct {
	ConditionType string  `json:"conditionType"`
	OperatorType  string  `json:"operatorType"`
	Operator      string  `json:"operator"`
	Scope         string  `json:"scope"`
	Value         float64 `json:"value"`
	Currency      string  `json:"currency"`
	Date          struct {
		Operator             string `json:"operator"`
		NumberOfDays         int    `json:"numberOfDays,omitempty"`
		IncludeCurrentPeriod bool   `json:"includeCurrentPeriod"`
		Range                string `json:"range"`
	} `json:"date"`
}

func (e *EcommerceTotalSpentCondition) GetType() string {
	return e.ConditionType
}

// E-commerce product purchased condition

// EcommerceProductPurchasedCondition represents the condition for filtering contacts
// based on their ecommerce product purchase history. This structure is used to define
// the criteria for searching contacts who have purchased specific products within a
// given scope, category, or date range.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must be set to "ecommerce_product_purchased".
//   - ShopScope: Defines the shop scope. Can be "all" or a specific Shop ID.
//   - CategoryScope: Defines the category scope. Can be "all" or a specific Category ID.
//   - OperatorType: Specifies the operator type. Must be set to "equal_operator".
//   - Operator: Defines the operator for the condition. Allowed values are "is" or "is_not".
//   - ProductScope: Specifies the product scope. Can be "all" or a specific Product ID.
//   - DateOperator: Defines the date range operator. Allowed values include:
//     "today", "yesterday", "last_7_days", "last_30_days", "last_n_days", "this_week",
//     "last_week", "this_month", "last_month", "last_2_months", "all_time", "date_to",
//     "date_from", or "custom".
//   - Value: Required for "date_to", "date_from", or "custom" date operators. For "date_to"
//     and "date_from", it should be a date formatted as "yyyy-mm-dd". For "custom", it should
//     be an ISO8601-compliant date interval string in the "<start_date>/<end_date>" format.
//   - NumberOfDays: Required for the "last_n_days" operator. Specifies the number of days
//     for the date range.
//   - IncludeCurrentPeriod: A flag that determines if the current day is included in the
//     chosen date range. This flag is required for the "last_n_days" date operator and can
//     also be used with "last_7_days" and "last_30_days" date operators.
type EcommerceProductPurchasedCondition struct {
	ConditionType        string `json:"conditionType"`
	ShopScope            string `json:"shopScope"`
	CategoryScope        string `json:"categoryScope"`
	OperatorType         string `json:"operatorType"`
	Operator             string `json:"operator"`
	ProductScope         string `json:"productScope"`
	DateOperator         string `json:"dateOperator"`
	Value                string `json:"value,omitempty"`
	NumberOfDays         int    `json:"numberOfDays,omitempty"`
	IncludeCurrentPeriod bool   `json:"includeCurrentPeriod,omitempty"`
}

func (e *EcommerceProductPurchasedCondition) GetType() string {
	return e.ConditionType
}

// EcommerceBrandPurchasedCondition represents a condition used to filter contacts
// based on their purchase history of a specific brand in an e-commerce shop.
//
// Fields:
// - ConditionType: Specifies the type of condition. Must be set to "ecommerce_brand_purchased".
// - Scope: The ID of the shop where the purchase was made.
// - OperatorType: Specifies the type of operator. Must be set to "equal_operator".
// - Operator: Defines the comparison operator. Can be "is" or "not_is".
// - Value: The brand name used for the condition.
type EcommerceBrandPurchasedCondition struct {
	ConditionType string `json:"conditionType"` //Must be set to "ecommerce_brand_purchased"
	Scope         string `json:"scope"`         //The ID of the shop
	OperatorType  string `json:"operatorType"`  //Must be set to "equal_operator"
	Operator      string `json:"operator"`      //Any of "is" "not_is"
	Value         string `json:"value"`         //The brand
}

func (e *EcommerceBrandPurchasedCondition) GetType() string {
	return e.ConditionType
}

// EcommerceAbandonedCartCondition represents the condition for filtering contacts
// based on abandoned e-commerce carts in the GetResponse system.
//
// Fields:
// - ConditionType: Specifies the type of condition. Must be set to "ecommerce_abandoned_cart".
// - ShopScope: Defines the scope of the shop. Can be "all" or a specific Shop ID.
// - CartValue: Contains details about the cart's value, including:
//   - Currency: The currency code according to ISO4217 (e.g., USD, GBP, EUR, PLN).
//   - Value: The numeric value of the cart.
//   - Operator: The comparison operator for the cart value. Possible values are:
//     "numeric_lt", "numeric_gt", "numeric_eq", "numeric_not_eq",
//     "numeric_lt_eq", "numeric_gt_eq".
//   - DateOperator: Specifies the date condition. Possible values include:
//     "date_from", "anytime", "never", "today", "yesterday", "last_n_days",
//     "this_week", "last_week", "this_month", "last_month", "custom".
//   - Value: Required for "date_from" and "last_n_days". For "date_from", it should
//     be a date formatted in ISO8601. For "last_n_days", it should be a stringified number.
//   - IncludeCurrentPeriod: Indicates whether to include the current period. Required
//     only for "last_n_days".
//   - Range: Allowed only for "custom". Specifies a date interval string compliant
//     with ISO8601, formatted as <start_date>/<end_date>.
type EcommerceAbandonedCartCondition struct {
	ConditionType string `json:"conditionType"`
	ShopScope     string `json:"shopScope"`
	CartValue     struct {
		Currency string  `json:"currency"`
		Value    float64 `json:"value"`
		Operator string  `json:"operator"`
	} `json:"cartValue"`
	DateOperator         string `json:"dateOperator"`
	Value                string `json:"value,omitempty"`
	IncludeCurrentPeriod bool   `json:"includeCurrentPeriod,omitempty"`
	Range                string `json:"range,omitempty"`
}

func (e *EcommerceAbandonedCartCondition) GetType() string {
	return e.ConditionType
}

// SmsSentCondition represents the condition for filtering contacts based on SMS sending criteria.
// It includes the type of condition, the SMS ID, and optionally the delivery status.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must be set to "sms_sent".
//   - SMSID: The unique identifier of the SMS.
//   - DeliveryStatus: (Optional) Specifies the delivery status of the SMS.
//     Can be one of "delivered", "undelivered", or "any".
//     Omit this field if the delivery status is not relevant.
type SmsSentCondition struct {
	ConditionType  string `json:"conditionType"`            //Must be sent to "sms_sent"
	SMSID          string `json:"smsId"`                    //ID of the SMS
	DeliveryStatus string `json:"deliveryStatus,omitempty"` //Any of "delivered" "undelivered" "any". Omit this field if the  delvivery status is not relevant.
}

func (s *SmsSentCondition) GetType() string {
	return s.ConditionType
}

// SmsLinkClickedCondition represents a condition used to filter contacts
// based on whether they clicked a link in an SMS message.
//
// Fields:
// - ConditionType: Specifies the type of condition. Must always be set to "sms_link_clicked".
// - SMSID: The unique identifier of the SMS message.
// - ClickTrackID: The unique identifier of the SMS click track.
type SmsLinkClickedCondition struct {
	ConditionType string `json:"conditionType"`
	SMSID         string `json:"smsId"`
	ClickTrackID  string `json:"clickTrackId"` //ID of the SMS click track
}

func (s *SmsLinkClickedCondition) GetType() string {
	return s.ConditionType
}

// SmsLinkClickedCondition represents a condition used to filter contacts
// based on whether they have not clicked a link in an SMS message.
//
// Fields:
// - ConditionType: Specifies the type of condition. Must always be set to "sms_link_not_clicked".
// - SMSID: The unique identifier of the SMS message.
// - ClickTrackID: The unique identifier of the SMS click track.
type SmsLinkNotClickedCondition struct {
	ConditionType string `json:"conditionType"`
	SMSID         string `json:"smsId"`
	ClickTrackID  string `json:"clickTrackId"`
}

// CustomEventCondition represents the conditions for filtering contacts based on a custom event.
// It includes the type of condition, the custom event ID, the occurrence status, and date-related filters.
//
// Fields:
//   - ConditionType: Specifies the type of condition. Must be set to "custom_event".
//   - CustomEventId: The ID of the custom event to filter by.
//   - Occurrence: Indicates whether the event has "occurred" or "not_occurred".
//   - DateOperator: Specifies the date filter operator. Possible values include:
//     "today", "yesterday", "last_7_days", "last_30_days", "last_n_days", "this_week",
//     "last_week", "this_month", "last_month", "last_2_months", "anytime", "date_to",
//     "date_from", "custom".
//   - Date: (Optional) Required only for "date_to", "date_from", and "custom" operators.
//     Represents the date or date range in one of the following formats:
//   - yyyy-mm-dd
//   - ISO8601 datetime string
//   - ISO 8601 date interval string in the format <start_date>/<end_date>.
type CustomEventCondition struct {
	ConditionType string `json:"conditionType"`  //must be set to "custom_event"
	CustomEventId string `json:"customEventId"`  //ID of the custom event
	Occurrence    string `json:"occurrence"`     //Any of "occurred" "not_occurred"
	DateOperator  string `json:"dateOperator"`   //Any of "today" "yesterday" "last_7_days" "last_30_days" "last_n_days" "this_week" "last_week" "this_month" "last_month" "last_2_months" "anytime" "date_to" "date_from" "custom"
	Date          string `json:"date,omitempty"` //Required and used only for date_to, date_from, and customDate. Can be either the date formatted as yyyy-mm-dd, ISO8601 datetime string, or a date interval string compliant with ISO 8601, supported format: <start_date>/<end_date>
}

func (c *CustomEventCondition) GetType() string {
	return c.ConditionType
}

// Generic fallback condition for unsupported types. Please report if this case occurs.
type GenericCondition map[string]any //fallback

func (g *GenericCondition) GetType() string {
	if t, ok := (*g)["conditionType"].(string); ok {
		return t
	}
	return ""
}
