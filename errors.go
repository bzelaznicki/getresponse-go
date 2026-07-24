package getresponse

import "fmt"

// APIError represents an error returned by the GetResponse API. It is returned
// by any client method whose request produced a non-2xx HTTP response.
type APIError struct {
	HttpStatus      int            `json:"httpStatus"`
	Code            int            `json:"code"`
	CodeDescription string         `json:"codeDescription"`
	Message         string         `json:"message"`
	MoreInfo        string         `json:"moreInfo"`
	Context         map[string]any `json:"context"`
	Uuid            string         `json:"uuid"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("GetResponse API Error: %s (Code: %d, HTTP: %d) - %s", e.Message, e.Code, e.HttpStatus, e.CodeDescription)
}
