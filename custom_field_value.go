package getresponse

// CustomFieldValue represents the value of a custom field on a contact.
type CustomFieldValue struct {
	CustomFieldID string   `json:"customFieldId"`
	Name          string   `json:"name"`
	Value         []string `json:"value"`
	Values        []string `json:"values"`
	Type          string   `json:"type"`
	FieldType     string   `json:"fieldType"`
	ValueType     string   `json:"valueType"`
}
